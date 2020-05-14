package interfaces

import (
	"Gin-Gonic/models"
)

type IRepository interface {
	AddUser(email models.Email) error
	AddFriend(friends models.Friends) error
	FindFriendsOfUser(m models.Email) []string
	FindCommonFriends(friends models.Friends)[]string
	FollowFriend(subscribe models.Request) error
	BlockFriend(subscribe models.Request) error
	NonBlockByEmail(sender models.Sender) []string
	ICheck
}
