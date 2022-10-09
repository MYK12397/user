package storage

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/google/martian/v3/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	userCollection = "users"
)

// User represents user object in request body.
type User struct {
	ID             string     `bson:"id"`
	Created        *time.Time `bson:"created"`
	Modified       *time.Time `bson:"modified,omitempty"`
	FirstName      string     `bson:"firstName"`
	LastName       string     `bson:"lastName"`
	ProfilePicture string     `bson:"profilePicture"`
	PhoneNumber    string     `bson:"phoneNumber"`
	Email          string     `bson:"email"`
	Country        string     `bson:"country"`
	DateOfBirth    string     `bson:"dateOfBirth,omitempty"`
}

// Save inerts a new document for user in database.
func Save(user *User) error {

	client := GetClient()

	collection := client.Database(os.Getenv("DB_NAME")).Collection(userCollection)
	ctime := time.Now()
	user.ID = primitive.NewObjectID().Hex()
	user.Created = &ctime
	user.Modified = &ctime

	_, err := collection.InsertOne(context.TODO(), user)

	if err != nil {
		log.Errorf("user not create in database", http.StatusInternalServerError)
		return err
	}

	return nil
}
