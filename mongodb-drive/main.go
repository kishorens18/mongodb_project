package main

import (
	"fmt"
	reastaurants "mongodb/reastaurants_service"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	mongoClient *mongo.Client
)

func main() {
	//client, _ := config.ConnectDataBase()
	//collection := config.GetCollection(client, "sample_training", "zips")

	fmt.Println("MongoDB server started")

	// products := []interface{}{
	// 	models.Product{ID: primitive.NewObjectID(), Name: "samsung", Price: 1000000, Description: "wow"},
	// 	models.Product{ID: primitive.NewObjectID(), Name: "iphone", Price: 9000000, Description: "oh,my god"},
	// 	models.Product{ID: primitive.NewObjectID(), Name: "poco", Price: 11000000, Description: "god itself"},
	// }

	// services.InsertProductList(products)

	products, _ := reastaurants.FindRestaurant()
	for index, product := range products {
		fmt.Println(index, product)
	}

}
