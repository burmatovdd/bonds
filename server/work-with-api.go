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

	bondsMap := make(map[string]string)
	err = json.Unmarshal(jsonData, &bondsMap)
	if err != nil {
		fmt.Println("err: ", err)
	}

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
		addBondToDB(bondInfo)
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
	users := getAllUsers()
	//cap - длина для срезов
	if cap(users) == 0 {
		c.JSON(http.StatusOK, struct {
			Response Response `json:"response"`
		}{
			Response{
				false,
			},
		})
	}

	for i := 0; i < len(users); i++ {
		if users[i].Login == login {
			fmt.Println("find login")
			if users[i].Password == password {
				fmt.Println("find password")
				c.JSON(http.StatusOK, struct {
					Response Response `json:"response"`
				}{
					Response{
						true,
					},
				})
				break
			} else {
				c.JSON(http.StatusOK, struct {
					Response Response `json:"response"`
				}{
					Response{
						false,
					},
				})
				break
			}
		} else {
			c.JSON(http.StatusOK, struct {
				Response Response `json:"response"`
			}{
				Response{
					false,
				},
			})
			break
		}
	}
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
	var userInfo = User{
		name,
		login,
		password,
		[]BondInfo{},
	}
	addUserToDB(userInfo)
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
