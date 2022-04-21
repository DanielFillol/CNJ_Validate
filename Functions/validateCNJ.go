package Functions

func ValidateCNJ(cnj string) (bool, string, error) {
	var flag bool

	nCNJ, err := ReturnStructCNJ(cnj)
	if err != nil {
		return false, "", err
	}

	nvd, err := ValidVD(cnj)
	if err != nil {
		return false, "", err
	}

	if nCNJ.VerifyingDigit != nvd {
		flag = false
	} else {
		flag = true
	}

	return flag, nvd, nil
}
