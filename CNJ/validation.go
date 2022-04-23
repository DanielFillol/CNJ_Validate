package CNJ

import (
	"errors"
	"math/big"
)

const (
	MOD = "97"
	SUB = "98"
)

// ValidateCNJ returns if a given cnj number is valid or not using  ValidVD
func ValidateCNJ(cnj string) (bool, error) {
	decomposedCNJ, err := DecomposeCNJ(cnj)
	if err != nil {
		return false, err
	}

	vd, err := validVD(decomposedCNJ.ArgNumber)
	if err != nil {
		return false, err
	}

	if decomposedCNJ.VerifyingDigit != vd {
		return false, nil
	} else {
		return true, nil
	}
}

// validVD returns the verifyng digit from a given cnj using ArgNumber
func validVD(argNumber string) (string, error) {
	cnjInt, bl := new(big.Int).SetString(argNumber, 10)
	if bl != true {
		return "", errors.New("cant convert ArgNumber into big int")
	}

	baseResInt, bl := new(big.Int).SetString(MOD, 10)
	if bl != true {
		return "", errors.New("cant convert 97 into big int")
	}

	baseSubInt, bl := new(big.Int).SetString(SUB, 10)
	if bl != true {
		return "", errors.New("cant convert 98 into big int")
	}

	div := new(big.Int).Mod(cnjInt, baseResInt)
	vd := new(big.Int).Sub(baseSubInt, div)

	var nvd string
	if len(vd.String()) < 2 {
		nvd = "0" + vd.String()
	} else {
		nvd = vd.String()
	}

	return nvd, nil
}
