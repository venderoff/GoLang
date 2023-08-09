package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Id   string `json:"_id"`
	Name string `json:"name"`
	Mail string `json:"mail"`
}

var userCollection = db().Database("spring").Collection("student")

func createStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var person User

	err := json.NewDecoder(r.Body).Decode(&person) // storing person

	if err != nil {
		fmt.Println("Error in parsing Json")
	}

	insertCollection, err := userCollection.InsertOne(context.TODO(), person)
	if err != nil {
		log.Println("Problem in insertion", err, insertCollection)
	}

	// json.NewEncoder(w).Encode(insertCollection.) // return mongoDB

	json.NewEncoder(w).Encode(insertCollection) // return mongoDB

}
