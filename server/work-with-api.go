package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"io/ioutil"
	helper "modules/helpers"
	"modules/middleware"
	"net/http"
	"os"
	"strconv"
)

func yearPost(c *gin.Context) {
	yearMap := make(map[string]string)

	err := c.BindJSON(&yearMap)
	if err != nil {
		fmt.Println("err: ", err)
	}

	id := c.GetString("uid")

	res := findUserInDBById(id)
	if res.User_id == id {
		test2 := TakeData(yearMap["year"], res.Bond)
		c.JSON(http.StatusOK, struct {
			AllInfos AllInfo `json:"allInfos"`
		}{
			test2.Coupons,
		})
	} else {
		fmt.Println("need to login")
	}

}

func bondsPost(c *gin.Context) {
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println("err: ", err)
	}

	bondsMap := make(map[string]string)
	err = json.Unmarshal(jsonData, &bondsMap)
	if err != nil {
		fmt.Println("err: ", err)
	}
	id := c.GetString("uid")

	res := findUserInDBById(id)
	if res.User_id == id {
		count, _ := strconv.ParseFloat(bondsMap["count"], 64)

		res.Bond = append(res.Bond, Bond{bondsMap["bond"], count})
		addBondToUserInDB(id, res.Bond)
	} else {
		fmt.Println("need to login")
	}

}

func delete(c *gin.Context) {
	deleteMap := make(map[string]string)

	err := c.BindJSON(&deleteMap)
	if err != nil {
		fmt.Println("err: ", err)
	}

	var name string
	for _, value := range deleteMap {
		name = string(value)
	}
	deleteBond(name)
}

func loginUser(c *gin.Context) {
	var userToken UserToken
	loginMap := make(map[string]string)

	err := c.BindJSON(&loginMap)
	if err != nil {
		fmt.Println("err: ", err)
	}

	res := findUserInDB(loginMap["login"])

	result := res.Login == loginMap["login"] && res.Password == hashPassword(loginMap["password"])

	if result == true {
		userToken.Token, userToken.Refresh_token, _ = helper.GenerateAllTokens(res.Login, res.Name, res.User_type, res.User_id)
		fmt.Println("login token: ", userToken.Token)
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

func register(c *gin.Context) {
	var user User
	var userToken UserToken

	registerMap := make(map[string]string)

	err := c.BindJSON(&registerMap)

	if err != nil {
		fmt.Println("err: ", err.Error())
	}

	name := registerMap["name"]
	login := registerMap["login"]
	password := hashPassword(registerMap["password"])
	user.ID = generateId()
	user.User_type = "User"
	user.User_id = user.ID.Hex()
	userToken.Token, userToken.Refresh_token, _ = helper.GenerateAllTokens(login, name, user.User_type, user.User_id)

	var userInfo = User{
		user.ID,
		name,
		password,
		login,
		user.User_id,
		user.User_type,
		[]Bond{},
	}

	payload, err := json.Marshal(gin.H{
		"message": userInfo.Login,
	})

	if err != nil {
		fmt.Println("err: ", err.Error())
	}
	fmt.Println("payload: ", string(payload))
	if err := redisClient.Publish(ctx, "sendMail", payload).Err(); err != nil {
		panic(err)
	}

	res := addUserToDB(userInfo)
	c.JSON(http.StatusOK, gin.H{
		"token": userToken,
		"res":   res,
	})
}
func HandleRequest() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"Content-Type,access-control-allow-origin, access-control-allow-headers, Authorization"},
	}))
	router.POST("/api/login", loginUser)
	router.POST("/api/register", register)
	router.Use(middleware.Middleware1())
	router.POST("/api/year", yearPost)
	router.POST("/api/bonds", bondsPost)
	router.POST("/api/delete", delete)
	err := router.Run(":8080")
	if err != nil {
		fmt.Println("err: ", err)
	}
}
