package interfaces

import (
	"Gin-Gonic/models"
)

type ICheck interface {
	CheckNonAddFriend(friends models.Friends) bool
	CheckNonFollow(subscribe models.Request) bool
	CheckNonBlock(subscribe models.Request) bool
}
