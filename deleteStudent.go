package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DeleteStudent(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)["id"] //get Parameter Value as string

	_id, err := primitive.ObjectIDFromHex(params) //convert params to mongoDB hex id

	if err != nil {
		fmt.Printf(err.Error())
	}

	opts := options.Delete().SetCollation(&options.Collation{}) // to //specify language-specific rules for string comparison, such as //rules for lettercase

	res, err := userCollection.DeleteOne(context.TODO(), bson.D{{"_id", _id}}, opts)

	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(res.DeletedCount) // return no of Douments deleted
}
