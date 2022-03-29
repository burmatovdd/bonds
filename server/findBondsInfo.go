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

var data []Row

var bonds = []Bond{}

func TakeData(year string) Test {
	var monthDict = make(map[string]float64)
	var bondInfos []BondInfo
	var yearSum float64
	bonds = getAllBonds()

	for i := 0; i < len(bonds); i++ {
		var bondInfo = BondInfo{
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
			fullDate := (dict[idx]["date"])[:7]
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
								Date:  fullDate,
								Value: value,
							}
							bondInfo.Coupons = append(bondInfo.Coupons, coupon)
						} else {
							monthDict[month] = value
							var coupon = Coupon{
								Date:  fullDate,
								Value: value,
							}
							bondInfo.Coupons = append(bondInfo.Coupons, coupon)
						}
					}
				}
			} else {
				value = value * bonds[i].Count
				yearSum = yearSum + value
				//fmt.Println(bonds[i].Name+"|", value, "|"+"  ", bonds[i].Count, "  |")
			}
		}
		bondInfos = append(bondInfos, bondInfo)
	}

	keys := make([]string, 0, len(monthDict))
	for key := range monthDict {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	var monthInfos []Months
	for _, key := range keys {
		var monthInfo = Months{
			key,
			monthDict[key],
		}
		monthInfos = append(monthInfos, monthInfo)
	}

	//var allInfos []AllInfo
	var allInfo = AllInfo{
		bondInfos,
		monthInfos,
	}

	//allInfos = append(allInfos, allInfo)

	tmpl := Test{
		allInfo,
	}

	return tmpl
}
