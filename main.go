package main

import (
	"fmt"
	"os"
	"github.com/W-ptra/Golang_CRUD_API/Database"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	db,err := database.GetConnection()
	if err != nil{
		fmt.Println("Error to connect to database:\n",err)
	}

	err = db.AutoMigrate(&database.Student{})
	if err != nil{
		fmt.Println("Error while migrating database:\n",err)
	}

	server := newAPIServer(fmt.Sprintf("%v:%v",os.Getenv("HOST"),os.Getenv("PORT")))
	server.run()
}