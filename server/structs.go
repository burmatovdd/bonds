package main

import (
	"encoding/xml"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Document struct {
	XMLName xml.Name `xml:"document"`
	Rows    []Row    `xml:"data>rows>row"`
}

type Row struct {
	XMLName xml.Name `xml:"row" json:"-"`
	Date    string   `xml:"coupondate,attr" json:"date"`
	Value   string   `xml:"value,attr" json:"value"`
}

type Bond struct {
	Name  string  `json:"name"`
	Count float64 `json:"count"`
}

type AllInfo struct {
	BondInfo []BondInfo `json:"bondInfos"`
	Months   []Months   `json:"months"`
}

type UserToken struct {
	Token         string `json:"token"`
	Refresh_token string `json:"refresh_token"`
}

type User struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      string             `json:"name" validate:"required,min=2,max=100"`
	Password  string             `json:"password" validate:"required,min=8""`
	Login     string             `json:"login" validate:"required, min=5"`
	User_id   string             `json:"user_id"`
	User_type string             `json:"user_type" validate:"required,eq=ADMIN|eq=USER""`
	Bond      []Bond             `json:"bond"`
}

type BondInfo struct {
	Bond    Bond     `json:"bond"`
	Coupons []Coupon `json:"coupons"`
}

type Coupon struct {
	Date  string  `json:"date"`
	Value float64 `json:"value"`
}

type Months struct {
	Date  string  `json:"date"`
	Value float64 `json:"value"`
}

type Test struct {
	Coupons AllInfo `json:"coupons"`
}
