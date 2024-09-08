package database

import (
	"gorm.io/driver/postgres"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"fmt"
	"os"
)

func GetConnection() (*gorm.DB,error){
	godotenv.Load()

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=Asia/Shanghai",os.Getenv("DB_HOST"),os.Getenv("DB_USER"),os.Getenv("DB_PASSWORD"),os.Getenv("DB_NAME"),os.Getenv("DB_PORT"),os.Getenv("DB_SSLMODE"))
	db,err := gorm.Open(postgres.Open(dsn),&gorm.Config{})

	if err!=nil{
		return nil,err
	}

	return db,nil
}

func CreateStudent(db *gorm.DB,student Student) error{
	result := db.Create(&student)

	return result.Error
}

func GetStudent(db *gorm.DB) ([]Student,error){
	var studentList []Student
	result := db.Find(&studentList)
	return studentList,result.Error
}

func GetStudentById(db *gorm.DB,id int) (Student,error){
	var student Student
	result := db.First(&student,"Id = ?",id)
	return student,result.Error
}

func UpdateStudentById(db *gorm.DB,student Student) error{
	return db.Model(&Student{}).Where("Id = ?",student.Id).Updates(Student{
		Name: student.Name,
		Age: student.Age,
		GPA: student.GPA,
		Address: student.Address,
	}).Error
}

func DeleteStudentById(db *gorm.DB,id int) error{
	return db.Delete(&Student{},id).Error
}