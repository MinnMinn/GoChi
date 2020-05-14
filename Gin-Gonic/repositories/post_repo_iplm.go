package repositories

import (
	"Gin-Gonic/models"
	"Gin-Gonic/driver"
	"database/sql"
)

type service struct {
	Db *sql.DB
}

func NewPostRepo(db *sql.DB) PostService {
	return &service {
		Db: db,
	}
}

func (s *service) AllPosts() ([]models.Post, error){
	//posts := make([]Model.Post, 0)
	//
	//rows, err := s.Db.Query("SELECT * FROM posts")
	//if err != nil {
	//	return posts, err
	//}
	//defer rows.Close()
	//for rows.Next() {
	//	post := Model.Post{}
	//	// SELECT id, name, gender, email FROM public.users;
	//	err := rows.Scan(&post.ID, &post.Title, &post.Content)
	//	if err != nil {
	//		break
	//	}
	//
	//	posts = append(posts, post)
	//}
	//
	//err = rows.Err()
	//if err != nil {
	//	return posts, err
	//}
	//
	//return posts, nil
	var post models.Post

	query, err := driver.DBConn().Query("SELECT * FROM posts")
	catch(err)

	var posts []models.Post
	for query.Next(){
		var id int
		var title, content string

		err = query.Scan(&id, &title, &content)
		if err != nil {
			panic(err.Error())
		}

		post.ID = id
		post.Title = title
		post.Content = content
		posts = append(posts, post)
	}

	return posts, err
}

func (s *service) DetailPost(id int64) (models.Post, error){
	var post models.Post

	query, err := driver.DBConn().Query("SELECT * FROM posts WHERE id=?", id)
	catch(err)

	for query.Next(){
		var id int
		var title, content string

		err = query.Scan(&id, &title, &content)
		if err != nil {
			panic(err.Error())
		}

		post.ID = id
		post.Title = title
		post.Content = content
	}

	return post, err
}

func (s *service) CreatePost(p models.Post) (error){
	query, err := driver.DBConn().Prepare("Insert posts SET title=?, content=?")
	catch(err)

	_, err = query.Exec(p.Title, p.Content)
	catch(err)
	defer query.Close()
	return err
}

func (s *service) UpdatePost(p models.Post) (error){
	query, err := driver.DBConn().Prepare("Update posts set title=?, content=? where id=?")
	catch(err)
	_, er := query.Exec(p.Title, p.Content, p.ID)
	catch(er)

	defer query.Close()
	return er
}

func (s *service) DeletePost(id int64) (error){
	query, err := driver.DBConn().Prepare("delete from posts where id=?")
	catch(err)
	_, er := query.Exec(id)
	catch(er)
	query.Close()
	return nil
}










//func AllPosts()([]Model.Post) {
//	var posts []Model.Post
//
//	query, err := driver.DBConn().Query("SELECT * FROM posts")
//	catch(err)
//
//	for query.Next(){
//		var id int
//		var title, content string
//
//		err = query.Scan(&id, &title, &content)
//		if err != nil {
//			panic(err.Error())
//		}
//		var post Model.Post
//
//		post.ID = id
//		post.Title = title
//		post.Content = content
//		posts = append(posts, post)
//	}
//
//	return posts
//}
//
//func DetailPost(id string)(Model.Post) {
//	var post Model.Post
//
//	query, err := driver.DBConn().Query("SELECT * FROM posts WHERE id=?", id)
//	catch(err)
//
//	for query.Next(){
//		var id int
//		var title, content string
//
//		err = query.Scan(&id, &title, &content)
//		if err != nil {
//			panic(err.Error())
//		}
//
//		post.ID = id
//		post.Title = title
//		post.Content = content
//	}
//
//	return post
//}
//
//func CreatePost(p Model.Post) {
//	var post Model.Post
//
//	query, err := driver.DBConn().Prepare("Insert posts SET title=?, content=?")
//	catch(err)
//
//	_, er := query.Exec(post.Title, post.Content)
//	catch(er)
//	defer query.Close()
//}
//
//func UpdatePost(id string) {
//	var post = DetailPost(id)
//
//	query, err := driver.DBConn().Prepare("Update posts set title=?, content=? where id=?")
//	catch(err)
//	_, er := query.Exec(post.Title, post.Content, id)
//	catch(er)
//
//	defer query.Close()
//}
//
//func DeletePost(id int) {
//
//	query, err := driver.DBConn().Prepare("delete from posts where id=?")
//	catch(err)
//	_, er := query.Exec(id)
//	catch(er)
//	query.Close()
//}










//func AllPosts(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	var post Model.Post
//	json.NewDecoder(r.Body).Decode(&post)
//
//	query, err := driver.DBConn().Query("SELECT * FROM posts")
//	catch(err)
//
//	var posts []Model.Post
//	for query.Next(){
//		var id int
//		var title, content string
//
//		err = query.Scan(&id, &title, &content)
//		if err != nil {
//			panic(err.Error())
//		}
//
//		post.ID = id
//		post.Title = title
//		post.Content = content
//		posts = append(posts, post)
//	}
//	defer query.Close()
//	json.NewEncoder(w).Encode(posts)
//
//	respondwithJSON(w, http.StatusCreated, map[string]string{"message": "successfully created"})
//}
//
//func DetailPost(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	var post Model.Post
//	json.NewDecoder(r.Body).Decode(&post)
//
//	query, err := driver.DBConn().Query("SELECT * FROM posts")
//	catch(err)
//
//	for query.Next(){
//		var id int
//		var title, content string
//
//		err = query.Scan(&id, &title, &content)
//		if err != nil {
//			panic(err.Error())
//		}
//
//		post.ID = id
//		post.Title = title
//		post.Content = content
//	}
//	defer query.Close()
//	json.NewEncoder(w).Encode(post)
//
//	respondwithJSON(w, http.StatusCreated, map[string]string{"message": "successfully created"})
//}
//
//func CreatePost(w http.ResponseWriter, r *http.Request) {
//	var post Model.Post
//	json.NewDecoder(r.Body).Decode(&post)
//
//	query, err := driver.DBConn().Prepare("Insert posts SET title=?, content=?")
//	catch(err)
//
//	_, er := query.Exec(post.Title, post.Content)
//	catch(er)
//	defer query.Close()
//
//	respondwithJSON(w, http.StatusCreated, map[string]string{"message": "successfully created"})
//}
//
//func UpdatePost(w http.ResponseWriter, r *http.Request) {
//	var post Model.Post
//	id := chi.URLParam(r, "id")
//	json.NewDecoder(r.Body).Decode(&post)
//
//	query, err := driver.DBConn().Prepare("Update posts set title=?, content=? where id=?")
//	catch(err)
//	_, er := query.Exec(post.Title, post.Content, id)
//	catch(er)
//
//	defer query.Close()
//
//	respondwithJSON(w, http.StatusOK, map[string]string{"message": "update successfully"})
//
//}
//
//func DeletePost(w http.ResponseWriter, r *http.Request) {
//	id := chi.URLParam(r, "id")
//
//	query, err := driver.DBConn().Prepare("delete from posts where id=?")
//	catch(err)
//	_, er := query.Exec(id)
//	catch(er)
//	query.Close()
//
//	respondwithJSON(w, http.StatusOK, map[string]string{"message": "successfully deleted"})
//}