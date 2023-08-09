package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UpdateStudent(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	type updateBody struct {
		Name string `json:"name"`
		Mail string `json:"mail`
	}
	var body updateBody
	e := json.NewDecoder(r.Body).Decode(&body)
	if e != nil {
		fmt.Println(e)
	}

	filter := bson.D{{"name", body.Name}} // converting value to BSON
	after := options.After                // returning updated document

	returnOpt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}

	update := bson.D{{"$set", bson.D{{"mail", body.Mail}}}}

	updateResult := userCollection.FindOneAndUpdate(context.TODO(), filter, update, &returnOpt)

	var result primitive.M

	_ = updateResult.Decode(&result)

	json.NewEncoder(w).Encode(result)
}
