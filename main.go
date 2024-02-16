package main

import (
	"go-app1/router"
)

func main() {

	r := router.InitRouter()

	r.Run()
}
