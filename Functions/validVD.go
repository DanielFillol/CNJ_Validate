package Functions

import (
	"github.com/Darklabel91/CNJ_Validate/Error"
	"math/big"
)

func ValidVD(cnj string) string {
	baseRes := "97"
	baseSub := "98"

	nCNJ, err := ReturnStructCNJ(cnj)
	Error.CheckError(err)

	cnjInt, _ := new(big.Int).SetString(nCNJ.ArgNumber, 10)
	baseResInt, _ := new(big.Int).SetString(baseRes, 10)
	baseSubInt, _ := new(big.Int).SetString(baseSub, 10)

	div := new(big.Int).Mod(cnjInt, baseResInt)
	nvd := new(big.Int).Sub(baseSubInt, div)

	return nvd.String()
}
