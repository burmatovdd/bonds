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
	"sort"
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

type BondInfo struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
	Count float64 `json:"count"`
	Date  string  `json:"date"`
}

type InfoPerYear struct {
	Year  string  `json:"year"`
	Month string  `json:"month"`
	Value float64 `json:"value"`
}

type Bond struct {
	Name  string  `json:"name"`
	Count float64 `json:"count"`
}

var input string
var data []Row
var all []InfoPerYear
var info []BondInfo
var bonds = [7]Bond{
	//{"RU000A102RX6", 1},
	//{"RU000A1007Z2", 1},
	//{"RU000A0ZYCR1", 1},
	//{"RU000A1041B2", 1},
	//{"RU000A0ZZZ17", 1},
	//{"RU000A103F27", 1},// 9,42%

	{"RU000A103UL3", 2},
	{"RU000A1049P5", 2},
	{"RU000A1040V2", 2},
	{"RU000A103WD6", 2},
	{"RU000A104CJ3", 1},
	{"RU000A1043G7", 1},
	{"RU000A104FG2", 2},
}
var monthDict = make(map[string]float64)

func TakeData(year string) {
	var yearSum float64

	for i := 0; i < len(bonds); i++ {
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
			return
		}

		result, err := json.Marshal(&data.Rows)
		if nil != err {
			fmt.Println("Error marshalling to JSON", err)
			return
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
							info = append(info, BondInfo{Name: bonds[i].Name, Value: value, Count: bonds[i].Count, Date: dict[idx]["date"]})
						} else {
							monthDict[month] = value // month:value
							info = append(info, BondInfo{Name: bonds[i].Name, Value: value, Count: bonds[i].Count, Date: dict[idx]["date"]})
						}
					}
				}
			} else {
				value = value * bonds[i].Count
				yearSum = yearSum + value
				fmt.Println(bonds[i].Name+"|", value, "|"+"  ", bonds[i].Count, "  |")
			}
		}
	}
	keys := make([]string, 0, len(monthDict))
	for key := range monthDict {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		all = append(all, InfoPerYear{Year: year, Month: key, Value: monthDict[key]})
		fmt.Println(year+" -", key, ":", monthDict[key])
	}

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
	fmt.Println("year: ", year)
	TakeData(year)
	c.JSON(200, struct {
		InfoPerYear []InfoPerYear `json:"info_per_year"`
		BondInfo    []BondInfo    `json:"bond_info"`
	}{
		all, info,
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
