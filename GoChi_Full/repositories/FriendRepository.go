package repositories

import (
	"SP_FriendManagement_API_TungNguyen/models"
	"database/sql"
)

type Database struct {
	Connect *sql.DB
}

func (db Database) GetAllEmail() []string {
	email, err := db.Connect.Query("SELECT `email` FROM `user`")
	catch(err)
	var emails []string
	for email.Next(){
		var e string
		err = email.Scan(&e)
		catch(err)
		emails = append(emails, e)
	}
	return emails
}

func (db Database) AddUser(email models.Email) bool {
	addUser, err := db.Connect.Prepare("INSERT INTO `user` (`email`) VALUES (?)")
	catch(err)
	_, err = addUser.Exec(email.Email)
	catch(err)
	defer addUser.Close()
	return true
}

func (db Database) AddFriend(friends models.Friends) bool{
	addFriend, err := db.Connect.Prepare("INSERT `connection` SET user_id=(SELECT `id` FROM `user` WHERE `email`=?), connect_id=(SELECT `id` FROM `user` WHERE `email`=?)")
	catch(err)
	_, err = addFriend.Exec(friends.Friends[0], friends.Friends[1])
	catch(err)
	defer addFriend.Close()
	return true
}

func (db Database) FindFriendsOfUser(m models.Email) []string {
	emailFriends, err := db.Connect.Query("SELECT `email` FROM `user` WHERE `id` IN (SELECT `connect_id` FROM `connection` WHERE `user_id` = (SELECT `id` FROM `user` WHERE `email`=?))", m.Email)
	catch(err)
	var email []string
	for emailFriends.Next(){
		var e string
		err = emailFriends.Scan(&e)
		catch(err)
		email = append(email, e)
	}
	return email
}

func (db Database) FindCommonFriends(friends models.Friends)[]string{
	commonFriends, err := db.Connect.Query("SELECT `email` FROM `user` WHERE `id` IN (" +
		"SELECT `user_id` FROM `connection` JOIN (" +
		"SELECT `id` FROM `user` WHERE `email` IN ( ?, ?)) t " +
		"ON `connect_id` = `id` GROUP BY `user_id` HAVING COUNT(`user_id`) > 1)", friends.Friends[0], friends.Friends[1])
	catch(err)
	var email []string
	for commonFriends.Next(){
		var e string
		commonFriends.Scan(&e)
		email = append(email, e)
	}
	return email
}

func (db Database) SubscribeFriend(subscribe models.Request) bool {
	subscribeUser, err := db.Connect.Prepare("INSERT `follow` SET `user_id`=(SELECT `id` FROM `user` WHERE `email`=?), follow_id=(SELECT `id` FROM `user` WHERE `email`=?)")
	catch(err)
	_, err = subscribeUser.Exec(subscribe.Requestor, subscribe.Target)
	catch(err)
	defer subscribeUser.Close()
	return true
}

func (db Database) BlockFriend(subscribe models.Request) bool {
	blockUser, err := db.Connect.Prepare("INSERT `block` SET `user_id`=(SELECT `id` FROM `user` WHERE `email`=?), block_id=(SELECT `id` FROM `user` WHERE `email`=?)")
	catch(err)
	_, err = blockUser.Exec(subscribe.Requestor, subscribe.Target)
	catch(err)
	defer blockUser.Close()
	return true
}

func (db Database) RegisteredCheck(email models.Email) bool {
	user, err := db.Connect.Query("SELECT `id` FROM `user` WHERE `email` = ?", email.Email)
	catch(err)

	for user.Next(){
		return true
	}
	defer user.Close()
	return false
}

func (db Database) FriendedCheck(friends models.Friends) bool {
	connect, err := db.Connect.Query("SELECT `user_id` FROM `connection` WHERE `user_id` = (SELECT `id` FROM `user` WHERE `email`=?) AND `connect_id` = (SELECT `id` FROM `user` WHERE `email`=?)", friends.Friends[0], friends.Friends[1])
	catch(err)

	for connect.Next(){
		return true
	}

	defer connect.Close()
	return false
}

func (db Database) SubscribeFriendCheck(req models.Request) bool {
	subscribe, err := db.Connect.Query("SELECT `user_id` FROM `follow` WHERE `user_id` = (SELECT `id` FROM `user` WHERE `email`=?) AND `follow_id` = (SELECT `id` FROM `user` WHERE `email`=?)", req.Requestor, req.Target)
	catch(err)

	for subscribe.Next(){
		return true
	}

	defer subscribe.Close()
	return false
}

func (db Database) BlockFriendCheck(subscribe models.Request) bool {
	block, err := db.Connect.Query("SELECT `user_id` FROM `block` WHERE `user_id` = (" +
		"SELECT `id` FROM `user` WHERE `email`=?) AND `block_id` = (" +
		"SELECT `id` FROM `user` WHERE `email`=?)", subscribe.Requestor, subscribe.Target)
	catch(err)

	for block.Next(){
		return true
	}

	defer block.Close()
	return false
}

func (db Database) NonBlockByEmail(sender models.Sender) []string {
	nonBlockId, err := db.Connect.Query("SELECT `email` FROM `user` WHERE `id` NOT IN (" +
		"SELECT `block_id` FROM `block` JOIN( " +
		"SELECT `id` FROM `user` WHERE `email` = ?) `u` ON `user_id` = `u`.`id`)", sender.Sender)
	catch(err)
	var emails []string
	for nonBlockId.Next() {
		var email string
		err = nonBlockId.Scan(&email)
		catch(err)
		emails = append(emails, email)
	}
	return emails
}

func catch(err error) {
	if err != nil {
		panic(err)
	}
}