package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type WebsiteData struct {
	URL     string `bson:"url"`
	Content string `bson:"content"`
	Title   string `bson:"title"`
}

func main() {
	// Create a MongoDB client
	clientOptions := options.Client().ApplyURI("mongodb://root:example@localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	// Select a database and collection
	collection := client.Database("local").Collection("student")

	// create group stage
	// groupStage := bson.D{
	// 	{"$group", bson.D{
	// 		{"_id", "$last_name"},
	// 		{"averageGrade", bson.D{{"$avg", "$grade"}}},
	// 		{"Grade", bson.D{{"$push", "$grade"}}},
	// 		{"firstNames", bson.D{{"$push", "$first_name"}}},
	// 		{"lastNames", bson.D{{"$first", "$last_name"}}},
	// 		{"Quantity", bson.D{{"$count", bson.D{}}}},
	// 	}}}

	groupStage := bson.M{
		"$group": bson.M{
			"_id":          "$last_name",
			"averageGrade": bson.M{"$avg": "$grade"},
			"Grade":        bson.M{"$push": "$grade"},
			"firstNames":   bson.M{"$push": "$first_name"},
			"lastNames":    bson.M{"$first": "$last_name"},
			"Quantity":     bson.M{"$count": bson.M{}},
		}}.project({ item: 1, status: 1 });

	// pass the pipeline to the Aggregate() method
	cursor, err := collection.Aggregate(context.TODO(), bson.A{groupStage})
	if err != nil {
		panic(err)
	}

	// display the results
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	// for _, result := range results {
	// 	fmt.Printf("Average Grade: %v \n", result["averageGrade"])
	// 	fmt.Printf("Grade: %v \n", result["Grade"])
	// 	fmt.Printf("First names: %v \n", result["firstNames"])
	// 	fmt.Printf("Last names: %v \n", result["lastNames"])
	// 	fmt.Printf("Quantity: %v \n\n", result["Quantity"])
	// }

	message, _ := json.MarshalIndent(results, "", " ")
	fmt.Println(string(message))

}
