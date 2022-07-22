package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	FirstName string             `json:"firstName" validate:"required"`
	LastName  string             `json:"lastName"`
	Email     string             `json:"email" validate:"email,required"`
	Password  string             `json:"password" validate:"required"`
}

type Post struct {
	Id          primitive.ObjectID   `json:"id,omitempty" bson:"_id,omitempty"`
	Title       string               `json:"title,omitempty" validate:"required"`
	Description string               `json:"description,omitempty"`
	User        primitive.ObjectID   `json:"user,omitempty"`
	Email       string               `json:"email,omitempty" validate:"email"`
	Address     string               `json:"address,omitempty" validate:"required"`
	PhoneNumber string               `json:"phoneNumber,omitempty" validate:"required"`
	ImageUrls   []string             `json:"imageUrls"`
	Status      string               `json:"status,omitempty" validate:"required,eq=LOST|eq=FOUND"`
	Claims      []primitive.ObjectID `json:"claims"`
	CreatedAt   time.Time            `json:"createdAt,omitempty"`
}

type Claim struct {
	Id          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title,omitempty" validate:"required"`
	Description string             `json:"description,omitempty"`
	User        primitive.ObjectID `json:"user,omitempty"`
	Email       string             `json:"email,omitempty" validate:"email"`
	ImageUrls   []string           `json:"imageUrls"`
	Address     string             `json:"address,omitempty" validate:"required"`
	PhoneNumber string             `json:"phoneNumber,omitempty" validate:"required"`
	Post        primitive.ObjectID `json:"post,omitempty" validate:"required"`
	CreatedAt   time.Time          `json:"createdAt,omitempty"`
	Status      string             `json:"status,omitempty" validate:"eq=PENDING|eq=ACCEPTED|eq=REJECTED"`
}
