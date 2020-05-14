package router

import (
	"SP_FriendManagement_API_TungNguyen/driver"
	"SP_FriendManagement_API_TungNguyen/repositories"
	"SP_FriendManagement_API_TungNguyen/services"
	"database/sql"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"net/http"
	"time"
)

var router *chi.Mux
var dbcon *sql.DB

func Router() {
	DBcon := driver.DBConn()
	log.Println("LET'S GO !!! ")
	var API = repositories.Database{DBcon}
	var service = services.FriendsService{API}
	router = chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Get("/", service.GetAllEmail)
	router.Post("/addUser", service.AddUser)
	router.Post("/addFriend", service.AddFriend)
	router.Get("/findFriendsOfUser", service.FindFriendsOfUser)
	router.Get("/findCommonFriends", service.FindCommonFriends)
	router.Post("/subscribeFriend", service.SubscribeFriend)
	router.Post("/blockFriend", service.BlockFriend)
	router.Get("/receiveUpdateFromEmail", service.ReceiveUpdatesFromEmail)

	http.ListenAndServe(":8005", logger())
}

func logger() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(time.Now(), r.Method, r.URL)
		router.ServeHTTP(w, r)
	})
}