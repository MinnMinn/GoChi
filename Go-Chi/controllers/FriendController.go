package controllers

import (
	"Go-Chi/models"
	"Go-Chi/services"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

func AddFriend(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var friends models.Friends
	json.NewDecoder(r.Body).Decode(&friends)
	var block = models.Request{Requestor: friends.Friends[0], Target:friends.Friends[1]}
	var beBlock = models.Request{Requestor: friends.Friends[1], Target:friends.Friends[0]}

	var checkNonAddFriend = services.CheckNonAddFriend(friends)
	var checkNonBlock = services.CheckNonBlock(block)
	var checkNonBeBlock = services.CheckNonBlock(beBlock)

	if checkNonAddFriend && checkNonBlock && checkNonBeBlock{
		services.AddFriend(friends)
		respondwithJSON(w, http.StatusCreated, map[string]bool{"success": true})
	} else {
		respondwithJSON(w, http.StatusInternalServerError, map[string]string{"message": "failed add friend"})
	}
}

func FindFriendsOfUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var mail models.Email
	json.NewDecoder(r.Body).Decode(&mail)
	var status = models.Friends{Success: true,Friends: services.FindFriendsOfUser(mail), Count: len(services.FindFriendsOfUser(mail))}
	if len(status.Friends) >0 {
		respondwithJSON(w, http.StatusOK, status)
	} else {
		respondwithJSON(w, http.StatusNotFound, map[string]string{"message": "Not Found Friends Of User"})
	}
}

func FindCommonFriends(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var friends models.Friends
	json.NewDecoder(r.Body).Decode(&friends)
	var status = models.Friends{Success: true, Friends: services.FindCommonFriends(friends), Count: len(services.FindCommonFriends(friends))}
	if len(status.Friends) >0 {
		respondwithJSON(w, http.StatusOK, status)
	} else {
		respondwithJSON(w, http.StatusNotFound, map[string]string{"message": "Not Found Common Friends"})
	}
}

func FollowFriend(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var subscribe models.Request
	json.NewDecoder(r.Body).Decode(&subscribe)

	var checkNonFollow = services.CheckNonFollow(subscribe)

	if checkNonFollow{
		services.FollowFriend(subscribe)
		respondwithJSON(w, http.StatusCreated, map[string]bool{"success": true})
	} else {
		respondwithJSON(w, http.StatusInternalServerError, map[string]string{"message": "failed follow"})
	}
}

func BlockFriend(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var block models.Request
	json.NewDecoder(r.Body).Decode(&block)

	var checkNonBlock = services.CheckNonBlock(block)
	if checkNonBlock {
		services.BlockFriend(block)
		respondwithJSON(w, http.StatusCreated, map[string]bool{"success": true})
	} else {
		respondwithJSON(w, http.StatusInternalServerError, map[string]string{"failed": "You were blocked this account !!!"})
	}
}

func ReceiveUpdatesFromEmail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var sender models.Sender
	json.NewDecoder(r.Body).Decode(&sender)
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

	respondwithJSON(w, http.StatusCreated, models.Recipients{Success: len(receiveUpdates) > 0, Recipients:receiveUpdates})
}

func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	fmt.Println(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}