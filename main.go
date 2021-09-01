package main

import (
	"example.com/estudo/database"
	"example.com/estudo/server"
)

func main() {

	database.LigarDB()
	server := server.NewServer()
	server.Run()
}
