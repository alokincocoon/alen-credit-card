package main

import (
	// "github.com/gin-gonic/gin"
	"go-app1/router"
	// "github.com/edgedb/edgedb-go"
)

// var db *edgedb.Client

func main() {

	r := router.InitRouter()

	r.Run()
}
