package services

import (
	"context"
	"log"
	"time"

	"my-go-api/pkg/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	databaseName   = "go-db"
	collectionName = "user"
)

func GetAllUsers(client *mongo.Client) ([]models.User, error) {
	collection := client.Database(databaseName).Collection(collectionName)
	cur, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	var results []models.User
	for cur.Next(context.TODO()) {
		var result models.User
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, result)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func SaveUser(client *mongo.Client, user models.User) error {
	collection := client.Database(databaseName).Collection(collectionName)
	user.CreationTimestamp = time.Now()
	user.UpdateTimestamp = time.Now()

	_, err := collection.InsertOne(context.TODO(), user)
	return err
}

func GetUserByID(client *mongo.Client, id string) (*models.User, error) {
	collection := client.Database(databaseName).Collection(collectionName)
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var user models.User
	err = collection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func DeleteUserByID(client *mongo.Client, id string) error {
	collection := client.Database(databaseName).Collection(collectionName)
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = collection.DeleteOne(context.TODO(), bson.M{"_id": objectID})
	return err
}
