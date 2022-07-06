package bondsInfoMethods

import (
	"net/http"
)

type BondsInfoMethodsService struct {
	parser BondsInfoMethods
}

type BondsInfoMethods interface {
	ResponseParse(resp *http.Response) ([]map[string]string, error)
	Check(year string, data string, month string) bool
}

func (service *BondsInfoMethodsService) Check(year string, date string, month string) bool {
	if year != "any" && date == year && month != "" {
		return true
	}
	return false
}
