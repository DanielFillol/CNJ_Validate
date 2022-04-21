package Functions

import (
	"errors"
	"strconv"
)

func OriginCourt(cnj string) (string, string, error) {
	nCNJ, err := ReturnStructCNJ(cnj)
	if err != nil {
		return "", "", err
	}

	segment1, segment2, err := Segment(cnj)
	if err != nil {
		return "", "", err
	}

	oc, err := strconv.Atoi(nCNJ.Court)
	if err != nil {
		return "", "", err
	}

	sg, err := strconv.Atoi(nCNJ.Segment)
	if err != nil {
		return "", "", err
	}

	var text1 string
	var text2 string

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
			return "", "", errors.New("invalid cnj number, this court , does not exist")
		}
	} else if oc == 90 {
		if sg == 4 {
			text1 = "processo originário"
			text2 = "Conselho da Justiça Federal"
		} else if sg == 5 {
			text1 = "processo originário"
			text2 = "Conselho Superior da Justiça do Trabalho"
		} else {
			return "", "", errors.New("invalid cnj number, this court does not exist")
		}
	} else {
		if sg == 4 {
			err = validOC(oc, 24)
			if err != nil {
				return "", "", err
			}
			text1 = "região"
			text2 = strconv.Itoa(oc)
		} else if sg == 5 {
			err = validOC(oc, 24)
			if err != nil {
				return "", "", err
			}
			text1 = "região"
			text2 = strconv.Itoa(oc)
		} else if sg == 6 || sg == 8 {
			err = validOC(oc, 27)
			if err != nil {
				return "", "", err
			}
			text1 = "unidade federativa"
			text2 = strconv.Itoa(oc)
		} else if sg == 7 {
			err = validOC(oc, 12)
			if err != nil {
				return "", "", err
			}
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
				return "", "", errors.New("invalid cnj number, this court does not exist")
			}
		}
	}

	return text1, text2, nil
}

func validOC(elem int, max int) error {
	if elem > max || elem < 1 {
		return errors.New("invalid cnj number, this court does not exist")
	} else {
		return nil
	}
}
