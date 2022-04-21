package Functions

func ReturnValidCNJ(cnj string) (string, error) {
	flag, nvd, err := ValidateCNJ(cnj)
	if err != nil {
		return "", err
	}

	if flag == false {
		nCNJ, err := ReturnStructCNJ(cnj)
		if err != nil {
			return "", err
		}

		return nCNJ.LawsuitNumber + "-" + nvd + "." + nCNJ.ProtocolYear + "." + nCNJ.Segment + "." + nCNJ.Court + "." + nCNJ.SourceUnit, nil
	} else {
		return cnj, nil
	}
}
