package interfaces

import (
	"SP_FriendManagement_API_TungNguyen/models"
)

type IRepository interface {
	GetAllEmail() []string
	AddUser(email models.Email) bool
	AddFriend(friends models.Friends) bool
	FindFriendsOfUser(m models.Email) []string
	FindCommonFriends(friends models.Friends)[]string
	SubscribeFriend(subscribe models.Request) bool
	BlockFriend(subscribe models.Request) bool
	NonBlockByEmail(sender models.Sender) []string
	ICheck
}
