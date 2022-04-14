package Functions

import (
	"errors"
	"github.com/Darklabel91/CNJ_Validate/Error"
	"strconv"
)

func OriginCourt(cnj string) (string, string, error) {
	var text1 string
	var text2 string
	var err error

	nCNJ, err0 := ReturnStructCNJ(cnj)
	Error.CheckError(err0)

	segment1, segment2, err1 := Segment(cnj)
	Error.CheckError(err1)

	oc, err2 := strconv.Atoi(nCNJ.Court)
	Error.CheckError(err2)

	sg, err3 := strconv.Atoi(nCNJ.Segment)
	Error.CheckError(err3)

	if nCNJ.Court == "00" {
		if sg == 1 || sg == 2 || sg == 3 {
			text1 = "processo originário"
			text2 = segment1 + segment2
		} else if sg == 5 {
			text1 = "processo originário"
			text2 = "Tribunal Superior do Trabalho (TST)"
		} else if sg == 6 {
			text1 = "processo originário"
			text2 = "Tribunal Superior do Eleitoral (TSE)"
		} else if sg == 7 || sg == 9 {
			text1 = "processo originário"
			text2 = "Superior Tribunal Militar (STM)"
		} else {
			err = errors.New("invalid cnj number, this court does not exist")
		}
	} else if oc == 90 {
		if sg == 4 {
			text1 = "processo originário"
			text2 = "Conselho da Justiça Federal"
		} else if sg == 5 {
			text1 = "processo originário"
			text2 = "Conselho Superior da Justiça do Trabalho"
		} else {
			err = errors.New("invalid cnj number, this court does not exist")
		}
	} else {
		if sg == 4 {
			err = validOC(oc, 24)
			text1 = "região"
			text2 = strconv.Itoa(oc)
		} else if sg == 5 {
			err = validOC(oc, 24)
			text1 = "região"
			text2 = strconv.Itoa(oc)
		} else if sg == 6 || sg == 8 {
			err = validOC(oc, 27)
			text1 = "unidade federativa"
			text2 = strconv.Itoa(oc)
		} else if sg == 7 {
			err = validOC(oc, 12)
			text1 = "circunscrição judiciária"
			text2 = strconv.Itoa(oc)
		} else if sg == 9 {
			if oc == 13 {
				text1 = "Minas Gerais - MG"
			} else if oc == 21 {
				text1 = "Rio Grande do Sul - RS"
			} else if oc == 26 {
				text1 = "São Paulo - SP"
			} else {
				err = errors.New("invalid cnj number, this court does not exist")
			}
		}
	}

	return text1, text2, err
}

func validOC(elem int, max int) error {
	if elem > max || elem < 1 {
		return errors.New("invalid cnj number, this court does not exist")
	} else {
		return nil
	}
}
