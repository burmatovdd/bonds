package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
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

type Info struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
	Count float64 `json:"count"`
	Date  string  `json:"date"`
}

type AllInfo struct {
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
var all []AllInfo
var info []Info
var bonds = [8]Bond{
	{"RU000A100YG1", 10},
	{"RU000A100LV7", 5},
	{"RU000A100YK3", 2},
	{"RU000A100YG1", 3},
	{"RU000A0ZYU39", 19},
	{"RU000A0ZYL22", 6},
	{"RU000A103DT2", 1},
	{"SU26229RMFS3", 11},
}
var monthDict = make(map[string]float64)

func TakeData(year string) {
	var yearSum float64
	//var monthDict = make(map[string]float64)
	//monthDict := map[string]float64{}

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
		//var monthArr = make([]float64,0)
		if err := json.Unmarshal(result, &dict); err != nil {
			panic(err)
		}

		for idx := range dict {

			date := (dict[idx]["date"])[:4]
			month := (dict[idx]["date"])[5:7]
			value, err := strconv.ParseFloat(dict[idx]["value"], 64)

			if err != nil {
				fmt.Println(err)
			}

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
							info = append(info, Info{Name: bonds[i].Name, Value: value, Count: bonds[i].Count, Date: dict[idx]["date"]})
						} else {
							monthDict[month] = value // month:value
							info = append(info, Info{Name: bonds[i].Name, Value: value, Count: bonds[i].Count, Date: dict[idx]["date"]})
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
		all = append(all, AllInfo{Year: year, Month: key, Value: monthDict[key]})
		fmt.Println(year+" -", key, ":", monthDict[key])
	}
	fmt.Println("amount per year : ", yearSum, "rub")

	fmt.Println(monthDict)
}

func HandleRequest() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/home_page.html")
	})
	http.HandleFunc("/result", func(w http.ResponseWriter, r *http.Request) {

		input = r.FormValue("year")

		fmt.Fprintf(w, "Год: %s ", input)
		fmt.Println("year: ", input)
		TakeData(input)
		for key, value := range monthDict {
			fmt.Fprintf(w, " Дата: %s Значение: %f", input+"-"+key, value)
		}
	})
	fmt.Println("Server is listening...")
	http.ListenAndServe(":8080", nil)
}

func main() {

	HandleRequest()
}
