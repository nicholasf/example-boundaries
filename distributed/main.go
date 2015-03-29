package main

import "github.com/gin-gonic/gin"

func main() {

	go runHouses(":8000")
	go runUsers(":8001")
}
