package main

import (
	"Gin-Gonic/router"
)

func main() {
	router.Router().Run(":8005")
}