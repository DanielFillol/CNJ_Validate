package Functions

import "github.com/Darklabel91/CNJ_Validate/Error"

func ValidateCNJ(cnj string) (bool, string) {
	var flag bool

	nCNJ, err0 := ReturnStructCNJ(cnj)
	Error.CheckError(err0)

	nvd := ValidVD(cnj)

	if nCNJ.VerifyingDigit != nvd {
		flag = false
	} else {
		flag = true
	}

	return flag, nvd
}
