package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var courses []Course

func main() {
	router := mux.NewRouter()

	courses = append(courses, Course{ID: "124134", Name: "FullStack Django Developer Freelance ready", Price: "299", Link: "https://courses.learncodeonline.in/learn/FullStack-Django-Developer-Freelance-ready",
		Author: &Author{Firstname: "Hitesh", Lastname: "Choudhary"}})
	courses = append(courses, Course{ID: "154434", Name: "Full stack with Django and React", Price: "299", Link: "https://courses.learncodeonline.in/learn/Full-stack-with-Django-and-React",
		Author: &Author{Firstname: "Hitesh", Lastname: "Choudhary"}})
	courses = append(courses, Course{ID: "198767", Name: "Complete React Native bootcamp", Price: "199", Link: "https://courses.learncodeonline.in/learn/Complete-React-Native-Mobile-App-developer",
		Author: &Author{Firstname: "Hitesh", Lastname: "Choudhary"}})

	router.HandleFunc("/api/courses", getCourses).Methods("GET")
	router.HandleFunc("/api/course/{id}", getSingleCourse).Methods("GET")
	router.HandleFunc("/api/course/create", createCourse).Methods("POST")
	router.HandleFunc("/api/course/update/{id}", updateCourse).Methods("PUT")
	router.HandleFunc("/api/course/delete/{id}", deleteCourse).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type Course struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Price  string  `json:"price"`
	Link   string  `json:"link"`
	Author *Author `json:"author"`
}

func getCourses(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(courses)
}

func getSingleCourse(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)

	for _, item := range courses {
		if item.ID == params["id"] {
			json.NewEncoder(res).Encode(item)
			return
		}
	}

	json.NewEncoder(res).Encode("No course found")
}

func createCourse(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	var course Course

	_ = json.NewDecoder(req.Body).Decode(&course)
	course.ID = strconv.Itoa(rand.Intn(1000000))

	courses = append(courses, course)
	json.NewEncoder(res).Encode(course)
}

func deleteCourse(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)

	for i, item := range courses {
		if item.ID == params["id"] {
			courses = append(courses[:i], courses[i+1:]...)
			break
		}
	}
	json.NewEncoder(res).Encode(courses)
}

func updateCourse(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)

	for i, item := range courses {
		if item.ID == params["id"] {
			courses = append(courses[:i], courses[i+1:]...)
			var course Course

			_ = json.NewDecoder(req.Body).Decode(&course)
			course.ID = params["id"]
			courses = append(courses, course)
			json.NewEncoder(res).Encode(course)
			return
		}
	}
}
