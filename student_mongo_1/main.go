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

type Student struct {
	FirstName string  `bson:"first_name"`
	LastName  string  `bson:"last_name"`
	Grade     float64 `bson:"grade"`
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

	// Define your aggregation pipeline with $group and $project stages
	pipeline := []bson.M{
		{
			"$group": bson.M{
				"_id":          "$last_name",
				"averageGrade": bson.M{"$avg": "$grade"},
				"Grade":        bson.M{"$push": "$grade"},
				"firstNames":   bson.M{"$push": "$first_name"},
				"lastNames":    bson.M{"$first": "$last_name"},
				"Quantity":     bson.M{"$sum": 1},
			},
		},
		{
			"$project": bson.M{
				"_id":          0,
				"averageGrade": 1,
				"Grade":       "$Grade",
				"firstNames":   "$firstNames",
				"lastNames":    "$lastNames",
				"Quantity":     "$Quantity",
				"is_passed": bson.M{
					"$cond": bson.M{
						"if":   bson.M{"$gte": []interface{}{"$averageGrade", 7}},
						"then": "Passed",
						"else": "Failed",
					},
				},
			},
		},
	}

	// Pass the pipeline to the Aggregate() method
	cursor, err := collection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		panic(err)
	}

	// Display the results
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	message, _ := json.MarshalIndent(results, "", " ")
	fmt.Print(string(message))
	fmt.Println()
}
