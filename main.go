package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

//user format struct
type User struct {
	Id      			int    	`json:"id"`
	UserType			string	`json:"user_type"` // either Tutor or Student
	Name   				string 	`json:"name"`
	PrimarySubject 		string 	`json:"primary_subject"`
	SecondarySubject 	string 	`json:"secondary_subject"`
	Description 		string 	`json:"description"`
}

//search format struct for the searched subject
type Search struct{
	PrimarySubject   				string 	`json:"primary_subject"`
	SecondarySubject   				string 	`json:"secondary_subject"`
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

//creates static list of tutors
func createTutorList() {
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

//binds subject JSON and posts matched tutors JSON to the webserver
func PostTutors(c *gin.Context) {
	var item Search
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	primary := item.PrimarySubject
	var newTutors []User = match(primary)
	c.JSON(http.StatusOK, gin.H{"matched_tutors": newTutors})
}

//function to match subject to tutors and build list
func match( subject string ) []User {
	var newTutors []User
	for _, tutor := range Tutors{
		if subject == tutor.SecondarySubject{
			newTutors = append(newTutors, tutor)
		}
	}
	return newTutors
}

func main() {
	router := gin.Default()
	//router.GET("/api/topics", )
	createTutorList()
	router.POST("/api/tutors", PostTutors)
	router.Run(":8090")
}