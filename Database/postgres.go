package database

import (
	router1 "github.com/W-ptra/Golang_CRUD_API/Router"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getConnection() (*gorm.DB,error){
	dsn := "host=localhost user=your_user password=your_password dbname=your_dbname port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db,err := gorm.Open(postgres.Open(dsn),&gorm.Config{})

	if err!=nil{
		return nil,err
	}

	return db,nil
}

func createStudent(db *gorm.DB,student router1.Student) error{
	newStudent := Student{
		Name: student.Name,
		Age: student.Age,
		GPA: student.GPA,
		Address: Address(student.Address),
	}

	result := db.Create(&newStudent)

	return result.Error
}

func getStudentById(db *gorm.DB,id int) (Student,error){
	var student Student
	result := db.First(&student,"Id = ?",id)
	return student,result.Error
}

func updateStudentById(db *gorm.DB,student Student) error{
	return db.Model(&Student{}).Where("Id = ?",student.Id).Updates(Student{
		Name: student.Name,
		Age: student.Age,
		GPA: student.GPA,
		Address: student.Address,
	}).Error
}

func deleteStudentById(db *gorm.DB,id int) error{
	return db.Delete(&Student{},id).Error
}