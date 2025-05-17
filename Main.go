package main

import (
	"github.com/NurochmanR/GO-JWT/controllers"
	"github.com/NurochmanR/GO-JWT/initializers"
	"github.com/NurochmanR/GO-JWT/middleware"
	"github.com/gin-gonic/gin"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.ConnetToDatabase()
}

func main() {
	router:=gin.Default()
	router.POST("/SignUp", controllers.SignUp)
	router.POST("/Login", controllers.Login)
	router.GET("/Validate", middleware.RequireAuth ,controllers.Validator)
	router.Run()
}
