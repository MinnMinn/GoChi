package services

import (
	"Go-Chi/models"
	"Go-Chi/driver"
	"Go-Chi/repositories"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

func AllPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := driver.DBConn()
	postRepo := repositories.NewPostRepo(db)
	var posts, _ = postRepo.AllPosts()
	json.NewEncoder(w).Encode(posts)
	respondwithJSON(w, http.StatusCreated, map[string]string{"message": "successfully created"})
}

func DetailPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := driver.DBConn()
	postRepo := repositories.NewPostRepo(db)
	id := chi.URLParam(r, "id")
	number, _ := strconv.ParseInt(id, 10, 0)
	var post, _ = postRepo.DetailPost(number)
	json.NewEncoder(w).Encode(post)
	respondwithJSON(w, http.StatusCreated, map[string]string{"message": "successfully created"})
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := driver.DBConn()
	postRepo := repositories.NewPostRepo(db)
	var post models.Post
	json.NewDecoder(r.Body).Decode(&post)
	postRepo.CreatePost(post)

	respondwithJSON(w, http.StatusCreated, map[string]string{"message": "successfully created"})
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := driver.DBConn()
	postRepo := repositories.NewPostRepo(db)
	var post models.Post
	json.NewDecoder(r.Body).Decode(&post)
	id := chi.URLParam(r, "id")
	number, _ := strconv.ParseInt(id, 10, 0)
	post.ID = int(number)
	postRepo.UpdatePost(post)

	respondwithJSON(w, http.StatusOK, map[string]string{"message": "update successfully"})

}

// DeletePost remove a spesific post
func DeletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := driver.DBConn()
	postRepo := repositories.NewPostRepo(db)
	id := chi.URLParam(r, "id")
	number, _ := strconv.ParseInt(id, 10, 0)
	postRepo.DeletePost(number)

	respondwithJSON(w, http.StatusOK, map[string]string{"message": "successfully deleted"})
}

func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	fmt.Println(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}