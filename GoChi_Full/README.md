# SP_FriendManagement_API_TungNguyen

## Installation & Run

#### Enviroment
This project uses phpMyadmin database inside with docker-compose. You can compose with dockerfile or create your own phpMyadmin database without it
You should set the database config with yours or leave it as default on [driver_mysql.go](https://github.com/s3corp-github/SP_FriendManagement_API_TungNguyen/tree/master/driver/driver_mysql.go)
```go
dbDriver := "mysql"
dbUser := "root"
dbPass := "root"
dbName := "golang"
dbPort := "3306"
dbHost := "fullstack-mysql"
```

For run docker-compose, run these following commands in project's root folder:

```bash
docker system prune --volumes
docker-compose up
```

#### phpMyadmin Database
```bash
# http://localhost:9090/
```

#### API Endpoint
```bash
# http://localhost:8081
```

## API Details

### 1. As a user, I need an API to create a friend connection between two email addresses.

Uri:
```
/addFriend
```
Method: POST

The API should receive the following JSON request:
```json
{
    "friends":
    [
        "andy@example.com",
        "john@example.com"
    ]
}
```
The Api response:
On success: HttpStatus = 200
```json
{
    "success": true
}
```
On error: HttpStatus = 400
```json
{
    "error" : "You Were Friended Or Be Blocked Each Other !!!"
}
```
### 2. As a user, I need an API to retrieve the friends list for an email address.
Uri:
```
/findFriendOfUser
```
The API should receive the following JSON request:
```json
{
    "email": "andy@example.com"
}
```
The Api response:
On success: HttpStatus = 200
```json
{
    "count": 2,
    "friends": [
        "john@example.com",
        "lisa@example.com"
    ],
    "success": true
}
```
On error: HttpStatus = 400
```json
{
    "error": "Not Found Friends Of andy@example.com"
}
```
### 3. As a user, I need an API to retrieve the common friends list between two email addresses.
Uri:
```
/findCommonFriends
```
The API should receive the following JSON request:
```json
{
    "friends": [
        "andy@example.com",
        "john@example.com"
    ]
}
```
The Api response:
On success: HttpStatus = 200
```json
{
    "success": true,
    "friends" :
    [
        "common@example.com"
    ],
    "count" : 1
}
```
On error: HttpStatus = 400
```json
{
    "error" : "Not Found Common Friends !!!"
}
```
### 4. As a user, I need an API to subscribe to updates from an email address.
Uri:
```
/subscribeFriend
```
The API should receive the following JSON request:
```json
{
    "requestor": "lisa@example.com",
    "target": "john@example.com"
}
```
On success: HttpStatus = 200
```json
{
    "success": true
}
```
On error: HttpStatus = 400
```json
{
    "error" : "You Were Followed Or Be Blocked !!!"
}
```
### 5. As a user, I need an API to block updates from an email address.
Uri:
```
/blockFriend
```

Suppose "andy@example.com" blocks "john@example.com":
The API should receive the following JSON request:
```json
{
    "requestor": "andy@example.com",
    "target": "john@example.com"
}
```
On success: HttpStatus = 200
```json
{
    "success": true
}
```
On error: HttpStatus = 400
```json
{
    "error": "Friends Blocked !!!"
}
```
### 6. As a user, I need an API to retrieve all email addresses that can receive updates from an email address.
Uri:
```
/receiveUpdatesFromEmail
```

The API should receive the following JSON request:
```json
{
    "sender": "john@example.com",
    "text": "Hello World! kate@example.com"
}
```
On success: HttpStatus = 200
```json
{
    "success": true,
    "recipients":
    [
        "andy@example.com",
        "common@example.com",
        "lisa@example.com",
        "kate@example.com"
    ]
}
```
On error: HttpStatus = 400
```json
{
    "error": "User is not exits",
    "success": false
}
```

### 7. As a user, I need an API to register my email account to the application Friend Management.
Uri:
```
/addUser
```

The API should receive the following JSON request:
```json
{
    "email": "minn@example.com"
}
```
On success: HttpStatus = 200
```json
{
    "success": true
}
```
On error: HttpStatus = 400
```json
{
    "error" : "User Exist Already !!!"
}
```