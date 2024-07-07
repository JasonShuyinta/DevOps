package services

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	"my-go-api/pkg/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func UpdateUser(client *mongo.Client, user models.User) error {
	collection := client.Database(databaseName).Collection(collectionName)
	objectID := user.ID

	//Update document
	update := bson.M{
		"$set": bson.M{
			"username":        user.Username,
			"email":           user.Email,
			"password":        user.Password,
			"dateOfBirth":     user.DateOfBirth,
			"updateTimestamp": time.Now(),
		},
	}

	result, err := collection.UpdateOne(context.TODO(), bson.M{"_id": objectID}, update)
	if err != nil {
		return err
	}

	// Check if any document was modified
	if result.MatchedCount == 0 {
		return err
	}

	return nil

}

// Create MultipleUsers
func CreateMultipleUsers(client *mongo.Client, n int) error {
	collection := client.Database(databaseName).Collection(collectionName)
	var operations []mongo.WriteModel

	for i := 0; i < n; i++ {
		username := "john_doe_" + strconv.Itoa(i)
		email := "john.doe" + strconv.Itoa(i) + "@example.com"
		user := models.User{
			Username:          username,
			Email:             email,
			Password:          "password",
			DateOfBirth:       time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC),
			CreationTimestamp: time.Now(),
			UpdateTimestamp:   time.Now(),
		}
		operations = append(operations, mongo.NewInsertOneModel().SetDocument(user))
	}

	_, err := collection.BulkWrite(context.TODO(), operations, options.BulkWrite().SetOrdered(false))
	if err != nil {
		return err
	}
	return nil

}

// ChunkSize defines the size of each chunk for parallel processing

// CreateMultipleUsers creates multiple users in parallel.
func CreateMultipleUsersGoRoutine(client *mongo.Client, n int) error {
	collection := client.Database(databaseName).Collection(collectionName)
	var ChunkSize = n / 4
	var wg sync.WaitGroup
	errorChannel := make(chan error, n/ChunkSize)

	// Function to handle user creation for a chunk of users
	createChunk := func(start, end int) {
		defer wg.Done()
		var operations []mongo.WriteModel
		for i := start; i < end; i++ {
			username := "john_doe_" + strconv.Itoa(i)
			email := "john.doe" + strconv.Itoa(i) + "@example.com"
			user := models.User{
				Username:          username,
				Email:             email,
				Password:          "password",
				DateOfBirth:       time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC),
				CreationTimestamp: time.Now(),
				UpdateTimestamp:   time.Now(),
			}
			operations = append(operations, mongo.NewInsertOneModel().SetDocument(user))
		}

		// Perform the bulk write operation
		_, err := collection.BulkWrite(context.TODO(), operations, options.BulkWrite().SetOrdered(false))
		if err != nil {
			errorChannel <- fmt.Errorf("error creating users from %d to %d: %w", start, end, err)
		}
	}

	// Create chunks and start goroutines
	for i := 0; i < n; i += ChunkSize {
		start := i
		end := i + ChunkSize
		if end > n {
			end = n
		}
		wg.Add(1)
		go createChunk(start, end)
	}

	// Wait for all goroutines to complete
	wg.Wait()
	close(errorChannel)

	// Check for errors from any goroutine
	for err := range errorChannel {
		if err != nil {
			return fmt.Errorf("error in user creation: %w", err)
		}
	}

	return nil
}

// Count Users
func CountUsers(client *mongo.Client) (int64, error) {
	collection := client.Database(databaseName).Collection(collectionName)
	result, err := collection.CountDocuments(context.TODO(), bson.D{})
	if err != nil {
		return 0, err
	}
	return result, nil
}

func DeleteAllUsers(client *mongo.Client) (*mongo.DeleteResult, error) {
	collection := client.Database(databaseName).Collection(collectionName)
	result, err := collection.DeleteMany(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	return result, nil
}
