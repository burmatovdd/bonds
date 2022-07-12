package api

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	apiMethods "modules/internal/server/api/apiMethods"
	"modules/internal/server/api/middleware"
)

type ApiStartService struct {
	method *ApiStart
}

type ApiStart interface {
	HandleRequest()
}

func (method *ApiStartService) HandleRequest() {
	service := apiMethods.ApiMethodsService{}
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"Content-Type,access-control-allow-origin, access-control-allow-headers, Authorization"},
	}))
	router.POST("/api/login", service.Login)
	router.POST("/api/register", service.Register)
	router.Use(middleware.Middleware())
	router.POST("/api/year", service.YearPost)
	router.POST("/api/bonds", service.BondsPost)
	router.POST("/api/delete", service.Delete)
	err := router.Run(":8080")
	if err != nil {
		fmt.Println("err: ", err)
	}
}
