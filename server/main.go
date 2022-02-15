package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"strconv"
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

type BondInfo struct {
	Bond Bond `json:"bond"`
	Coupons []Coupon `json:"coupons"`
}

type Coupon struct {
	Date string `json:"date"`
	Value float64 `json:"value"`
}

type Test struct {
	Coupons []BondInfo `json:"coupons"`
}

var input string
var data []Row

var bonds = [7]Bond{
	{
		"RU000A103UL3",
		2,
	},
	{"RU000A1049P5", 2},
	{"RU000A1040V2", 2},
	{"RU000A103WD6", 2},
	{"RU000A104CJ3", 1},
	{"RU000A1043G7", 1},
	{"RU000A104FG2", 2},
}
var monthDict = make(map[string]float64)

func TakeData(year string) Test {
	var bondInfos []BondInfo
	var yearSum float64

	for i := 0; i < len(bonds); i++ {
		var bondInfo = BondInfo {
			Bond{
				bonds[i].Name,
				bonds[i].Count,
			},
			[]Coupon{},
		}
		fmt.Println("Name" + "        |" + " Value" + " | " + "Count " + " |" + "Date" + "      |")
		response, err := http.Get("https://iss.moex.com/iss/statistics/engines/stock/markets/bonds/bondization/" + bonds[i].Name + "?iss.meta=off&iss.only=coupons&coupons.columns=coupondate,value")
		if err != nil {
			log.Fatal(err)
		}

		// defer очищаем ресуры
		defer response.Body.Close()
		byteValue, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		data := &Document{}

		err = xml.Unmarshal(byteValue, &data)
		if nil != err {
			fmt.Println("Error unmarshalling from XML", err)

		}

		result, err := json.Marshal(&data.Rows)
		if nil != err {
			fmt.Println("Error marshalling to JSON", err)
		}

		var dict []map[string]string
		if err = json.Unmarshal(result, &dict); err != nil {
			panic(err)
		}

		for idx := range dict {

			date := (dict[idx]["date"])[:4]
			month := (dict[idx]["date"])[5:7]
			value, _ := strconv.ParseFloat(dict[idx]["value"], 64)

			if year != "any" {
				if date == year {
					if month != "" {
						value = math.Round((value*bonds[i].Count)*100) / 100
						yearSum = math.Round((yearSum+value)*100) / 100
						fmt.Println(bonds[i].Name+"|", value, "|"+"  ", bonds[i].Count, "  |", dict[idx]["date"])
						// проверяем есть ли такой ключ
						_, exist := monthDict[month]
						// если да, то...
						if exist {
							monthDict[month] += value // month:value
							var coupon = Coupon{
								Date: dict[idx]["date"],
								Value: value,
							}
							bondInfo.Coupons = append(bondInfo.Coupons,coupon )
						} else {
							monthDict[month] = value // month:value
							var coupon = Coupon{
								Date: dict[idx]["date"],
								Value: value,
							}
							bondInfo.Coupons = append(bondInfo.Coupons, coupon)						}
					}
				}
			} else {
				value = value * bonds[i].Count
				yearSum = yearSum + value
				fmt.Println(bonds[i].Name+"|", value, "|"+"  ", bonds[i].Count, "  |")
			}
		}
		bondInfos = append(bondInfos,bondInfo)
	}
	tmpl := Test{
		bondInfos,
	}
	return tmpl
}

func Post(c *gin.Context) {
	jsonData, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Println(string(jsonData))

	yearMap := make(map[string]string)
	err := json.Unmarshal(jsonData, &yearMap)
	if err != nil {
		fmt.Println("err: ", err)
	}

	var year string
	for _, value := range yearMap {
		year = string(value)
	}

	test2 := TakeData(year)
	c.JSON(200, struct {
		BondInfos []BondInfo `json:"bondInfos"`
	}{
		test2.Coupons,
	})
}

func HandleRequest() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"Content-Type,access-control-allow-origin, access-control-allow-headers"},
	}))
	router.POST("/result", Post)
	router.Run("localhost:8080")
}

func main() {
	HandleRequest()
}
