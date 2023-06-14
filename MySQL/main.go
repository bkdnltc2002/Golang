package main

import(
	"fmt"
	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Student struct{
	ID uint `gorm:"primaryKey" json:"student_id"` 
	Name string `gorm:"type:varchar(50)" json:"student_name"`
	ClassID uint `gorm:"foreignKey"`
	// Subjects []Subject `gorm:"many2many:student_subjects;`
}

type Class struct{
	ID uint `gorm:"primaryKey" json:"class_id"`
	Name string `gorm:"type:varchar(50)" json:"class_name"`
}

type Subject struct{
	ID uint `gorm:"primaryKey" json:"subject_id"`
	Name string `gorm:"type:varchar(50)" json:"subject_name"`
}

func main(){
	dsn :="root:root123@tcp(localhost:30001)/Students"
	db,err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err!= nil{
		log.Fatal(err)
	}

	// Automigrate the table
	err = db.AutoMigrate(&Student{}, &Class{}, &Subject{})
	if err!= nil{
		log.Fatal(err)
	}

	// Find a student by ID
	var student Student
	err = db.First(&student, 1).Error
	if err!= nil{
		log.Fatal(err)
	}
	fmt.Println(student)

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