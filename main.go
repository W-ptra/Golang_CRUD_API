package main

import (
	"fmt"
	"net/http"
	"os"
	"github.com/joho/godotenv"
	"github.com/W-ptra/Golang_CRUD_API/Router"
)

func main() {
	godotenv.Load()

	mux := http.NewServeMux()
	mux.HandleFunc("/api/student",router.StudentRouter)

	http.ListenAndServe(fmt.Sprintf("%v:%v",os.Getenv("HOST"),os.Getenv("PORT")),mux)
}