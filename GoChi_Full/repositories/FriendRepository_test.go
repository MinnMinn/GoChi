package repositories

import (
	"SP_FriendManagement_API_TungNguyen/models"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
)

func TestAddUser(t *testing.T) {
	email := models.Email{"minn@example.com"}
	db, mock, mErr := sqlmock.New()
	if mErr != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", mErr)
	}
	defer db.Close()
	mock.ExpectPrepare("INSERT INTO `user`").ExpectExec().WillReturnResult(sqlmock.NewResult(1, 1))

	Fdb := Database{db}
	res := Fdb.AddUser(email)
	assert.Equal(t, true, res)
}

func TestAddFriend(t *testing.T) {
	addFriend := models.Friends{Friends: []string{"john@example.com", "andy@example.com"}}
	db, mock, mErr := sqlmock.New()
	if mErr != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", mErr)
	}
	defer db.Close()
	mock.ExpectPrepare("INSERT `connection`").ExpectExec().WillReturnResult(sqlmock.NewResult(1, 1))

	Fdb := Database{db}
	res := Fdb.AddFriend(addFriend)
	assert.Equal(t, true, res)
}

func TestFindFriendsOfUser(t *testing.T) {
	email := models.Email{
		"andy@example.com",
	}
	var listUserEmailMock = []string{"john@example.com", "lisa@example.com"}
	UserEmailsTableMock := sqlmock.NewRows([]string{"email"}).AddRow("john@example.com").AddRow("lisa@example.com")
	db, mock, mErr := sqlmock.New()
	if mErr != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", mErr)
	}
	defer db.Close()
	mock.ExpectQuery("SELECT `email` FROM").WithArgs("andy@example.com").WillReturnRows(UserEmailsTableMock)
	Fdb := Database{db}
	res := Fdb.FindFriendsOfUser(email)
	assert.Equal(t, res, listUserEmailMock)
}

func TestFindCommonFriends(t *testing.T) {
	var listUserEmailMock = []string{"common@example.com"}
	friends := models.Friends{Friends: []string{"andy@example.com", "john@example.com"}}
	UserEmailsTableMock := sqlmock.NewRows([]string{"email"}).AddRow("common@example.com")
	db, mock, mErr := sqlmock.New()
	if mErr != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", mErr)
	}
	defer db.Close()
	mock.ExpectQuery("SELECT `email` FROM").WithArgs("andy@example.com", "john@example.com").WillReturnRows(UserEmailsTableMock)
	Fdb := Database{db}
	res := Fdb.FindCommonFriends(friends)
	assert.Equal(t, res, listUserEmailMock)
}

func TestFollowFriend(t *testing.T) {
	friends := models.Request{Requestor: "lisa@example.com", Target: "john@example.com"}
	db, mock, mErr := sqlmock.New()
	if mErr != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", mErr)
	}
	defer db.Close()
	mock.ExpectPrepare("INSERT `follow`").ExpectExec().WillReturnResult(sqlmock.NewResult(1, 1))

	Fdb := Database{db}
	res := Fdb.SubscribeFriend(friends)
	assert.Equal(t, true, res)
}

func TestBlockFriend(t *testing.T) {
	request := models.Request{Requestor: "andy@example.com", Target: "john@example.com"}
	db, mock, mErr := sqlmock.New()
	if mErr != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", mErr)
	}
	defer db.Close()
	mock.ExpectPrepare("INSERT `block`").ExpectExec().WillReturnResult(sqlmock.NewResult(1, 1))

	Fdb := Database{db}
	res := Fdb.BlockFriend(request)
	assert.Equal(t, true, res)
}

func TestRegisteredCheck(t *testing.T) {
	var result = true
	email := models.Email{"minn@example.com"}
	UserIdTableMock := sqlmock.NewRows([]string{"id"}).AddRow(1)
	db, mock, mErr := sqlmock.New()
	if mErr != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", mErr)
	}
	defer db.Close()
	mock.ExpectQuery("SELECT `id` FROM").WithArgs("minn@example.com").WillReturnRows(UserIdTableMock)
	Fdb := Database{db}
	res := Fdb.RegisteredCheck(email)
	assert.Equal(t, res, result)
}

func TestFriendedCheck(t *testing.T) {
	var result = true
	friends := models.Friends{Friends: []string{"andy@example.com", "john@example.com"}}
	UserIdTableMock := sqlmock.NewRows([]string{"id"}).AddRow(3)
	db, mock, mErr := sqlmock.New()
	if mErr != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", mErr)
	}
	defer db.Close()
	mock.ExpectQuery("SELECT `user_id` FROM").WithArgs("andy@example.com", "john@example.com").WillReturnRows(UserIdTableMock)
	Fdb := Database{db}
	res := Fdb.FriendedCheck(friends)
	assert.Equal(t, res, result)
}

func TestFollowFriendCheck(t *testing.T) {
	var result = true
	request := models.Request{Requestor: "andy@example.com", Target: "john@example.com"}
	UserIdTableMock := sqlmock.NewRows([]string{"id"}).AddRow(3)
	db, mock, mErr := sqlmock.New()
	if mErr != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", mErr)
	}
	defer db.Close()
	mock.ExpectQuery("SELECT `user_id` FROM").WithArgs("andy@example.com", "john@example.com").WillReturnRows(UserIdTableMock)
	Fdb := Database{db}
	res := Fdb.SubscribeFriendCheck(request)
	assert.Equal(t, res, result)
}

func TestBlockFriendCheck(t *testing.T) {
	var result = true
	request := models.Request{Requestor: "andy@example.com", Target: "john@example.com"}
	UserIdTableMock := sqlmock.NewRows([]string{"id"}).AddRow(3)
	db, mock, mErr := sqlmock.New()
	if mErr != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", mErr)
	}
	defer db.Close()
	mock.ExpectQuery("SELECT `user_id` FROM").WithArgs("andy@example.com", "john@example.com").WillReturnRows(UserIdTableMock)
	Fdb := Database{db}
	res := Fdb.BlockFriendCheck(request)
	assert.Equal(t, res, result)
}

func TestCheckBlock(t *testing.T) {
	var result = true
	request := models.Request{Requestor: "andy@example.com", Target: "john@example.com"}
	UserIdTableMock := sqlmock.NewRows([]string{"id"}).AddRow(3)
	db, mock, mErr := sqlmock.New()
	if mErr != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", mErr)
	}
	defer db.Close()
	mock.ExpectQuery("SELECT `user_id` FROM").WithArgs("andy@example.com", "john@example.com").WillReturnRows(UserIdTableMock)
	Fdb := Database{db}
	res := Fdb.BlockFriendCheck(request)
	assert.Equal(t, res, result)
}

func TestNonBlockByEmail(t *testing.T) {
	result := []string{"andy@example.com", "common@example.com", "lisa@example.com"}
	request := models.Sender{Sender: "john@example.com"}
	UserEmailTableMock := sqlmock.NewRows([]string{"email"}).AddRow("andy@example.com").AddRow("common@example.com").AddRow("lisa@example.com")
	db, mock, mErr := sqlmock.New()
	if mErr != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", mErr)
	}
	defer db.Close()
	mock.ExpectQuery("SELECT `email` FROM").WithArgs("john@example.com").WillReturnRows(UserEmailTableMock)
	Fdb := Database{db}
	res := Fdb.NonBlockByEmail(request)
	assert.Equal(t, res, result)
}
