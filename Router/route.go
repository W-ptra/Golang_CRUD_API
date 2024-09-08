package router1

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"github.com/W-ptra/Golang_CRUD_API/Database"
)

type Address struct{
	Street 		string 	`json:"street"`
	Province 	string	`json:"province"`
	Country 	string	`json:"country"`
}

type Student struct{
	Id 			int		`json:"id"`
	Name 		string	`json:"name"`
	Age 		int		`json:"age"`
	GPA 		float64	`json:"gpa"`
	Address  	Address	`json:"address"`
}

func StudentGet(w http.ResponseWriter,r *http.Request){
	db,err := database.GetConnection()
	if err != nil{
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message":"Something went wrong, can't get all student"})
		return
	}

	data,err := database.GetStudent(db)
	
	if err != nil{
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message":"Something went wrong, can't get all student"})
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(data)
}

func StudentGetById(w http.ResponseWriter,r *http.Request){

	studentId,err := strconv.Atoi(r.PathValue("id"))
	if err!=nil{
		http.Error(w,"Invalid path variable, expected to be Integer",http.StatusBadRequest)
		return
	}

	db,err := database.GetConnection()

	if err != nil{
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message":"Something went wrong, can't get student by id"})
		return
	}

	data,err := database.GetStudentById(db,studentId)

	if err != nil{
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message":"Something went wrong, can't get student by id"})
		return
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func StudentPost(w http.ResponseWriter,r *http.Request){
	body,_ := io.ReadAll(r.Body)
	defer r.Body.Close()

	var student Student
	if err := json.Unmarshal(body,&student);err!=nil{
		http.Error(w,"Invalid Json",http.StatusBadRequest)
		return
	}

	log.Println(student)
	newAddress := database.Address{
		Street: student.Address.Street,
		Province: student.Address.Province,
		Country: student.Address.Country,
	}
	newStudent := database.Student{
		Name: student.Name,
		Age: student.Age,
		GPA: student.GPA,
		Address: newAddress,
	}
	db,err := database.GetConnection()
	
	if err != nil{
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message":"Something went wrong, can't create new student"})
		return
	}

	err = database.CreateStudent(db,newStudent)

	if err != nil{
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message":"Something went wrong, can't create new student"})
		return
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message":"successfully created new student"})
}

func StudentPut(w http.ResponseWriter,r *http.Request){
	studentId,err := strconv.Atoi(r.PathValue("id"))
	if err!=nil{
		http.Error(w,"Invalid path variable, expected to be Integer",http.StatusBadRequest)
		return
	}

	body,_ := io.ReadAll(r.Body)
	defer r.Body.Close()

	var student Student
	if err:=json.Unmarshal(body,&student);err!=nil{
		http.Error(w,"Invalid Json",http.StatusBadRequest)
		return
	}
	student.Id = studentId
	log.Println(student)

	db,err := database.GetConnection()
	if err != nil{
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message":"Something went wrong, can't update student"})
		return
	}

	newAddress := database.Address{
		Street: student.Address.Street,
		Province: student.Address.Province,
		Country: student.Address.Country,
	}
	newStudent := database.Student{
		Name: student.Name,
		Age: student.Age,
		GPA: student.GPA,
		Address: newAddress,
	}

	err = database.UpdateStudentById(db,newStudent)

	if err != nil{
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message":"Something went wrong, can't update student"})
		return
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message":"successfully update student","studentId":strconv.Itoa(studentId)})
}

func StudentDelete(w http.ResponseWriter,r *http.Request){
	studentId,err := strconv.Atoi(r.PathValue("id"))
	if err!=nil{
		http.Error(w,"Invalid path variable, expected to be Integer",http.StatusBadRequest)
		return
	}

	db,err := database.GetConnection()

	if err != nil{
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message":"Something went wrong, can't delete student"})
		return
	}

	err = database.DeleteStudentById(db,studentId)

	if err != nil{
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message":"Something went wrong, can't delete student"})
		return
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message":"successfully delete student","studentId":strconv.Itoa(studentId)})
}