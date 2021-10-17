package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	//here create some database and golang
	ID 											primitive.ObjectID	`bson:"_id"`
	First_name 							*string							`json:"first_name" validate:"required, min=2, max=255"`
	Last_name 							*string							`json:"last_name" validate:"required, min=2, max=255"`
	Password  							*string							`json:"password" validate:"required, min=2, max=255"`
	Email 									*string							`json:"email" validate:"email, required"`
	Phone 									*string							`json:"phone" validate:"required"`
	Token 									*string							`json:"token"`
	User_type 							*string							`json:"user_type" validate:"required, eq=ADMIN|eq=USER"`
	Refresh_token 					*string							`json:"refresh_token"`
	created_at 							*time.Time					`json:"updated_at"`
	updated_at 							*time.Time					`json:"created_at"`
}