package Functions

import (
	"errors"
	"github.com/Darklabel91/CNJ_Validate/Error"
	"strconv"
)

func Segment(cnj string) (string, string, error) {
	var courtSegment string
	var short string
	var err error

	nCNJ, err0 := ReturnStructCNJ(cnj)
	Error.CheckError(err0)

	sg, err1 := strconv.Atoi(nCNJ.Segment)
	Error.CheckError(err1)

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
		err = errors.New("invalid cnj number, this segment does not exist")
	}

	return courtSegment, short, err
}
