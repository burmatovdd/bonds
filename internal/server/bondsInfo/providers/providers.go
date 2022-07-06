package providers

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"modules/internal/server/bondsInfo/bondsInfoMethods"
	"modules/internal/server/structs"
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

type MoexBondInfoProvider struct {
}

func (provider *MoexBondInfoProvider) GetBond(ticker string) structs.Bond {
	var bondInfo = structs.Bond{
		Name:    ticker,
		Coupons: []structs.Coupon{},
	}
	response, err := http.Get("https://iss.moex.com/iss/statistics/engines/stock/markets/bonds/bondization/" + ticker + "?iss.meta=off&iss.only=coupons&coupons.columns=coupondate,value")
	if err != nil {
		log.Fatal(err)
	}

	dict, err := xmlResponseParse(response)

	for idx := range dict {

		fullDate := (dict[idx]["date"])[:7]
		value, _ := strconv.ParseFloat(dict[idx]["value"], 64)
		var coupon = structs.Coupon{
			Date:  fullDate,
			Value: value,
		}
		bondInfo.Coupons = append(bondInfo.Coupons, coupon)
	}
	return bondInfo
}

// GetBondsForYear bonds - массив бондов( тикер, кол-во),
//bondsInfos - массив структуры Bond,
//info - данные которые получаем по конкретному тикеру
func (provider *MoexBondInfoProvider) GetBondsForYear(year string, bonds []structs.Bond) []structs.Bond {
	service := bondsInfoMethods.BondsInfoMethodsService{}
	for i := 0; i < len(bonds); i++ {
		info := provider.GetBond(bonds[i].Name)

		for j := 0; j < len(info.Coupons); j++ {
			date := info.Coupons[j].Date[:4]
			month := info.Coupons[j].Date[5:7]
			if service.Check(year, date, month) {
				bonds[i].Coupons = append(bonds[i].Coupons, info.Coupons[j])
			}
		}

	}
	return bonds
}

func xmlResponseParse(resp *http.Response) ([]map[string]string, error) {
	byteValue, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = resp.Body.Close()
	if nil != err {
		fmt.Println("Error response.Body.Close", err)
		return nil, err
	}

	data := &Document{}

	err = xml.Unmarshal(byteValue, &data)
	if nil != err {
		fmt.Println("Error unmarshalling from XML", err)
		return nil, err
	}

	result, err := json.Marshal(&data.Rows)
	if nil != err {
		fmt.Println("Error marshalling to JSON", err)
		return nil, err
	}

	var dict []map[string]string
	if err = json.Unmarshal(result, &dict); err != nil {
		panic(err)
	}
	return dict, nil
}
