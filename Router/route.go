package router

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"io"
)

type Address struct{
	Street string 	`json:"street"`
	Province string	`json:"province"`
	Country string	`json:"country"`
}

type Student struct{
	Id int			`json:"id"`
	Name string		`json:"name"`
	Age int			`json:"age"`
	GPA float64		`json:"gpa"`
	Address  Address `json:"address"`
}

func studentGet(w http.ResponseWriter,r *http.Request){
	dummyData := Student{
		Id: 1,
		Name: "Stewy",
		Age: 20,
		GPA: 3.2,
		Address: Address{
			Street: "21th Orange Street",
			Province: "Unknown",
			Country: "Japan",
		},
	}

	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(dummyData)
}

func studentPost(w http.ResponseWriter,r *http.Request){
	body,_ := io.ReadAll(r.Body)
	defer r.Body.Close()

	var student Student
	if err := json.Unmarshal(body,&student);err!=nil{
		http.Error(w,"Invalid Json",http.StatusBadRequest)
		return
	}

	log.Println(student)

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"Message":"Successfully Created New Student"})
}

func studentPut(w http.ResponseWriter,r *http.Request){

}

func studentDelete(w http.ResponseWriter,r *http.Request){

}

func StudentRouter(w http.ResponseWriter,r *http.Request){
	log.Printf("%v %v %v %v %v %v",r.Method,r.Host,r.URL,r.Body,r.UserAgent(),time.Now())
	switch (r.Method){
	case http.MethodGet:
		studentGet(w,r)
	case http.MethodPost:
		studentPost(w,r)
	case http.MethodPut:
		studentPut(w,r)
	case http.MethodDelete:
		studentDelete(w,r)
	}
}