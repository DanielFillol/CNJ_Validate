package Functions

import (
	"math/big"
)

func ValidVD(cnj string) (string, error) {
	nCNJ, err := ReturnStructCNJ(cnj)
	if err != nil {
		return "", err
	}

	cnjInt, _ := new(big.Int).SetString(nCNJ.ArgNumber, 10)
	baseResInt, _ := new(big.Int).SetString("97", 10)
	baseSubInt, _ := new(big.Int).SetString("98", 10)

	div := new(big.Int).Mod(cnjInt, baseResInt)
	nvd := new(big.Int).Sub(baseSubInt, div)

	return nvd.String(), nil
}
