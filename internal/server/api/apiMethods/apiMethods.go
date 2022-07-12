package apiMethods

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"io/ioutil"
	"modules/internal/server/UserModel"
	"modules/internal/server/bondsInfo"
	"modules/internal/server/bondsInfo/providers"
	"modules/internal/server/hash"
	helper "modules/internal/server/helpers"
	mongo "modules/internal/server/mongoDB"
	"net/http"
	"os"
)

type ApiMethodsService struct {
	method       ApiMethods
	mongoService mongo.MongoService
}

type ApiMethods interface {
	YearPost(c *gin.Context)
	BondsPost(c *gin.Context)
	Delete(c *gin.Context)
	Login(c *gin.Context)
	Register(c *gin.Context)
}

func (method *ApiMethodsService) YearPost(c *gin.Context) {
	service := method.mongoService
	provider := providers.MoexBondInfoProvider{}
	yearMap := make(map[string]string)

	err := c.BindJSON(&yearMap)
	if err != nil {
		fmt.Println("YearPost: err: ", err)
	}

	id := c.GetString("uid")

	res := service.FindUserInDBById(id)
	if res.User_id == id {
		response := provider.GetBondsForYear(yearMap["year"], res.Bonds)
		c.JSON(http.StatusOK, struct {
			Bonds []bondsInfo.UserBonds `json:"bonds"`
		}{
			response,
		})
		fmt.Println()
	} else {
		fmt.Println("need to login")
	}

}

func (method *ApiMethodsService) BondsPost(c *gin.Context) {
	service := mongo.MongoService{}
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println("err: ", err)
	}

	bond := bondsInfo.UserBonds{}
	err = json.Unmarshal(jsonData, &bond)
	if err != nil {
		fmt.Println("err: ", err)
	}
	id := c.GetString("uid")

	res := service.FindUserInDBById(id)
	if res.User_id == id {
		service.AddBondToUserInDB(id, bond)
	} else {
		fmt.Println("need to login")
	}

}

func (method *ApiMethodsService) Delete(c *gin.Context) {
	service := mongo.MongoService{}
	deleteMap := make(map[string]string)

	err := c.BindJSON(&deleteMap)
	if err != nil {
		fmt.Println("err: ", err)
	}

	var name string
	for _, value := range deleteMap {
		name = string(value)
	}
	service.DeleteBond(name)
}

func (method *ApiMethodsService) Login(c *gin.Context) {
	hashService := hash.HashService{}
	service := mongo.MongoService{}
	var userToken bondsInfo.UserToken
	loginMap := make(map[string]string)

	err := c.BindJSON(&loginMap)
	if err != nil {
		fmt.Println("err: ", err)
	}
	res := service.FindUserInDB(loginMap["login"])

	result := res.Login == loginMap["login"] && res.Password == hashService.Password(loginMap["password"])

	if result == true {
		userToken.Token, userToken.Refresh_token, _ = helper.GenerateAllTokens(res.Login, res.Name, res.User_type, res.User_id)
	}
	fmt.Println("result: ", result)
	c.JSON(http.StatusOK, gin.H{
		"response": result,
		"token":    userToken,
	})
}

var redisClient = redis.NewClient(&redis.Options{
	Addr: os.Getenv("REDIS_CONN"),
})

func (method *ApiMethodsService) Register(c *gin.Context) {
	mongoService := mongo.MongoService{}
	hashService := hash.HashService{}
	var user UserModel.User
	var userToken bondsInfo.UserToken

	registerMap := make(map[string]string)

	err := c.BindJSON(&registerMap)

	if err != nil {
		fmt.Println("err: ", err.Error())
	}

	name := registerMap["name"]
	login := registerMap["login"]
	password := hashService.Password(registerMap["password"])
	user.ID = mongoService.GenerateId()
	user.User_type = "User"
	user.User_id = user.ID.Hex()
	userToken.Token, userToken.Refresh_token, _ = helper.GenerateAllTokens(login, name, user.User_type, user.User_id)

	var userInfo = UserModel.User{
		ID:        user.ID,
		Name:      name,
		Password:  password,
		Login:     login,
		User_id:   user.User_id,
		User_type: user.User_type,
		Bonds:     []bondsInfo.UserBonds{},
	}

	payload, err := json.Marshal(gin.H{
		"message": userInfo.Login,
	})

	if err != nil {
		fmt.Println("err: ", err.Error())
	}
	if err := redisClient.Publish(c, "sendMail", payload).Err(); err != nil {
		panic(err)
	}

	res := mongoService.AddUserToDB(userInfo)
	c.JSON(http.StatusOK, gin.H{
		"token": userToken,
		"res":   res,
	})
}
