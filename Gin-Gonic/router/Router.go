package router

import (
	"Gin-Gonic/controllers"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Static("/public", "./public")

	router := r.Group("/api")
	{
		router.POST("/addUser", controllers.AddUser)
		router.POST("/addFriend", controllers.AddFriend)
		router.GET("/findFriendOfUser", controllers.FindFriendsOfUser)
		router.GET("/findCommonFriends", controllers.FindCommonFriends)
		router.POST("/followFriend", controllers.FollowFriend)
		router.POST("/blockFriend", controllers.BlockFriend)
		router.GET("/receiveUpdatesFromEmail", controllers.ReceiveUpdatesFromEmail)
	}

	return r
}