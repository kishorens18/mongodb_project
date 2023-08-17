package services

import (
	"fmt"
	"mongodb/config"
	"mongodb/models"
	"time"

	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func ProductContext() *mongo.Collection {
	client, _ := config.ConnectDataBase()
	return config.GetCollection(client, "product", "products")
}

func InsertProduct(product models.Product) (*mongo.InsertOneResult, error) {
	// var product models.Product
	// product.ID = primitive.NewObjectID()
	// product.Name = "iphone"
	// product.Description = "It working"
	// product.Price = 100000
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// client, _ := config.ConnectDataBase()
	// var productCollection *mongo.Collection = config.GetCollection(client, "inventory", "products")
	result, err := ProductContext().InsertOne(ctx, product)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	return result, nil
}

func InsertProductList(products []interface{}) (*mongo.InsertManyResult, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, err := ProductContext().InsertMany(ctx, products)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	return result, nil
}

func Update(products []interface{}) (*mongo.InsertManyResult, error) {
	// var product models.Product
	// product.ID = primitive.NewObjectID()
	// product.Name = "iphone"
	// product.Description = "It working"
	// product.Price = 100000
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// client, _ := config.ConnectDataBase()
	// var productCollection *mongo.Collection = config.GetCollection(client, "inventory", "products")
	result, err := ProductContext().InsertMany(ctx, products)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	return result, nil
}

func FindProducts() ([]*models.Product, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel() // Cancel the context to release resources even if the function returns early
	filter := bson.D{{"name", "samsung"}}

	result, err := ProductContext().Find(ctx, filter)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err // Return the error here
	}

	var products []*models.Product
	for result.Next(ctx) {
		product := &models.Product{}
		err := result.Decode(product)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err := result.Err(); err != nil {
		return nil, err
	}

	if len(products) == 0 {
		return []*models.Product{}, nil
	}

	return products, nil
}
