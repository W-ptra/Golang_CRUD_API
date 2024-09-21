package database

import (
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbConnection *gorm.DB

func GetConnection() (*gorm.DB,error){
	if dbConnection == nil{

		godotenv.Load()
		dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=Asia/Shanghai", 
			os.Getenv("DB_HOST"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_SSLMODE"),
		)
		connection,err := gorm.Open(postgres.Open(dsn),&gorm.Config{})
		if err!=nil{
			return nil,err
		}
		dbConnection = connection
	}
	return dbConnection,nil
}

func Migration(){
	db,err := GetConnection()
	if err != nil{
		fmt.Println("Error to connect to database:\n",err)
	}

	err = db.AutoMigrate(&Student{})
	if err != nil{
		fmt.Println("Error while migrating database:\n",err)
	}
}

func CreateStudent(student Student) error{
	db,err := GetConnection()
	if err!=nil{
		return err
	}
	result := db.Create(&student)
	return result.Error
}

func GetStudent() ([]Student,error){
	db,err := GetConnection()
	if err!=nil{
		return nil,err
	}
	var studentList []Student
	result := db.Find(&studentList)
	return studentList,result.Error
}

func GetStudentById(id int) (Student,error){
	var student Student
	db,err := GetConnection()
	if err!=nil{
		return Student{},err
	}
	result := db.First(&student,"Id = ?",id)
	return student,result.Error
}

func UpdateStudentById(student Student) error{
	db,err := GetConnection()
	if err!=nil{
		return err
	}
	studentDB,err := GetStudentById(student.Id)
	if err != nil{
		log.Println(err)
		return err
	}

	studentDB.Name = student.Name
	studentDB.Age = student.Age
	studentDB.GPA = student.GPA
	studentDB.Street = student.Street
	studentDB.Province = student.Province
	studentDB.Country = student.Country
	
	if err := db.Save(studentDB).Error; err != nil {
		return err
	}

	return nil
}

func DeleteStudentById(id int) error{
	db,err := GetConnection()
	if err!=nil{
		return err
	}
	student,err := GetStudentById(id)
	if err!=nil{
		return err
	}

	return db.Delete(&Student{},student.Id).Error
}