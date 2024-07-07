package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"` // Let MongoDB handle the ID generation
	Username          string             `bson:"username" json:"username"`
	Email             string             `bson:"email" json:"email"`
	Password          string             `bson:"password" json:"password"`
	DateOfBirth       time.Time          `bson:"dateOfBirth" json:"dateOfBirth"`
	CreationTimestamp time.Time          `bson:"creationTimestamp" json:"creationTimestamp"`
	UpdateTimestamp   time.Time          `bson:"updateTimestamp" json:"updateTimestamp"`
}
