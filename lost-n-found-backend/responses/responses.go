package responses

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HttpResponse struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

type UserResponse struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	FirstName string             `json:"firstName" validate:"required"`
	LastName  string             `json:"lastName"`
	Email     string             `json:"email" validate:"email,required"`
	AuthToken string             `json:"authToken" validate:"required"`
}
type PostResponse struct {
	Id          primitive.ObjectID   `json:"id,omitempty" bson:"_id,omitempty"`
	Title       string               `json:"title,omitempty" validate:"required"`
	Description string               `json:"description,omitempty"`
	Email       string               `json:"email,omitempty" validate:"email"`
	Address     string               `json:"address,omitempty" validate:"required"`
	PhoneNumber string               `json:"phoneNumber,omitempty" validate:"required"`
	ImageUrls   []string             `json:"imageUrls"`
	Status      string               `json:"status,omitempty" validate:"required,eq=LOST|eq=FOUND"`
	Claims      []primitive.ObjectID `json:"claims"`
	CreatedAt   time.Time            `json:"createdAt,omitempty"`
	UserDetails UserResponse         `json:"userDetails,omitempty"`
}
type ClaimResponse struct {
	Id          primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Title       string             `json:"title,omitempty" validate:"required"`
	Description string             `json:"description,omitempty"`
	Email       string             `json:"email,omitempty" validate:"email"`
	ImageUrls   []string           `json:"imageUrls"`
	Address     string             `json:"address,omitempty" validate:"required"`
	PhoneNumber string             `json:"phoneNumber,omitempty" validate:"required"`
	CreatedAt   time.Time          `json:"createdAt,omitempty"`
	UserDetails UserResponse       `json:"userDetails,omitempty"`
	Post        primitive.ObjectID `json:"post,omitempty"`
	Status      string             `json:"status,omitempty" validate:"eq=PENDING|eq=ACCEPTED|eq=REJECTED"`
}
