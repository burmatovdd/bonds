package main

import "encoding/xml"

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
