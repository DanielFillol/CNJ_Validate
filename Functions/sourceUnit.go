package Functions

import (
	"errors"
	"strconv"
)

func SourceUnit(cnj string) (string, string, error) {
	nCNJ, err := ReturnStructCNJ(cnj)
	if err != nil {
		return "", "", err
	}

	sg, err := strconv.Atoi(nCNJ.Segment)
	if err != nil {
		return "", "", err
	}

	su := nCNJ.SourceUnit

	var text1 string
	var text2 string

	if su == "0000" {
		text1 = "competência originária"
		text2 = "tribunal"
	} else if su[0:1] == "9" {
		text1 = "competência originária"
		text2 = "turmas recursais"
	}

	if sg == 4 {
		text1 = "subseção judiciária"
		text2 = su
	} else if sg == 5 {
		text1 = "vara do trabalho"
		text2 = su
	} else if sg == 6 {
		text1 = "zona eleitoral"
		text2 = su
	} else if sg == 7 || sg == 9 {
		text1 = "auditoria militar"
		text2 = su
	} else if sg == 8 {
		text1 = "foro"
		text2 = su
	} else {
		er := errors.New("invalid cnj number, this segment does not exist")
		return "", "", er

	}

	return text1, text2, nil
}
