package database

import (
	"gorm.io/driver/postgres"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"fmt"
	"os"
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
	return db.Model(&Student{}).Where("Id = ?",student.Id).Updates(Student{
		Name: student.Name,
		Age: student.Age,
		GPA: student.GPA,
		Street: student.Street,
		Province: student.Province,
		Country: student.Country,
	}).Error
}

func DeleteStudentById(id int) error{
	db,err := GetConnection()
	if err!=nil{
		return err
	}
	return db.Delete(&Student{},id).Error
}