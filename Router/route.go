package router1

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
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
	var dummyData []Student
	dummyData = append(dummyData,Student{
		Id: 1,
		Name: "Stewy",
		Age: 20,
		GPA: 3.2,
		Address: Address{
			Street: "21th Orange Street",
			Province: "Unknown",
			Country: "Japan",
		},
	})

	dummyData = append(dummyData,Student{
		Id: 2,
		Name: "Robert",
		Age: 22,
		GPA: 2.9,
		Address: Address{
			Street: "21th Mallard Street",
			Province: "Washington",
			Country: "USA",
		},
	})
	
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(dummyData)
}

func StudentGetById(w http.ResponseWriter,r *http.Request){

	studentId,err := strconv.Atoi(r.PathValue("id"))
	if err!=nil{
		http.Error(w,"Invalid path variable, expected to be Integer",http.StatusBadRequest)
		return
	}

	dummyData := map[string]interface{}{
		"studentId":studentId,
		"student": Student{
					Id: 1,
					Name: "Stewy",
					Age: 20,
					GPA: 3.2,
					Address: Address{
						Street: "21th Orange Street",
						Province: "Unknown",
						Country: "Japan",
					},
				},
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dummyData)
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

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message":"successfully delete student","studentId":strconv.Itoa(studentId)})
}