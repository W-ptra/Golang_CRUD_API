package main

import (
	"fmt"
	"os"
	"github.com/W-ptra/Golang_CRUD_API/Database"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	database.Migration()
	server := newAPIServer(fmt.Sprintf("%v:%v",os.Getenv("HOST"),os.Getenv("PORT")))
	server.run()
}