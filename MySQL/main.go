package main

import (
	// "database/sql"
	"fmt"
	"log"
	"sync"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Student struct{
	ID uint 
	StudentName string
	ClassID uint
	ClassInfo Class `gorm:"foreignKey:ClassID;->"`
	ExamInfo []Exam `gorm:"foreignKey:StudentID;->"`
}

type Class struct{
	ID uint
	ClassName string 
}

type Subject struct{
	ID uint 
	SubjectName string 
}

type StudentSubject struct{
	StudentID uint 
	SubjectID uint 
}

type Exam struct{
	ID uint 
	StudentID uint
	Mark uint 
}
var once sync.Once
var db *gorm.DB


func GetConn() *gorm.DB {
	once.Do( func(){
		dsn :="root:root123@tcp(localhost:30001)/Students"
		db,err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err!= nil{
			log.Fatal(err)
		}
	}) 
	return db
}

func main(){
	
	// Verify if connection is ok
	conn := GetConn()
	err := conn.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected ✓")
	err = conn.Close()
	if err != nil {
		log.Fatal(err)
	}
	// Automigrate the table
	// err = db.AutoMigrate(&Student{}, &Class{}, &Subject{})
	// if err!= nil{
	// 	log.Fatal(err)
	// }

	// Find a student by ID
	var students []Student
	err = db.Preload("ClassInfo").Preload("ExamInfo").Find(&students).Error
	if err!= nil{
		log.Fatal(err)
	}
	for _,s:= range students{
		fmt.Printf("%+v\n",s)
	}
	
	

	// Find all students
	// var students []Student
	// err = db.Find(&students).Error
	// if err!= nil{
	// 	log.Fatal(err)
	// }
	// fmt.Println(students)

	// Find all subjects for a student
	// err = db.Preload("subject").First(&student,1).Error
	// if err!= nil{
	// 	log.Fatal(err)
	// }
	// fmt.Println(student.Subjects)

	// Find all students in a class
	// var studentInClass []Student
	// err = db.Preload("subject").Where("class-id = ?",1).Find(&studentInClass).Error
	// if err!= nil{
	// 	log.Fatal(err)
	// }
	// fmt.Println(studentInClass)
}