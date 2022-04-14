package CNJ_Validate

import (
	"errors"
	"fmt"
	"github.com/Darklabel91/CNJ_Validate/CSV"
	"github.com/Darklabel91/CNJ_Validate/Error"
	"github.com/Darklabel91/CNJ_Validate/Functions"
	"github.com/Darklabel91/CNJ_Validate/Structs"
)

func AnalyzeCNJCSV(rawFilePath string, separator rune, nameResultFolder string) {
	raw := CSV.ReadCsvFile(rawFilePath, separator)
	createCSVs(raw, nameResultFolder)
	fmt.Println("Files created")
}

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

func createCSVs(raw []string, nameResultFolder string) {
	var analyzeCNJCSV []Structs.AnalysisCNJ

	for i := 0; i < len(raw); i++ {
		dataReturn, err := AnalyzeCNJ(raw[i])
		Error.CheckError(err)
		analyzeCNJCSV = append(analyzeCNJCSV, dataReturn)
	}

	CSV.ExportCSV("filesOK", nameResultFolder, analyzeCNJCSV)

}

func CNJWrite(number Structs.AnalysisCNJ) string {
	var preposition2 string
	var sourceU2 string
	var ct2 string
	var text string

	lawsuit := number.Detailed.LawsuitNumber
	year := number.Detailed.ProtocolYear
	segment1 := number.Segment1
	segment2 := number.Segment2
	sourceU1 := number.SourceUnit1
	ct1 := number.Court1

	if number.Detailed.District != "" {
		sourceU2 = number.Detailed.District
	} else {
		sourceU2 = number.SourceUnit2
	}

	if number.Detailed.UF != "" {
		ct2 = number.Detailed.UF
	} else {
		ct2 = number.Court2
	}

	if number.Detailed.SourceUnit != "8" {
		preposition2 = "o"
	} else {
		preposition2 = "a"
	}

	text = "Processo número: " + lawsuit + ", protocolado n" + preposition2 + " " + sourceU1 + " de " + sourceU2 + ", no ano " + year + " | " + ct1 + ": " + ct2 + " | " + segment1 + " (" + segment2 + ")"

	return text
}
