package repositories

import (
	"context"
	"errors"
	"log"
	"strings"
	"time"

	"github.com/abe16s/Online-Marketplace-API/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


type UserRepository struct {
	collection *mongo.Collection
}

// NewUserRepository creates a new UserRepository.
func NewUserRepository(client *mongo.Client, dbName string, collectionName string) *UserRepository {
	collection := client.Database(dbName).Collection(collectionName)

	// check if there is an index on the email field

	// Get a list of existing indexes
    cursor, err := collection.Indexes().List(context.TODO())
    if err != nil {
        log.Printf("could not list indexes: %v", err)
    }
    defer cursor.Close(context.TODO())

    var indexes []bson.M
    if err := cursor.All(context.TODO(), &indexes); err != nil {
        log.Printf("could not parse indexes: %v", err)
    }

    // Check if the "email" index already exists
    indexExists := false
    for _, index := range indexes {
        key := index["key"].(bson.M)
        if len(key) == 1 && key["email"] != nil {
            indexExists = true
            break
        }
    }

    // If the index does not exist, create it
    if !indexExists {
		indexModel := mongo.IndexModel{
			Keys:    bson.D{{Key: "email", Value: 1}}, // Create index on the "email" field
			Options: options.Index().SetUnique(true),    // Ensure the index is unique
		}
		
		// Create the index
		_, err = collection.Indexes().CreateOne(context.TODO(), indexModel)
		if err != nil {
			log.Printf("could not create index: %v", err)
		}
	} else {
		log.Println("email index already exists")
	}
	
	return &UserRepository{
		collection: collection,
	}
}


// register new user with unique email and password
func (ur *UserRepository) CreateUser(user *models.User) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Check if user already exists
	for {
		_, err := ur.collection.InsertOne(ctx, user)
		
		// if user exists return error
		if mongo.IsDuplicateKeyError(err) {
			// Check if the duplicate key error is caused by the email field
			if strings.Contains(err.Error(), "email") {
				return nil, errors.New("email already exists")
			}
			// Check if the duplicate key error is caused by the _id field
			if strings.Contains(err.Error(), "_id") {
				continue
			}
		} else if err != nil {
			return nil, err
		}

		// else create new user
		return user, nil
	}
}


// login user 
func (ur *UserRepository) FindByEmail(email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var existingUser models.User
	err := ur.collection.FindOne(ctx, bson.M{"email": email}).Decode(&existingUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &existingUser, nil
}


