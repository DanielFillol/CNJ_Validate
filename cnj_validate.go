package CNJ_Validate

import (
	"errors"
	"github.com/Darklabel91/CNJ_Validate/Error"
	"github.com/Darklabel91/CNJ_Validate/Functions"
	"github.com/Darklabel91/CNJ_Validate/Structs"
	"strconv"
)

func AnalyzeCNJ(cnj string) (Structs.AnalysisCNJ, error) {
	var err error

	flag, nvd := Functions.ValidateCNJ(cnj)
	vCNJ := Functions.ReturnValidCNJ(cnj)

	sg1, sg2, err0 := Functions.Segment(cnj)
	su1, su2, err1 := Functions.SourceUnit(cnj)
	ct1, ct2, err2 := Functions.OriginCourt(cnj)
	dt, err3 := Functions.ReturnStructCNJ(cnj)

	if err0 != nil || err1 != nil || err2 != nil || err3 != nil {
		if err0 != nil {
			err = errors.New("cnj is invalid with the given segment")
		}
		if err1 != nil {
			err = errors.New("cnj is invalid with the given source unit")
		}
		if err2 != nil {
			err = errors.New("cnj is invalid with the given origin court")
		}
		if err3 != nil {
			err = errors.New("CNJ out of format, it must have 25 or 20 char")
		}
	}

	return Structs.AnalysisCNJ{
		ReceivedCNJ: cnj,
		ValidCNJ:    flag,
		CorrectCNJ:  vCNJ,
		ValidDigit:  nvd,
		Segment1:    sg1,
		Segment2:    sg2,
		SourceUnit1: su1,
		SourceUnit2: su2,
		Court1:      ct1,
		Court2:      ct2,
		Detailed:    dt,
	}, err

}

func CNJWrite(number Structs.AnalysisCNJ) string {
	var preposition string
	var text string

	lawsuit := number.Detailed.LawsuitNumber
	year := number.Detailed.ProtocolYear
	segment1 := number.Segment1
	segment2 := number.Segment2
	sourceU1 := number.SourceUnit1
	sourceU2 := number.Detailed.District
	ct1 := number.Court1
	ct2 := number.Detailed.UF

	sg, err1 := strconv.Atoi(number.Detailed.Segment)
	Error.CheckError(err1)

	if sg < 3 {
		preposition = "do"
	} else {
		preposition = "da"
	}

	text = "Processo número: " + lawsuit + ", ajuizado no ano de " + year + ", pertencente ao segmento " + preposition + " " + segment1 + " (" + segment2 + "), tendo como unidade de origem: " + sourceU1 + " de " + sourceU2 + " | " + ct1 + ": " + ct2
	return text
}
