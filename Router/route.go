package router1

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"github.com/W-ptra/Golang_CRUD_API/Database"
)

type Message struct{
	Message string `json:"message"`
}

type Student struct{
	Id 			int		`json:"id"`
	Name 		string	`json:"name"`
	Age 		int		`json:"age"`
	GPA 		float64	`json:"gpa"`
	Street 		string 	`json:"street"`
	Province 	string	`json:"province"`
	Country 	string	`json:"country"`
}

func setRespond(w http.ResponseWriter,payload interface{},statusCode int){
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(payload)
}

func StudentGet(w http.ResponseWriter,r *http.Request){
	data,err := database.GetStudent()
	if err != nil{
		setRespond(w,Message{"Something went wrong"},500)
		return
	}
	setRespond(w,data,200)
}

func StudentGetById(w http.ResponseWriter,r *http.Request){
	studentId,err := strconv.Atoi(r.PathValue("id")) // get parameter value 
	if err !=nil{
		setRespond(w,Message{"Something went wrong"},500)
		return
	}

	data,err := database.GetStudentById(studentId)
	if err != nil && err.Error() == "record not found"{
		setRespond(w,Message{"Student not found"},404)
		return
	}
	setRespond(w,data,200)
}

func StudentPost(w http.ResponseWriter,r *http.Request){
	body,_ := io.ReadAll(r.Body)
	defer r.Body.Close()

	var student Student
	if err := json.Unmarshal(body,&student);err!=nil{
		setRespond(w,Message{"Invalid Json"},500)
		return
	}
	log.Println(student)
	newStudent := database.Student{
		Name: 		student.Name,
		Age: 		student.Age,
		GPA: 		student.GPA,
		Street: 	student.Street,
		Province: 	student.Province,
		Country: 	student.Country,
	}
	err := database.CreateStudent(newStudent)

	if err != nil{
		setRespond(w,Message{"Something went wrong, can't create new student"},500)
		return
	}
	setRespond(w,Message{"successfully created new student"},200)
}

func StudentPut(w http.ResponseWriter,r *http.Request){
	studentId,err := strconv.Atoi(r.PathValue("id"))
	if err!=nil{
		setRespond(w,Message{"Invalid path variable, expected to be Integer"},400)
		return
	}

	body,_ := io.ReadAll(r.Body)
	defer r.Body.Close()

	var student Student
	if err:=json.Unmarshal(body,&student);err!=nil{
		setRespond(w,Message{"Invalid Json"},400)
		return
	}
	student.Id = studentId
	log.Println(student)

	newStudent := database.Student{
		Id: studentId,
		Name: student.Name,
		Age: student.Age,
		GPA: student.GPA,
		Street: student.Street,
		Province: student.Province,
		Country: student.Country,
	}

	err = database.UpdateStudentById(newStudent)

	if err != nil && err.Error() == "record not found"{
		setRespond(w,Message{"student not found"},404)
		return
	}
	setRespond(w,Message{"successfully update student"},200)
}

func StudentDelete(w http.ResponseWriter,r *http.Request){
	studentId,err := strconv.Atoi(r.PathValue("id"))
	if err!=nil{
		http.Error(w,"Invalid path variable, expected to be Integer",http.StatusBadRequest)
		return
	}
	err = database.DeleteStudentById(studentId)

	if err != nil && err.Error() == "record not found"{
		setRespond(w,Message{"student not found"},404)
		return
	}
	setRespond(w,Message{"successfully delete student"},200)
}