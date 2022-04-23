package CNJ

type AnalysisCNJ struct {
	ReceivedCNJ      string        // "CNJ_original",
	ValidCNJ         bool          // "CNJ_é_válido",
	SegmentName      string        // "Segmento",
	SegmentShort     string        // "Segmento_short",
	SourceUnitType   string        // "Tipo_Unidade_Judiciária",
	SourceUnitNumber string        // "Número_Unidade_Judiciária",
	CourtType        string        // "Tipo_Região",
	CourtNumber      string        // "Número_Região",
	Detailed         DecomposedCNJ //
}

type DecomposedCNJ struct {
	LawsuitNumber  string `json:"lawsuit_number,omitempty"`  // "Número_Processo",
	VerifyingDigit string `json:"verifying_digit,omitempty"` // "Dígito_Verificador",
	ProtocolYear   string `json:"protocol_year,omitempty"`   // "Ano Protocolo",
	Segment        string `json:"segment,omitempty"`         // "Poder Judiciário",
	Court          string `json:"court,omitempty"`           // "Região",
	SourceUnit     string `json:"source_unit,omitempty"`     // "Unidade Judiciária",
	ArgNumber      string `json:"arg_number,omitempty"`      //
	District       string `json:"district,omitempty"`        // "Nome_Unidade_Judiciária",
	UF             string `json:"UF,omitempty"`              // "Nome_Região",
}

//AnalyzeCNJ returns the complex struct AnalysisCNJ containing all the useful data from this package
func AnalyzeCNJ(cnj string) (AnalysisCNJ, error) {
	decomposedCNJ, err := DecomposeCNJ(cnj)
	if err != nil {
		return AnalysisCNJ{}, err
	}

	vCNJ, err := ValidateCNJ(cnj)
	if err != nil {
		return AnalysisCNJ{}, err
	}

	segment, err := GetSegment(decomposedCNJ.Segment)
	if err != nil {
		return AnalysisCNJ{}, err
	}

	sourceUnit, err := GetSourceUnit(decomposedCNJ.SourceUnit, segment)
	if err != nil {
		return AnalysisCNJ{}, err
	}

	originCourt, err := GetOriginCourt(decomposedCNJ.Court)
	if err != nil {
		return AnalysisCNJ{}, err
	}

	return AnalysisCNJ{
		ReceivedCNJ:      cnj,
		ValidCNJ:         vCNJ,
		SegmentName:      segment.Name,
		SegmentShort:     segment.Short,
		SourceUnitType:   sourceUnit.SourceUnitType,
		SourceUnitNumber: sourceUnit.SourceUnitNumber,
		CourtType:        originCourt.originCourtType,
		CourtNumber:      originCourt.originCourtNumber,
		Detailed:         decomposedCNJ,
	}, nil

}

func WriteCNJ(number AnalysisCNJ) string {
	lawsuit := number.Detailed.LawsuitNumber
	year := number.Detailed.ProtocolYear
	segment1 := number.SegmentName
	segment2 := number.SegmentShort
	sourceU1 := number.SourceUnitType
	ct1 := number.CourtType

	var sourceU2 string
	if number.Detailed.District != "" {
		sourceU2 = number.Detailed.District
	} else {
		sourceU2 = number.SourceUnitNumber
	}

	var ct2 string
	if number.Detailed.UF != "" {
		ct2 = number.Detailed.UF
	} else {
		ct2 = number.CourtNumber
	}

	var preposition string
	if number.Detailed.SourceUnit != "8" {
		preposition = "o"
	} else {
		preposition = "a"
	}

	return "Processo número: " + lawsuit + ", protocolado n" + preposition + " " + sourceU1 + " de " + sourceU2 + ", no ano " + year + " | " + ct1 + ": " + ct2 + " | " + segment1 + " (" + segment2 + ")"

}
