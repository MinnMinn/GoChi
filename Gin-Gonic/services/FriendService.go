package services

import (
	"Gin-Gonic/driver"
	"Gin-Gonic/models"
	"Gin-Gonic/repositories"
)

func CheckNonAddFriend(friends models.Friends) bool {
	db := driver.DBConn()
	connectRepository := repositories.FriendRepository(db)
	return  connectRepository.CheckNonAddFriend(friends)
}

func AddUser(email models.Email) error {
	db := driver.DBConn()
	connectRepository := repositories.FriendRepository(db)
	return  connectRepository.AddUser(email)
}

func AddFriend(friends models.Friends) error {
	db := driver.DBConn()
	connectRepository := repositories.FriendRepository(db)
	return  connectRepository.AddFriend(friends)
}

func FindFriendsOfUser(m models.Email) []string {
	db := driver.DBConn()
	connectRepository := repositories.FriendRepository(db)
	return  connectRepository.FindFriendsOfUser(m)
}

func FindCommonFriends(friends models.Friends) []string {
	db := driver.DBConn()
	connectRepository := repositories.FriendRepository(db)
	return  connectRepository.FindCommonFriends(friends)
}

func CheckNonFollow(subscribe models.Request) bool {
	db := driver.DBConn()
	connectRepository := repositories.FriendRepository(db)
	return  connectRepository.CheckNonFollow(subscribe)
}

func FollowFriend(subscribe models.Request) error {
	db := driver.DBConn()
	connectRepository := repositories.FriendRepository(db)
	return  connectRepository.FollowFriend(subscribe)
}

func CheckNonBlock(subscribe models.Request) bool {
	db := driver.DBConn()
	connectRepository := repositories.FriendRepository(db)
	return  connectRepository.CheckNonBlock(subscribe)
}

func BlockFriend(subscribe models.Request) error {
	db := driver.DBConn()
	connectRepository := repositories.FriendRepository(db)
	return  connectRepository.BlockFriend(subscribe)
}

func NonBlockByEmail(sender models.Sender) []string {
	db := driver.DBConn()
	connectRepository := repositories.FriendRepository(db)
	return connectRepository.NonBlockByEmail(sender)
}
