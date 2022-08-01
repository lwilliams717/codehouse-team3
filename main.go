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

type Topic struct {
	PrimarySubject      		string    	`json:"primary_subject"`
	SecondarySubjects			[]string	`json:"secondary_subjects"`
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

//creates static list of tutors
func createTutorList() {
	Tutors = append(Tutors, User{Id: GetNextId(), UserType: "Tutor", Name: "Mohammed", PrimarySubject: PrimarySubjects[0], SecondarySubject: Science[0]})
	Tutors = append(Tutors, User{Id: GetNextId(), UserType: "Tutor", Name: "Nushi", PrimarySubject: PrimarySubjects[1], SecondarySubject: Technology[0]})
	Tutors = append(Tutors, User{Id: GetNextId(), UserType: "Tutor", Name: "Wei", PrimarySubject: PrimarySubjects[2], SecondarySubject: Engineering[0]})
	Tutors = append(Tutors, User{Id: GetNextId(), UserType: "Tutor", Name: "David", PrimarySubject: PrimarySubjects[3], SecondarySubject: Math[0]})
	Tutors = append(Tutors, User{Id: GetNextId(), UserType: "Tutor", Name: "Antonio", PrimarySubject: PrimarySubjects[0], SecondarySubject: Science[1]})
	Tutors = append(Tutors, User{Id: GetNextId(), UserType: "Tutor", Name: "Anita", PrimarySubject: PrimarySubjects[1], SecondarySubject: Technology[1]})
	Tutors = append(Tutors, User{Id: GetNextId(), UserType: "Tutor", Name: "Svetlana", PrimarySubject: PrimarySubjects[2], SecondarySubject: Engineering[1]})
	Tutors = append(Tutors, User{Id: GetNextId(), UserType: "Tutor", Name: "Mo", PrimarySubject: PrimarySubjects[3], SecondarySubject: Math[1]})
	Tutors = append(Tutors, User{Id: GetNextId(), UserType: "Tutor", Name: "Xiaoyan", PrimarySubject: PrimarySubjects[0], SecondarySubject: Science[2]})
	Tutors = append(Tutors, User{Id: GetNextId(), UserType: "Tutor", Name: "Ana", PrimarySubject: PrimarySubjects[1], SecondarySubject: Technology[2]})
}

func PostTopics(c *gin.Context) {
	var topics []Topic
	topics = append(topics, Topic{PrimarySubject: PrimarySubjects[0], SecondarySubjects: Science[0:3]})
	topics = append(topics, Topic{PrimarySubject: PrimarySubjects[1], SecondarySubjects: Technology[0:4]})
	topics = append(topics, Topic{PrimarySubject: PrimarySubjects[2], SecondarySubjects: Engineering[0:4]})
	topics = append(topics, Topic{PrimarySubject: PrimarySubjects[3], SecondarySubjects: Math[0:4]})
	c.JSON(http.StatusOK, gin.H{"topics": topics})
}

//binds subject JSON and posts matched tutors JSON to the webserver
func PostTutors(c *gin.Context) {
	var item Search
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	secondary := item.SecondarySubject
	var newTutors []User = match(secondary)
	c.JSON(http.StatusOK, gin.H{"matched_tutors": newTutors})
}

//function to match subject to tutors and build list
func match( subject string ) []User {
	newTutors := []User{}
	for _, tutor := range Tutors{
		if subject == tutor.SecondarySubject{
			newTutors = append(newTutors, tutor)
		}
	}
	return newTutors
}

func main() {
	router := gin.Default()
	router.POST("/api/topics", PostTopics)
	createTutorList()
	router.POST("/api/tutors", PostTutors)
	router.Run(":8090")
}