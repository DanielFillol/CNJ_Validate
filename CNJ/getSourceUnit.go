package CNJ

import (
	"errors"
)

const (
	JusticeSection  = "subseção judiciária"
	LaborUnit       = "vara do trabalho"
	ElectoralUnit   = "zona eleitoral"
	MilitaryUnit    = "auditoria militar"
	CivilUnit       = "foro"
	CourtUnitSingle = "competência originária do Tribunal"
	CourtUnit       = "competência originária da Turma Recursal"
)

// SourceUnit returns two basic infos:
//	sourceUnitType:
//	-subseção judiciária
//	-vara do trabalho
//	-zona eleitoral
//	-auditoria militar
//	-foro
//	As an exception it returns competência originária do tribunal / turma recursal
//	sourceUnitNumber:
//	- the given source unit code
type SourceUnit struct {
	SourceUnitType   string
	SourceUnitNumber string
}

//getSourceUnit returns SourceUnit
func getSourceUnit(sourceUnit string, segment Segment) (SourceUnit, error) {

	if isZeroSequence(sourceUnit) {
		return SourceUnit{CourtUnitSingle, sourceUnit}, nil
	}

	if isFirstDigit9(sourceUnit) {
		return SourceUnit{CourtUnit, sourceUnit}, nil
	}

	su, err := parseSourceUnit(segment, sourceUnit)
	if err != nil {
		return SourceUnit{}, err
	}

	return su, nil

}

//The first digit of source unit been 9 means the lawsuit have "competência originária"
//on the court where it was initialized
func isFirstDigit9(sourceUnit string) bool {
	if sourceUnit[0:1] == "9" {
		return true
	}
	return false
}

//When all the source unit sequence is composed with only 0 means "competência originária da Turma Recursal"
//on the court where it was initialized
func isZeroSequence(sourceUnit string) bool {
	if sourceUnit == "0000" {
		return true
	}
	return false
}

//parseSourceUnit returns SourceUnit form a given Segment
func parseSourceUnit(segment Segment, sourceUnit string) (SourceUnit, error) {
	var sourceUnitType string
	var sourceUnitNumber string

	switch segment.Number {
	case 4:
		sourceUnitType = JusticeSection
		sourceUnitNumber = sourceUnit
	case 5:
		sourceUnitType = LaborUnit
		sourceUnitNumber = sourceUnit
	case 6:
		sourceUnitType = ElectoralUnit
		sourceUnitNumber = sourceUnit
	case 7, 9:
		sourceUnitType = MilitaryUnit
		sourceUnitNumber = sourceUnit
	case 8:
		sourceUnitType = CivilUnit
		sourceUnitNumber = sourceUnit
	default:
		return SourceUnit{}, errors.New("invalid cnj number, this segment does not exist")
	}

	return SourceUnit{SourceUnitType: sourceUnitType, SourceUnitNumber: sourceUnitNumber}, nil
}
