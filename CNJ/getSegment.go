package CNJ

import (
	"errors"
	"strconv"
)

// Segment returns three basic infos:
//	number - the segment code
//  name  - the full name of the segment
//  short - the segment nickname
type Segment struct {
	Number int
	Name   string
	Short  string
}

//courtSegment is a map of every possible segment allowed on CNJ format (goes from 1 to 9)
var courtSegment = map[int]Segment{
	1: {Number: 1, Name: "Supremo Tribunal Federal", Short: "STF"},
	2: {Number: 2, Name: "Conselho Nacional de Justiça", Short: "CNJ"},
	3: {Number: 3, Name: "Superior Tribunal de Justiça", Short: "STJ"},
	4: {Number: 4, Name: "Justiça Federal", Short: "JF"},
	5: {Number: 5, Name: "Justiça do Trabalho", Short: "JT"},
	6: {Number: 6, Name: "Justiça Eleitoral", Short: "JE"},
	7: {Number: 7, Name: "Justiça Militar da União", Short: "JMU"},
	8: {Number: 8, Name: "Justiça dos Estados e do Distrito Federal e Territórios", Short: "Comum"},
	9: {Number: 9, Name: "Justiça Militar Estadual", Short: "JME"},
}

// getSegment returns one of nine possible segments:
// 	Supremo Tribunal Federal
// 	Conselho Nacional de Justiça
// 	Superior Tribunal de Justiça
// 	Justiça Federal
// 	Justiça do Trabalho
// 	Justiça Eleitoral
// 	Justiça Militar da União
// 	Justiça dos Estados e do Distrito Federal e Territórios
// 	Justiça Militar Estadual
func getSegment(segment string) (Segment, error) {
	segmentCode, err := strconv.Atoi(segment)
	if err != nil {
		return Segment{}, err
	}

	segmentName, ok := courtSegment[segmentCode]
	if !ok {
		return Segment{}, errors.New("this segment code is invalid: must be between 1 and 9 and single digit")
	}

	return segmentName, nil
}
