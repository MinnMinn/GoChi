package interfaces

import (
	"Go-Chi/models"
)

type IRepository interface {
	AddFriend(friends models.Friends) error
	FindFriendsOfUser(m models.Email) []string
	FindCommonFriends(friends models.Friends)[]string
	FollowFriend(subscribe models.Request) error
	BlockFriend(subscribe models.Request) error
	NonBlockByEmail(sender models.Sender) []string
	ICheck
}
