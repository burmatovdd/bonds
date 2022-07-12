package UserModel

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"modules/internal/server/bondsInfo"
)

type User struct {
	ID        primitive.ObjectID    `bson:"_id"`
	Name      string                `json:"name" validate:"required,min=2,max=100"`
	Password  string                `json:"password" validate:"required,min=8"`
	Login     string                `json:"login" validate:"required, min=5"`
	User_id   string                `json:"user_id"`
	User_type string                `json:"user_type" validate:"required,eq=ADMIN|eq=USER"`
	Bonds     []bondsInfo.UserBonds `json:"bonds"`
}
