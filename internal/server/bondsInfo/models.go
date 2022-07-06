package bondsInfo

import "modules/internal/server/structs"

type BondInfoProvider interface {
	GetBond(ticker string) structs.Bond
	GetBondsForYear(year string, bonds []structs.Bond) []structs.Bond
}
