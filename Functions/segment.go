package Functions

import (
	"errors"
	"strconv"
)

func Segment(cnj string) (string, string, error) {
	nCNJ, err := ReturnStructCNJ(cnj)
	if err != nil {
		return "", "", err
	}

	sg, err := strconv.Atoi(nCNJ.Segment)
	if err != nil {
		return "", "", err
	}

	var courtSegment string
	var short string

	if sg == 1 {
		courtSegment = "Supremo Tribunal Federal"
		short = "STF"
	} else if sg == 2 {
		courtSegment = "Conselho Nacional de Justiça"
		short = "CNJ"
	} else if sg == 3 {
		courtSegment = "Superior Tribunal de Justiça"
		short = "STJ"
	} else if sg == 4 {
		courtSegment = "Justiça Federal"
		short = "JF"
	} else if sg == 5 {
		courtSegment = "Justiça do Trabalho"
		short = "JT"
	} else if sg == 6 {
		courtSegment = "Justiça Eleitoral"
		short = "JE"
	} else if sg == 7 {
		courtSegment = "Justiça Militar da União"
		short = "JMU"
	} else if sg == 8 {
		courtSegment = "Justiça dos Estados e do Distrito Federal e Territórios"
		short = "Justiça Comum"
	} else if sg == 9 {
		courtSegment = "Justiça Militar Estadual"
		short = "JME"
	} else {
		er := errors.New("invalid cnj number, this segment does not exist")
		return "", "", er

	}

	return courtSegment, short, nil
}
