package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	route := mux.NewRouter()
	db()
	s := route.PathPrefix("/api").Subrouter()

	//routes
	// Post : http://localhost:8000/api/createStudent
	s.HandleFunc("/createStudent", createStudent).Methods("POST")
	// Get : http://localhost:8000/api/getStudent/zara
	s.HandleFunc("/getStudent/{name}", GetStudents).Methods("GET")
	// http://localhost:8000/api/updateStudent
	s.HandleFunc("/updateStudent", UpdateStudent).Methods("PUT")

	s.HandleFunc("/delStudent/{id}", DeleteStudent).Methods("DELETE")
	log.Println("Server Running on PORT : 8000")
	log.Fatal(http.ListenAndServe(":8000", s)) // Run Server

}
