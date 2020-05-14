package interfaces

import (
	"SP_FriendManagement_API_TungNguyen/models"
)

type ICheck interface {
	RegisteredCheck(email models.Email) bool
	FriendedCheck(friends models.Friends) bool
	SubscribeFriendCheck(subscribe models.Request) bool
	BlockFriendCheck(subscribe models.Request) bool
}
