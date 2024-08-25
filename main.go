package main

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	server := newAPIServer(fmt.Sprintf("%v:%v",os.Getenv("HOST"),os.Getenv("PORT")))
	server.run()
}