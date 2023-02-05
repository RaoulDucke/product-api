package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "usr"
	password = "pwd"
	dbname   = "products"
)

func main() {
	r := gin.Default()

	r.GET("/products", func(c *gin.Context) { c.JSON(http.StatusOK, "hello") })

	r.Run()
}
