package router

import (
	"Go-Chi/controllers"
	"Go-Chi/services"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"time"
)

var router *chi.Mux

func Router() {
	router = chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Get("/posts", services.AllPosts)
	router.Get("/posts/{id}", services.DetailPost)
	router.Post("/posts", services.CreatePost)
	router.Put("/posts/{id}", services.UpdatePost)
	router.Delete("/posts/{id}", services.DeletePost)
	router.Post("/addFriend", controllers.AddFriend)
	router.Get("/findFriendOfUser", controllers.FindFriendsOfUser)
	router.Get("/findCommonFriends", controllers.FindCommonFriends)
	router.Post("/followFriend", controllers.FollowFriend)
	router.Post("/blockFriend", controllers.BlockFriend)
	router.Get("/receiveUpdatesFromEmail", controllers.ReceiveUpdatesFromEmail)
	http.ListenAndServe(":8005", logger())
}

func logger() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(time.Now(), r.Method, r.URL)
		router.ServeHTTP(w, r)
	})
}