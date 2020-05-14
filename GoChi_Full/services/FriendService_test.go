package services

import (
	"SP_FriendManagement_API_TungNguyen/models"
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

type APIMock struct {
	mock.Mock
}

func (mock *APIMock) GetAllEmail() []string {
	args := mock.Called()
	t := args.Get(0).([]string)
	return t
}

func (mock *APIMock) AddUser(email models.Email) bool {
	args := mock.Called(email.Email)
	return args.Bool(0)
}

func (mock *APIMock) AddFriend(friends models.Friends) bool{
	args := mock.Called(friends.Friends[0], friends.Friends[1])
	return args.Bool(0)
}

func (mock *APIMock) FindFriendsOfUser(email models.Email) []string {
	args := mock.Called(email.Email)
	t := args.Get(0).([]string)
	return t
}

func (mock *APIMock) FindCommonFriends(friends models.Friends)[]string{
	args := mock.Called(friends.Friends[0], friends.Friends[1])
	t := args.Get(0).([]string)
	return t
}

func (mock *APIMock) SubscribeFriend(subscribe models.Request) bool {
	args := mock.Called(subscribe.Requestor, subscribe.Target)
	return args.Bool(0)
}

func (mock *APIMock) BlockFriend(subscribe models.Request) bool {
	args := mock.Called(subscribe.Requestor, subscribe.Target)
	return args.Bool(0)
}

func (mock *APIMock) RegisteredCheck(email models.Email) bool {
	args := mock.Called(email.Email)
	return args.Bool(0)
}

func (mock *APIMock) FriendedCheck(friends models.Friends) bool {
	args := mock.Called(friends.Friends[0], friends.Friends[1])
	return args.Bool(0)
}

func (mock *APIMock) SubscribeFriendCheck(subscribe models.Request) bool {
	args := mock.Called(subscribe.Requestor, subscribe.Target)
	return args.Bool(0)
}

func (mock *APIMock) BlockFriendCheck(subscribe models.Request) bool {
	args := mock.Called(subscribe.Requestor, subscribe.Target)
	return args.Bool(0)
}

func (mock *APIMock) NonBlockByEmail(sender models.Sender) []string {
	args := mock.Called(sender.Sender, sender.Text)
	t := args.Get(0).([]string)
	return t
}

func TestAddFriend(t *testing.T) {
	var jsonStr = []byte(`{"friends": ["andy@example.com","john@example.com"]}`)
	req, err := http.NewRequest("POST", "/api/addFriend", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	API := new(APIMock)
	API.On("RegisteredCheck", "andy@example.com").Return(true)
	API.On("RegisteredCheck", "john@example.com").Return(true)
	API.On("FriendedCheck", "andy@example.com","john@example.com").Return(false)
	API.On("BlockFriendCheck", "andy@example.com","john@example.com").Return(false)
	API.On("BlockFriendCheck", "john@example.com","andy@example.com").Return(false)
	API.On("AddFriend", "andy@example.com","john@example.com").Return(true)
	API.On("AddFriend", "john@example.com","andy@example.com").Return(true)
	service := FriendsService{API}
	service.AddFriend(w, req)
	var res models.Success
	errPs := json.NewDecoder(w.Body).Decode(&res)
	assert.NoError(t, errPs)
	assert.Equal(t, res, models.Success{true})
}

func TestFindFriendsOfUser(t *testing.T) {
	var jsonStr = []byte(`{"email": "andy@example.com"}`)
	friends := []string {"john@example.com", "common@example.com"}
	req, err := http.NewRequest("GET", "/api/findFriendOfUser", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	API := new(APIMock)
	API.On("RegisteredCheck", "andy@example.com").Return(true)
	API.On("FindFriendsOfUser", "andy@example.com").Return(friends)
	service := FriendsService{API}
	service.FindFriendsOfUser(w, req)
	var res models.Friends
	errPs := json.NewDecoder(w.Body).Decode(&res)
	assert.NoError(t, errPs)
	ExpectedRe := models.Friends{true, friends, len(friends)}
	assert.Equal(t, res, ExpectedRe)
}

func TestFindCommonFriends(t *testing.T) {
	var jsonStr = []byte(`{"friends": ["andy@example.com","john@example.com"]}`)
	friends := models.Friends{
		Success: true,
		Friends : []string{"common@example.com"},
		Count : 1,
	}

	req, err := http.NewRequest("GET", "/api/findCommonFriends", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	API := new(APIMock)
	API.On("RegisteredCheck", "andy@example.com").Return(true)
	API.On("RegisteredCheck", "john@example.com").Return(true)
	API.On("FindCommonFriends", "andy@example.com","john@example.com").Return(friends.Friends)
	service := FriendsService{API}
	service.FindCommonFriends(w, req)
	var res models.Success
	errPs := json.NewDecoder(w.Body).Decode(&res)
	assert.NoError(t, errPs)
	assert.Equal(t, res, models.Success{true})
}

func TestSubscribeFriend(t *testing.T) {
	var jsonStr = []byte(`{"requestor": "andy@example.com", "target": "common@example.com"}`)
	req, err := http.NewRequest("POST", "/api/subscribeFriend", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	API := new(APIMock)
	API.On("RegisteredCheck", "andy@example.com").Return(true)
	API.On("RegisteredCheck", "common@example.com").Return(true)
	API.On("SubscribeFriendCheck", "andy@example.com","common@example.com").Return(false)
	API.On("SubscribeFriend", "andy@example.com","common@example.com").Return(false)

	service := FriendsService{API}
	service.SubscribeFriend(w, req)
	var res models.Success
	errPs := json.NewDecoder(w.Body).Decode(&res)
	assert.NoError(t, errPs)
	assert.Equal(t, res, models.Success{true})
}

func TestBlockFriend(t *testing.T) {
	var jsonStr = []byte(`{"requestor": "andy@example.com", "target": "john@example.com"}`)
	req, err := http.NewRequest("POST", "/api/blockFriend", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	API := new(APIMock)
	API.On("RegisteredCheck", "andy@example.com").Return(true)
	API.On("RegisteredCheck", "john@example.com").Return(true)
	API.On("BlockFriendCheck", "andy@example.com","john@example.com").Return(false)
	API.On("BlockFriend", "andy@example.com","john@example.com").Return(true)

	service := FriendsService{API}
	service.BlockFriend(w, req)
	var res models.Success
	errPs := json.NewDecoder(w.Body).Decode(&res)
	assert.NoError(t, errPs)
	assert.Equal(t, res, models.Success{true})
}

func TestReceiveUpdatesFromEmail(t *testing.T) {
	var jsonStr = []byte(`{"sender": "john@example.com", "text": "Hello World! kate@example.com"}`)
	recipients := models.Recipients{
		Success: true,
		Recipients:[]string{
		"lisa@example.com",
		"kate@example.com",
		},
	}
	req, err := http.NewRequest("POST", "/api/receiveUpdatesFromEmail", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	API := new(APIMock)
	API.On("RegisteredCheck", "john@example.com").Return(true)
	API.On("NonBlockByEmail", "john@example.com","Hello World! kate@example.com").Return(recipients.Recipients)
	API.On("FriendedCheck", "john@example.com","lisa@example.com").Return(true)
	API.On("SubscribeFriendCheck", "lisa@example.com","john@example.com").Return(true)
	API.On("FriendedCheck", "john@example.com","kate@example.com").Return(true)
	API.On("SubscribeFriendCheck", "kate@example.com","john@example.com").Return(true)

	service := FriendsService{API}
	service.ReceiveUpdatesFromEmail(w, req)
	var res models.Success
	errPs := json.NewDecoder(w.Body).Decode(&res)
	assert.NoError(t, errPs)
	assert.Equal(t, res, models.Success{true})
}

