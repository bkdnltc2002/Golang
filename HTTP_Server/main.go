package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	connection *gorm.DB
}

var (
	db      *Database
	dbMutex sync.Mutex
	once    sync.Once
)

func GetDatabase() *Database {
	once.Do(func() {
		dbMutex.Lock()
		defer dbMutex.Unlock()
		db = &Database{}
		db.initConnection()
	})
	return db
}

func (db *Database) initConnection() {
	dsn := "root:root123@tcp(localhost:30001)/Students"
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.connection = conn
}

func (db *Database) CloseConnection() {
	sqlDB, err := db.connection.DB()
	if err != nil {
		log.Fatalf("Failed to get SQL DB: %v", err)
	}
	sqlDB.Close()
}

type Student struct {
	ID          uint
	StudentName string
	ClassID     uint
	ClassInfo   Class `gorm:"foreignKey:ClassID;->"`
	// ExamInfo []Exam `gorm:"foreignKey:StudentID;->"`
}

type Class struct {
	ID        uint
	ClassName string
}

type Subject struct {
	ID          uint
	SubjectName string
}

type StudentSubject struct {
	StudentID uint
	SubjectID uint
}

// type Exam struct{
// 	ID uint
// 	StudentID uint
// 	Mark uint
// }

func getUserHandler(c *gin.Context) {
	db := GetDatabase()
	var student []Student
	result := db.connection.Find(&student)
	if result.Error != nil {
		c.String(http.StatusInternalServerError, "Failed to fetch data from the database")
		return
	}

	// for _,s:= range student{
	// 	fmt.Printf("%+v\n",s)
	// }
	var studentName []Student
	err := db.connection.Preload("ClassInfo").Find(&studentName).Error
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(studentName)

	c.JSON(http.StatusOK, student)
}

func search(c *gin.Context) {
	db := GetDatabase()
	searchQuery := c.Query("name")

	var students []Student
	err := db.connection.Where("student_name LIKE ?", "%"+searchQuery+"%").Find(&students).Error
	if err != nil {
		c.String(http.StatusInternalServerError, "The name does not exist")
		return
	}
	c.JSON(http.StatusOK, students)
}

func main() {
	r := gin.Default()

	r.GET("/v1/students", getUserHandler)
	r.GET("/v1/students/search", search)
	r.Run(":8080")
	defer GetDatabase().CloseConnection()
}
