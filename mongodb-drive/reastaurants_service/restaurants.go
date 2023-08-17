package reastaurants

import (
	"context"
	"fmt"
	"mongodb/config"
	models "mongodb/rest_models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ProductContext() *mongo.Collection {
	client, _ := config.ConnectDataBase()
	return config.GetCollection(client, "sample_restaurants", "restaurants")
}

func FindRestaurant() ([]*models.Restaurant, error) {
	ctx, _ := context.WithTimeout(context.Background(), 100*time.Second)
	filter := bson.D{} //here condition should be given
	limit := options.Find().SetLimit(10)
	result, err := ProductContext().Find(ctx, filter, limit)
	if err != nil {

		fmt.Println(err.Error())
		return nil, err
	} else {
		// fmt.Println(result)
		//build the array of products for the cursor that we received
		var products []*models.Restaurant
		for result.Next(ctx) {
			product := &models.Restaurant{}
			err := result.Decode(product)

			if err != nil {
				return nil, err
			}
			//fmt.Println(product)
			products = append(products, product)
		}
		if err := result.Err(); err != nil {
			return nil, err
		}
		if len(products) == 0 {
			return []*models.Restaurant{}, nil
		}
		return products, nil
	}
}
