package CNJ_Validate

import (
	"errors"
	"fmt"
	"github.com/Darklabel91/CNJ_Validate/CSV"
	"github.com/Darklabel91/CNJ_Validate/Functions"
	"github.com/Darklabel91/CNJ_Validate/Structs"
)

func AnalyzeCNJCSV(rawFilePath string, separator rune, nameResultFolder string) error {
	raw, err := CSV.ReadCsvFile(rawFilePath, separator)
	if err != nil {
		return errors.New("CSV must have only one collum with cnj (without title)")
	}
	err = createCSVs(raw, nameResultFolder)
	if err != nil {
		return err
	}
	fmt.Println("Files created")
	return nil
}

func AnalyzeCNJ(cnj string) (Structs.AnalysisCNJ, error) {
	flag, nvd, err := Functions.ValidateCNJ(cnj)
	if err != nil {
		return Structs.AnalysisCNJ{}, err
	}

	vCNJ, err := Functions.ReturnValidCNJ(cnj)
	if err != nil {
		return Structs.AnalysisCNJ{}, err
	}

	sg1, sg2, err := Functions.Segment(cnj)
	if err != nil {
		return Structs.AnalysisCNJ{}, err
	}

	su1, su2, err := Functions.SourceUnit(cnj)
	if err != nil {
		return Structs.AnalysisCNJ{}, err
	}

	ct1, ct2, err := Functions.OriginCourt(cnj)
	if err != nil {
		return Structs.AnalysisCNJ{}, err
	}

	dt, err := Functions.ReturnStructCNJ(cnj)
	if err != nil {
		return Structs.AnalysisCNJ{}, err
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

func createCSVs(raw []string, nameResultFolder string) error {
	var analyzeCNJCSV []Structs.AnalysisCNJ

	for i := 0; i < len(raw); i++ {
		dataReturn, err := AnalyzeCNJ(raw[i])
		if err != nil {
			return err
		}

		analyzeCNJCSV = append(analyzeCNJCSV, dataReturn)
	}

	err := CSV.ExportCSV("filesOK", nameResultFolder, analyzeCNJCSV)
	if err != nil {
		return err
	}

	return nil
}

func CNJWrite(number Structs.AnalysisCNJ) string {
	lawsuit := number.Detailed.LawsuitNumber
	year := number.Detailed.ProtocolYear
	segment1 := number.Segment1
	segment2 := number.Segment2
	sourceU1 := number.SourceUnit1
	ct1 := number.Court1

	var sourceU2 string
	if number.Detailed.District != "" {
		sourceU2 = number.Detailed.District
	} else {
		sourceU2 = number.SourceUnit2
	}

	var ct2 string
	if number.Detailed.UF != "" {
		ct2 = number.Detailed.UF
	} else {
		ct2 = number.Court2
	}

	var preposition string
	if number.Detailed.SourceUnit != "8" {
		preposition = "o"
	} else {
		preposition = "a"
	}

	return "Processo número: " + lawsuit + ", protocolado n" + preposition + " " + sourceU1 + " de " + sourceU2 + ", no ano " + year + " | " + ct1 + ": " + ct2 + " | " + segment1 + " (" + segment2 + ")"

}
