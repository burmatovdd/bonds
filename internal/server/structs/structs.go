package structs

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Bond для хранения облигаций в базе и для работы с ней
type Bond struct {
	Name    string   `json:"name"`
	Coupons []Coupon `json:"coupons"`
}

// Coupon для работы с купонами
type Coupon struct {
	Date  string  `json:"date"`
	Value float64 `json:"value"`
}

type UserToken struct {
	Token         string `json:"token"`
	Refresh_token string `json:"refresh_token"`
}

type UserBonds struct {
	Bond  Bond `json:"bond"`
	Count int  `json:"count"`
}

type User struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      string             `json:"name" validate:"required,min=2,max=100"`
	Password  string             `json:"password" validate:"required,min=8"`
	Login     string             `json:"login" validate:"required, min=5"`
	User_id   string             `json:"user_id"`
	User_type string             `json:"user_type" validate:"required,eq=ADMIN|eq=USER"`
	Bonds     []UserBonds        `json:"bonds"`
}
