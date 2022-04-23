package CNJ

import (
	"errors"
	"strconv"
)

const (
	OriginalLawsuit = "Processo Originário"
	MarcialCourt    = "Tribunal Militar Estadual"
	Region          = "região"
	Estate          = "unidade federativa"
	JudicialC       = "circunscrição judiciária"
)

// OriginCourt returns two information:
//
// originCourtType:
//
//	região,
//	unidade federativa,
//	circunscrição judiciária
//	in cases were the lawsuit is not initiate at 1st instance (primeiro grau) it returns:
//	processo originário
//
// originCourtNumber:
// 	the given court number
type OriginCourt struct {
	originCourtType   string
	originCourtNumber string
}

//GetOriginCourt returns OriginCourt
func GetOriginCourt(court string, segment Segment) (OriginCourt, error) {
	var originC OriginCourt
	var err error

	switch court {
	case "00":
		originC, err = parseCourt00(segment)
		if err != nil {
			return OriginCourt{}, err
		}
	case "90":
		originC, err = parseCourt90(segment)
		if err != nil {
			return OriginCourt{}, err
		}
	default:
		originC, err = parseCourtOther(segment, court)
		if err != nil {
			return OriginCourt{}, err
		}
	}

	return originC, nil
}

//Court number 00 means that the given lawsuit is "processo originário"
//, in plain english: this lawsuit does not start at 1st degree
// for that reason entails:
// any state court (2nd degree)
// Supremo Tribunal Federal,
// Conselho Superior de Justiça
// Superior Tribunal de Justiça
// Tribunal Superior do Trabalho
// Tribunal Superior Eleitoral
// Superior Tribunal Militar
func parseCourt00(segment Segment) (OriginCourt, error) {
	var originCourt string

	switch segment.Number {
	case 1, 2, 3:
		originCourt = segment.Name + segment.Short
	case 5:
		originCourt = "Tribunal Superior do Trabalho (TST)"
	case 6:
		originCourt = "Tribunal Superior do Eleitoral (TSE)"
	case 7, 9:
		originCourt = "Superior Tribunal Militar (STM)"
	default:
		return OriginCourt{}, errors.New("invalid cnj number, this court , does not exist")
	}
	return OriginCourt{OriginalLawsuit, originCourt}, nil

}

//Court number 90 means that a given lawsuit is "processo originário"
// in plain english: this lawsuit does not start at 1st degree
// , but it only entails two cases:
// Conselho da Justiça Federal and
// Conselho Superior da Justiça do Trabalho
func parseCourt90(segment Segment) (OriginCourt, error) {
	var originCourt string

	switch segment.Number {
	case 4:
		originCourt = "Conselho da Justiça Federal"
	case 5:
		originCourt = "Conselho Superior da Justiça do Trabalho"
	default:
		return OriginCourt{}, errors.New("invalid cnj number, this court does not exist")
	}

	return OriginCourt{OriginalLawsuit, originCourt}, nil
}

//Martial court returns the 3 possible locations of the origin Court
//
func parseMartialCourt(oc int) (OriginCourt, error) {
	var cType string
	var cNumber string

	switch oc {
	case 13:
		cType = MarcialCourt
		cNumber = "Minas Gerais - MG"
	case 21:
		cType = MarcialCourt
		cNumber = "Rio Grande do Sul - RS"
	case 26:
		cType = MarcialCourt
		cNumber = "São Paulo - SP"
	default:
		return OriginCourt{}, errors.New("invalid cnj number: this court does not exist")
	}

	return OriginCourt{cType, cNumber}, nil
}

//Other court numbers return the standard expect OriginCourt:
// região,
// unidade federativa
// circunscrição judiciária, and there given numbers
func parseCourtOther(segment Segment, orginCourt string) (OriginCourt, error) {
	var cType string
	var cNumber string

	oc, err := strconv.Atoi(orginCourt)
	if err != nil {
		return OriginCourt{}, err
	}

	if isCourtValid(segment, oc) != true {
		return OriginCourt{}, errors.New("invalid cnj number, this court does not exist")
	}

	switch segment.Number {
	case 4:
		cType = Region
		cNumber = strconv.Itoa(segment.Number)
	case 5:
		cType = Region
		cNumber = strconv.Itoa(segment.Number)
	case 6, 8:
		cType = Estate
		cNumber = strconv.Itoa(segment.Number)
	case 7:
		cType = JudicialC
		cNumber = strconv.Itoa(segment.Number)
	case 9:
		martialOC, err := parseMartialCourt(oc)
		if err != nil {
			return OriginCourt{}, err
		}
		cType = martialOC.originCourtType
		cNumber = martialOC.originCourtNumber
	default:
		return OriginCourt{}, errors.New("invalid cnj number: this court does not exist")
	}

	return OriginCourt{cType, cNumber}, nil
}

func isCourtValid(segment Segment, court int) bool {
	switch segment.Number {
	case 4, 5:
		err := ocError(segment, court)
		if err != nil {
			return false
		}
		return true
	case 6, 8:
		err := ocError(segment, court)
		if err != nil {
			return false
		}
		return true
	case 7:
		err := ocError(segment, court)
		if err != nil {
			return false
		}
		return true
	default:
		return false
	}
}

func ocError(segment Segment, court int) error {
	max, err := getMax(segment)
	if err != nil {
		return err
	}

	if court > max || court < 1 {
		return errors.New("invalid cnj number, this court does not exist")
	} else {
		return nil
	}
}

func getMax(segment Segment) (int, error) {
	var max int

	switch segment.Number {
	case 4, 5:
		max = 24
	case 6, 8:
		max = 27
	case 7:
		max = 12
	default:
		return 0, errors.New("can not get max court number: this court does not exist")
	}

	return max, nil
}
