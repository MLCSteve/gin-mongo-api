package main

import (
	"github.com/MLCSteve/gin-mongo-api.git/configs"
	"github.com/MLCSteve/gin-mongo-api.git/routes" //add this
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//run database
	configs.ConnectDB()

	//routes
	routes.UserRoute(router) //add this

	router.Run("localhost:6000")
}
