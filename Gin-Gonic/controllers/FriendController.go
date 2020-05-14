package controllers

import (
	"Gin-Gonic/models"
	"Gin-Gonic/services"
	"github.com/gin-gonic/gin"
	"regexp"
	"strings"
)

func AddUser(c * gin.Context) {
	var email models.Email
	c.ShouldBindJSON(&email)
	services.AddUser(email)
}

func AddFriend(c * gin.Context) {
	var friends models.Friends
	c.ShouldBindJSON(&friends)
	var block = models.Request{Requestor: friends.Friends[0], Target:friends.Friends[1]}
	var beBlock = models.Request{Requestor: friends.Friends[1], Target:friends.Friends[0]}

	var checkNonAddFriend = services.CheckNonAddFriend(friends)
	var checkNonBlock = services.CheckNonBlock(block)
	var checkNonBeBlock = services.CheckNonBlock(beBlock)
	success := models.Success{Success: checkNonAddFriend}
	if checkNonAddFriend && checkNonBlock && checkNonBeBlock{
		services.AddFriend(friends)
		c.JSON(200, success)
	} else {
		c.JSON(500, success)
	}
}

func FindFriendsOfUser(c * gin.Context) {
	var mail models.Email
	c.ShouldBindJSON(&mail)
	var status = models.Friends{Success: true,Friends: services.FindFriendsOfUser(mail), Count: len(services.FindFriendsOfUser(mail))}
	if len(status.Friends) >0 {
		c.JSON(200, status)
	} else {
		c.JSON(404, gin.H{
			"message" : "Not Found Friends Of User",
		})
	}
}

func FindCommonFriends(c * gin.Context) {
	var friends models.Friends
	c.ShouldBindJSON(&friends)
	var status = models.Friends{Success: true, Friends: services.FindCommonFriends(friends), Count: len(services.FindCommonFriends(friends))}
	if len(status.Friends) >0 {
		c.JSON(200, status)
	} else {
		c.JSON(404, gin.H{
			"message" : "Not Found Common Friends",
		})
	}
}

func FollowFriend(c * gin.Context) {
	var subscribe models.Request
	c.ShouldBindJSON(&subscribe)

	var checkNonFollow = services.CheckNonFollow(subscribe)
	success := models.Success{Success: checkNonFollow}
	if checkNonFollow{
		services.FollowFriend(subscribe)
		c.JSON(200, success)
	} else {
		c.JSON(500, gin.H{
			"message" : "You were followed",
		})
	}
}

func BlockFriend(c * gin.Context) {
	var block models.Request
	c.ShouldBindJSON(&block)

	var checkNonBlock = services.CheckNonBlock(block)
	success := models.Success{Success: checkNonBlock}
	if checkNonBlock {
		services.BlockFriend(block)
		c.JSON(200, success)
	} else {
		c.JSON(500, success)
	}
}

func ReceiveUpdatesFromEmail(c * gin.Context) {
	var sender models.Sender
	c.ShouldBindJSON(&sender)
	var receiveUpdates []string

	var emails = services.NonBlockByEmail(sender)
	for i := 0; i < len(emails); i++ {
		var friends = models.Friends{Friends:[]string{sender.Sender, emails[i]}}
		var subscribe = models.Request{Requestor:emails[i], Target: sender.Sender}
		var checkNonAddFriend = services.CheckNonAddFriend(friends)
		var checkNonFollow = services.CheckNonFollow(subscribe)
		if !checkNonAddFriend || !checkNonFollow {
			receiveUpdates = append(receiveUpdates, emails[i])
		}
	}

	var emailMentioned = strings.Split(sender.Text, " ")
	for a := range emailMentioned {
		var regexpMail, _ = regexp.MatchString("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$", emailMentioned[a])
		if regexpMail {
			receiveUpdates = append(receiveUpdates, emailMentioned[a])
		}
	}

	c.JSON(200, models.Recipients{Success: len(receiveUpdates) > 0, Recipients:receiveUpdates})
}