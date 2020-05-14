package repositories

import "Go-Chi/models"

type PostService interface {
	AllPosts() ([]models.Post, error)
	DetailPost(id int64) (models.Post, error)
	CreatePost(p models.Post) (error)
	UpdatePost(p models.Post) (error)
	DeletePost(id int64) (error)
}

//func AllPosts() ([]Model.Post) {
//	note := repository.AllPosts()
//	return note
//}
//
//func DetailPost(id string) (Model.Post) {
//	note := repository.DetailPost(id)
//	return note
//}
//
//func CreatePost(p Model.Post) {
//	repository.CreatePost(p)
//}
//
//func UpdatePost(id string){
//	repository.UpdatePost(id)
//}
//
//func DeletePost(id int){
//	repository.DeletePost(id)
//}