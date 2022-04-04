package routes

import (
    "gin-mongo-api/controllers" //add this
    "github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine)  {
    router.POST("/user", controllers.CreateUser())
    router.GET("/user/:userId", controllers.GetAUser()) //add this
    router.PUT("/user/:userId", controllers.EditAUser()) //add this
    router.DELETE("/user/:userId", controllers.DeleteAUser()) //add this
    router.GET("/users", controllers.GetAllUsers())
}



func AuthRoute(router *gin.Engine)  {
    router.POST("/authentication", controllers.Authentication())
}