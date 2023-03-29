package models

import (
	"github.com/jinzhu/gorm"
	"github.com/Yosephine25/go_students_crud_mysql/pkg/config"
)


var db *gorm.DB

type Student struct {
	gorm.Model
	NIM string `gorm:""json:"nim"`
	Name string `json:"name"`
	IPK string `json:"ipk"`
	Jurusan string `json:"jurusan"`
	Angkatan string `json:"angkatan"`
}
func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Student{})
}
func (b *Student) CreateStudent() *Student {
	db.NewRecord(b)
	db.Create(&b)
	return b
}
func GetAllStudents() []Student {
	var Students []Student
	db.Find(&Students)
	return Students
}
func GetStudentbyId(nim int64) (*Student, *gorm.DB) {
	var getStudent Student
	db := db.Where("NIM=11321065", nim).Find(&getStudent)
	return &getStudent, db
}

func DeleteStudent(nim int64) Student {
	var student Student
	db.Where("NIM=11321065", nim).Delete(student)
	return student
}
