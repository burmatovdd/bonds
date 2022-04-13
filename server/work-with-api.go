package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
)

func yearPost(c *gin.Context) {

	yearMap := make(map[string]string)

	err := c.BindJSON(&yearMap)
	if err != nil {
		fmt.Println("err: ", err)
	}

	var year string
	for _, value := range yearMap {
		year = string(value)
	}

	test2 := TakeData(year)
	c.JSON(http.StatusOK, struct {
		AllInfos AllInfo `json:"allInfos"`
	}{
		test2.Coupons,
	})
}

func bondsPost(c *gin.Context) {
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println("err: ", err)
	}
	//fmt.Println(string(jsonData))

	bondsMap := make(map[string]string)
	err = json.Unmarshal(jsonData, &bondsMap)
	if err != nil {
		fmt.Println("err: ", err)
	}

	//fmt.Println("bondsMap: ", bondsMap)
	var bond string
	var count float64
	for i := 1; i < len(bondsMap); i++ {
		bond = bondsMap["bond"]
		count, _ = strconv.ParseFloat(bondsMap["count"], 64)
		bonds = append(bonds, Bond{Name: bond, Count: count})
	}

	for i := 0; i < len(bonds); i++ {
		var bondInfo = Bond{
			bonds[i].Name,
			bonds[i].Count,
		}
		addToDB(bondInfo)
		//deleteBond(bondInfo.Name)
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
	loginMap := make(map[string]string)

	err := c.BindJSON(&loginMap)
	if err != nil {
		fmt.Println("err: ", err)
	}

	var login string
	var password string
	login = loginMap["login"]
	password = loginMap["password"]
	fmt.Println("login: ", login)
	fmt.Println("password: ", password)
}

func register(c *gin.Context) {
	registerMap := make(map[string]string)

	err := c.BindJSON(&registerMap)
	if err != nil {
		fmt.Println("err: ", err)
	}

	var name string
	var login string
	var password string
	name = registerMap["name"]
	login = registerMap["login"]
	password = registerMap["password"]
	fmt.Println("name: ", name)
	fmt.Println("login: ", login)
	fmt.Println("password: ", password)
}

func HandleRequest() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"Content-Type,access-control-allow-origin, access-control-allow-headers"},
	}))
	router.POST("/year", yearPost)
	router.POST("/bonds", bondsPost)
	router.POST("/delete", delete)
	router.POST("/login", loginUser)
	router.POST("/register", register)
	err := router.Run(":8080")
	if err != nil {
		fmt.Println("err: ", err)
	}
}
