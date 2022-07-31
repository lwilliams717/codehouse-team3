package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type User struct {
	Id      			int    	`json:"id"`
	UserType			string	`json:"user_type"` // either Tutor or Student
	Name   				string 	`json:"name"`
	PrimarySubject 		string 	`json:"primary_subject"`
	SecondarySubject 	string 	`json:"secondary_subject"`
	Description 		string 	`json:"description"`
}

var nextId int = 0
var Tutors []User

var Names = [...]string{
	"Mohammed",
	"Nushi",
	"Wei",
	"David",
	"Antonio",
	"Anita",
	"Svetlana",
	"Mo",
	"Xiaoyan",
	"Ana",
}

var PrimarySubjects = [...]string{
	"Science",
	"Technology",
	"Engineering",
	"Math",
}

var Science = [...]string{
	"Biology",
	"Physics",
	"Chemistry",
}
var Technology = [...]string{
	"Computer Science",
	"Digital Modeling",
	"Machine Learning",
	"Game Development",
}
var Engineering = [...]string{
	"Robotics",
	"Electronics",
	"Mechanical",
	"Civil",
}
var Math = [...]string{
	"Algebra",
	"Geometry",
	"Statistics",
	"Calculus",
}

func GetNextId() int {
	nextId++
	return nextId - 1
}

func createRandomTutorList(c *gin.Context) {
	Tutors = append(Tutors, User{Id: GetNextId(), UserType: "Tutor", Name: "Mohammed", PrimarySubject: "Technology", SecondarySubject: "Game Development"})
	Tutors = append(Tutors, User{Id: GetNextId(), UserType: "Tutor", Name: "Nushi", PrimarySubject: "Math", SecondarySubject: "Algebra"})
	Tutors = append(Tutors, User{Id: GetNextId(), UserType: "Tutor", Name: "Wei", PrimarySubject: "Math", SecondarySubject: "Geometry"})
	Tutors = append(Tutors, User{Id: GetNextId(), UserType: "Tutor", Name: "David", PrimarySubject: "Technology", SecondarySubject: "Computer Science"})
	Tutors = append(Tutors, User{Id: GetNextId(), UserType: "Tutor", Name: "Antonio", PrimarySubject: "Technology", SecondarySubject: "Machine Learning"})
	Tutors = append(Tutors, User{Id: GetNextId(), UserType: "Tutor", Name: "Anita", PrimarySubject: "Science", SecondarySubject: "Physics"})
	Tutors = append(Tutors, User{Id: GetNextId(), UserType: "Tutor", Name: "Svetlana", PrimarySubject: "Engineering", SecondarySubject: "Electronics"})
	Tutors = append(Tutors, User{Id: GetNextId(), UserType: "Tutor", Name: "Mo", PrimarySubject: "Math", SecondarySubject: "Calculus"})
	Tutors = append(Tutors, User{Id: GetNextId(), UserType: "Tutor", Name: "Xiaoyan", PrimarySubject: "Engineering", SecondarySubject: "Civil"})
	Tutors = append(Tutors, User{Id: GetNextId(), UserType: "Tutor", Name: "Ana", PrimarySubject: "Technology", SecondarySubject: "Digital Modeling"})
	c.JSON(http.StatusOK, gin.H{"Tutor List": Tutors})
}

func main() {
	router := gin.Default()
	router.GET("/", createRandomTutorList)
	router.Run(":8090")
}
