package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetConnection() (*gorm.DB,error){
	dsn := "host=localhost user=root password=root dbname=golang port=8000 sslmode=disable TimeZone=Asia/Shanghai"
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