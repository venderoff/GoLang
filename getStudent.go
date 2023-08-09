package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetStudents(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	//path Variable
	pathVar := mux.Vars(r)

	for _, p := range pathVar {
		var result primitive.M

		err1 := userCollection.FindOne(context.TODO(), bson.D{{"name", p}}).Decode(&result)
		if err1 != nil {
			fmt.Println(err1)
		}

		json.NewEncoder(w).Encode(result)
	}

}
