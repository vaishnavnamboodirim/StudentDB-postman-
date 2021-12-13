package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Student struct {
	Id    string  `json:"id"`
	Name  string  `json:"name"`
	Class string  `json:"class"`
	Marks float64 `json:"marks"`
}

var students []Student

func getStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)

}

func getStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range students {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Student{})
}

func createStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var student Student
	_ = json.NewDecoder(r.Body).Decode(&student)
	student.Id = strconv.Itoa(3)
	students = append(students, student)
	json.NewEncoder(w).Encode(student)
}

func updateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, items := range students {
		if items.Id == params["id"] {
			students = append(students[:index], students[index+1:]...)
			var student Student
			_ = json.NewDecoder(r.Body).Decode(&student)
			student.Id = params["id"]
			students = append(students, student)
			json.NewEncoder(w).Encode(student)
			return
		}
	}
	json.NewEncoder(w).Encode(students)
}

func deleteStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, items := range students {
		if items.Id == params["id"] {
			students = append(students[:index], students[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(students)
	log.Println("Student deleted from DB")
}

func main() {

	r := mux.NewRouter()

	students = append(students, Student{Id: "1", Name: "Student1", Class: "5A", Marks: 500})
	students = append(students, Student{Id: "2", Name: "Student2", Class: "5A", Marks: 450})

	r.HandleFunc("/api/students/", getStudents).Methods("GET")
	r.HandleFunc("/api/students/{id}", getStudent).Methods("GET")
	r.HandleFunc("/api/students", createStudent).Methods("POST")
	r.HandleFunc("/api/students/{id}", updateStudent).Methods("PUT")
	r.HandleFunc("/api/students/{id}", deleteStudent).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))

}
