package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"strconv"
)

type Document struct {
	XMLName xml.Name `xml:"document"`
	Rows    []Row   `xml:"data>rows>row"`
}

type Row struct {
	XMLName xml.Name `xml:"row" json:"-"`
	Date    string   `xml:"coupondate,attr" json:"date"`
	Value    string   `xml:"value,attr" json:"value"`
}

type Info struct {
	Name string `json:"name"`
	Date string `json:"info"`
	Value float64 `json:"value"`
	YearSum float64 `json:"yearSum"`
}

type Bond struct {
	Name string `json:"name"`
	Count float64 `json:"count"`
}

var data []Row
var tmpl []Info
var bonds = [8]Bond{
	{"RU000A100YG1",10},
	{"RU000A100LV7",5},
	{"RU000A100YK3",1},
	{"RU000A100YG1",1},
	{"RU000A0ZYU39",1},
	{"RU000A0ZYL22",1},
	{"RU000A103DT2",1},
	{"SU26229RMFS3",1},
}

func TakeData(year string){
	var yearSum float64
	//var monthDict = make(map[string]float64)
	monthDict := map[string]float64{}

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

			if year != "any"{
				if date == year {
					if month != ""{
						value = math.Round((value * bonds[i].Count)*100)/100
						yearSum = math.Round((yearSum + value)*100)/100
						fmt.Println(bonds[i].Name + "|", value,  "|" + "  " ,bonds[i].Count, "  |",dict[idx]["date"])
						tmpl = append(tmpl,Info{Name: bonds[i].Name,Date:dict[idx]["date"],Value:value})
						fmt.Println("monthDictOld: ",monthDict)
						// проверяем есть ли такой ключ
						_ ,exist := monthDict[month]
						// если да, то...
						if exist {
							monthDict[month] += value// month:value
						}else{
							monthDict[month] = value// month:value
						}
						fmt.Println("monthDictNew: ",monthDict)
					}
				}
			}else{
				value = value * bonds[i].Count
				yearSum = yearSum + value
				fmt.Println(bonds[i].Name + "|", value,  "|" + "  " ,bonds[i].Count, "  |")
				tmpl = append(tmpl,Info{Name: bonds[i].Name,Date:dict[idx]["date"],Value:value,YearSum: yearSum})
			}
		}
	}
	for key,sum := range monthDict{
		fmt.Println(year+"-"+key, ": ", sum, "rub")
	}
	fmt.Println("amount per year : ", yearSum, "rub")

	fmt.Println(monthDict)
}

func getAllBonds(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tmpl)
}

func main(){
	var year string
	fmt.Println("enter year or print 'any' if you want to see the bonds for the entire period:  ")
	fmt.Scan(&year)

	TakeData(year)

	r := mux.NewRouter()
	r.HandleFunc("/", getAllBonds).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))

}