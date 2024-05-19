package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// model
type Course struct {
	CourseId     string `json:"courseid"`
	CourseName   string `json:"coursename"`
	CoursePrice  int    `json:"price"`
	CourseAuthor string `json:"author"`
}
type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

// fake db
var courses []Course

// middleware
func (c *Course) isEmpty() bool {
	return (c.CourseName == "" )
}

func main() {
	fmt.Println("creating api")
}

// controller
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>raboo's Home</h1>"))
}
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getting All Courses called")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}
func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getting your id Courses called")
	// id
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("no course found with thus id")

}
func addCourse(w http.ResponseWriter, r *http.Request)  {
	// request has course details in json format
	fmt.Println("adding controller")
	w.Header().Set("Content-Type", "application/json")
	var course Course
	// decoding the req json body to ds Course
	
	if r.Body==nil{
		// send in json format error msg
		json.NewEncoder(w).Encode("empty req body")
		return
	}
	json.NewDecoder(r.Body).Decode(&course)
	 if course.isEmpty(){
		json.NewEncoder(w).Encode("courseName not provided")
	}else{
		// perfectly fine 
		// appending in db 
		courses=append(courses, course)
		json.NewDecoder(w).Decode(&courses)
	}
}