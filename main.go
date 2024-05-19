package main

import (
	"github.com/UjjwalMahar/llamakraft/initializers"
	"github.com/gin-gonic/gin"
	"github.com/UjjwalMahar/llamakraft/controllers"
	"github.com/UjjwalMahar/llamakraft/middlewares"
)


func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
	initializers.SyncDatabase()
}

func main() {

	r := gin.Default()
	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middlewares.RequireAuth , controllers.Validate)
	r.GET("/health", controllers.Health)
	r.Run()
}