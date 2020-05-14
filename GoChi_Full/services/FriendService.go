package services

import (
	"SP_FriendManagement_API_TungNguyen/interfaces"
	"SP_FriendManagement_API_TungNguyen/models"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

type FriendsService struct {
	API interfaces.IRepository
}

func (service FriendsService) GetAllEmail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	emails := service.API.GetAllEmail()
	if len(emails) == 0 {
		respondwithJSON(w, http.StatusCreated, "Emails exists none")
		return
	}
	respondwithJSON(w, http.StatusCreated, emails)
}

func (service FriendsService) AddUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var email models.Email
	json.NewDecoder(r.Body).Decode(&email)
	var regexpMail, _ = regexp.MatchString("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$", email.Email)
	if !regexpMail {
		respondwithJSON(w, http.StatusCreated, "Format Email Is Not True !!!")
		return
	}
	var checkExistUser = service.API.RegisteredCheck(email)
	if !checkExistUser {
		service.API.AddUser(email)
		respondwithJSON(w, http.StatusCreated, map[string]bool{"success": true})
		return
	} else {
		respondwithJSON(w, http.StatusInternalServerError, map[string]string{"error" : "User Exist Already !!!"})
		return
	}
}

func (service FriendsService) AddFriend(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var friends models.Friends
	json.NewDecoder(r.Body).Decode(&friends)
	var email1 = models.Email{friends.Friends[0]}
	var email2 = models.Email{friends.Friends[1]}
	var checkExistUser1 = service.API.RegisteredCheck(email1)
	var checkExistUser2 = service.API.RegisteredCheck(email2)

	if !checkExistUser1 || !checkExistUser2 {
		respondwithJSON(w, http.StatusInternalServerError, map[string]string{"error" : "Not Found Email !!!"})
		return
	}
	var block = models.Request{Requestor: email1.Email, Target:email2.Email}
	var beBlock = models.Request{Requestor: friends.Friends[1], Target:friends.Friends[0]}

	var checkFriended = service.API.FriendedCheck(friends)
	var checkBlockFriend = service.API.BlockFriendCheck(block)
	var checkBeBlockFriend = service.API.BlockFriendCheck(beBlock)
	if !checkFriended && !checkBlockFriend && !checkBeBlockFriend{
		beFriend := models.Friends{Success: true, Friends: []string{friends.Friends[1], friends.Friends[0]}, Count: 1}
		service.API.AddFriend(friends)
		service.API.AddFriend(beFriend)
		respondwithJSON(w, http.StatusCreated, map[string]bool{"success": true})
		return
	} else {
		respondwithJSON(w, http.StatusInternalServerError, map[string]string{"error" : "You Were Friended Or Be Blocked Each Other !!!"})
		return
	}
}

func (service FriendsService) FindFriendsOfUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var email models.Email
	json.NewDecoder(r.Body).Decode(&email)
	var checkExistUser = service.API.RegisteredCheck(email)
	if !checkExistUser {
		respondwithJSON(w, http.StatusInternalServerError, map[string]string{"error" : "Not Found Email !!!"})
		return
	}
	var status = models.Friends{Success: true,Friends: service.API.FindFriendsOfUser(email), Count: len(service.API.FindFriendsOfUser(email))}
	if len(status.Friends) > 0 {
		respondwithJSON(w, http.StatusOK, status)
		return
	} else {
		respondwithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Not Found Friends Of User"})
		return
	}
}

func (service FriendsService) FindCommonFriends(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var friends models.Friends
	json.NewDecoder(r.Body).Decode(&friends)
	var email1 = models.Email{friends.Friends[0]}
	var email2 = models.Email{friends.Friends[1]}
	var checkExistUser1 = service.API.RegisteredCheck(email1)
	var checkExistUser2 = service.API.RegisteredCheck(email2)

	if !checkExistUser1 || !checkExistUser2 {
		respondwithJSON(w, http.StatusInternalServerError, map[string]string{"error" : "Not Found Email !!!"})
		return
	}
	var status = models.Friends{Success: true, Friends: service.API.FindCommonFriends(friends), Count: len(service.API.FindCommonFriends(friends))}
	if len(status.Friends) >0 {
		respondwithJSON(w, http.StatusOK, status)
		return
	} else {
		respondwithJSON(w, http.StatusInternalServerError, map[string]string{"error" : "Not Found Common Friends !!!"})
		return
	}
}

func (service FriendsService) SubscribeFriend(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var subscribe models.Request
	json.NewDecoder(r.Body).Decode(&subscribe)
	var email1 = models.Email{subscribe.Requestor}
	var email2 = models.Email{subscribe.Target}
	var checkExistUser1 = service.API.RegisteredCheck(email1)
	var checkExistUser2 = service.API.RegisteredCheck(email2)

	if !checkExistUser1 || !checkExistUser2 {
		respondwithJSON(w, http.StatusInternalServerError, map[string]string{"error" : "Not Found Email !!!"})
		return
	}
	var checkFollowFriend = service.API.SubscribeFriendCheck(subscribe)
	if !checkFollowFriend{
		service.API.SubscribeFriend(subscribe)
		respondwithJSON(w, http.StatusCreated, map[string]bool{"success": true})
		return
	} else {
		respondwithJSON(w, http.StatusInternalServerError, map[string]string{"error" : "You Were Followed Or Be Blocked !!!"})
		return
	}
}

func (service FriendsService) BlockFriend(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var block models.Request
	json.NewDecoder(r.Body).Decode(&block)
	var email1 = models.Email{block.Requestor}
	var email2 = models.Email{block.Target}
	var checkExistUser1 = service.API.RegisteredCheck(email1)
	var checkExistUser2 = service.API.RegisteredCheck(email2)

	if !checkExistUser1 || !checkExistUser2 {
		respondwithJSON(w, http.StatusInternalServerError, map[string]string{"error" : "Not Found Email !!!"})
		return
	}
	var checkBlock = service.API.BlockFriendCheck(block)
	if !checkBlock {
		service.API.BlockFriend(block)
		respondwithJSON(w, http.StatusCreated, map[string]bool{"success": true})
		return
	} else {
		respondwithJSON(w, http.StatusInternalServerError, map[string]string{"error" : "Friends Blocked !!!"})
		return

	}
}

func (service FriendsService) ReceiveUpdatesFromEmail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var sender models.Sender
	json.NewDecoder(r.Body).Decode(&sender)
	var receiveUpdates []string
	var email = models.Email{sender.Sender}
	var checkExistUser = service.API.RegisteredCheck(email)

	if !checkExistUser {
		respondwithJSON(w, http.StatusInternalServerError, map[string]string{"error" : "Not Found Email !!!"})
		return
	}
	var emails = service.API.NonBlockByEmail(sender)
	for i := 0; i < len(emails); i++ {
		var friends = models.Friends{Friends:[]string{sender.Sender, emails[i]}}
		var subscribe = models.Request{Requestor:emails[i], Target: sender.Sender}
		var checkFriended = service.API.FriendedCheck(friends)
		var checkFollowFriend = service.API.SubscribeFriendCheck(subscribe)
		if checkFriended || checkFollowFriend {
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

	if len(receiveUpdates) > 0 {
		respondwithJSON(w, http.StatusCreated, models.Recipients{Success: len(receiveUpdates) > 0, Recipients:receiveUpdates})
		return
	} else {
		respondwithJSON(w, http.StatusInternalServerError, map[string]string{"error": "User is not exits"})
		return
	}
}

func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	fmt.Println(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}