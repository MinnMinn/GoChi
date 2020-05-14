package driver

import (
	"database/sql"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/go-sql-driver/mysql"
)

var router *chi.Mux

const (
	dbName = "Golang"
	dbUser = "root"
	dbPass = "secret"
	dbDriver = "mysql"
)

func DBConn() (db *sql.DB) {
	router = chi.NewRouter()
	router.Use(middleware.Recoverer)

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}