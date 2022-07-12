package bondsInfo

type BondInfoProvider interface {
	GetBond(ticker string) Bond
	GetBondsForYear(year string, bonds []Bond) []Bond
}

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
