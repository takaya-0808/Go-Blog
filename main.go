package main

import (
	"Go-Blog/app/src/infrastructure/router"
)

func main() {

	r := router.InitRouting()
	r.Run(":8022")
}
