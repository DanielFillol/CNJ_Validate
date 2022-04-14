package Functions

import (
	"errors"
	"github.com/Darklabel91/CNJ_Validate/Error"
	"strconv"
)

func SourceUnit(cnj string) (string, string, error) {
	var err error
	var text1 string
	var text2 string

	nCNJ, err0 := ReturnStructCNJ(cnj)
	Error.CheckError(err0)

	sg, err1 := strconv.Atoi(nCNJ.Segment)
	Error.CheckError(err1)

	su := nCNJ.SourceUnit

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
		err = errors.New("invalid cnj number, this segment does not exist")
	}

	return text1, text2, err
}
