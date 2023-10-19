package CNJ

import "errors"

// AnalysisCNJ returns:
//
//	receivedCNJ
//	validCNJ given by ValidateCNJ
//	segmentName and segmentShort given by getSegment
//	sourceUnitType and SourceUnitNumber given by getSourceUnit
//	courtType and courtNumber given by getOriginCourt
//	detailed is a struct on it's on given by DecomposeCNJ
type AnalysisCNJ struct {
	ReceivedCNJ      string        `json:"received_cnj,omitempty"`
	ValidCNJ         bool          `json:"valid_cnj,omitempty"`
	SegmentName      string        `json:"segmentName,omitempty"`
	SegmentShort     string        `json:"segment_short,omitempty"`
	SourceUnitType   string        `json:"source_unit_type,omitempty"`
	SourceUnitNumber string        `json:"source_unit_number,omitempty"`
	CourtType        string        `json:"court_type,omitempty"`
	CourtNumber      string        `json:"court_number,omitempty"`
	Detailed         DecomposedCNJ `json:"Detailed"`
}

// DecomposedCNJ returns the CNJ in decomposed manner:
//
//	lawsuitCNJFormat [NNNNNNN]-[DD].[AAAA].[J].[CT].[0000]
//	 lawsuitNumber: [NNNNNNN]
//	 verifyingDigit: [DD]
//	 protocolYear: [AAAA]
//	 segment: [J]
//	 court: [CT]
//	 sourceUnit: [0000]
//	 argNumber: [NNNNNNNAAAAJCT0000] + "00"
//	 district: district where the lawsuit was proposed, frequently a city name
//	 uf: the uf of the correspondent district
//	 tj: the court shortname [TJ/STJ/TRF..]
type DecomposedCNJ struct {
	LawsuitCNJFormat string `json:"lawsuitCNJFormat,omitempty"`
	LawsuitNumber    string `json:"lawsuit_number,omitempty"`
	VerifyingDigit   string `json:"verifying_digit,omitempty"`
	ProtocolYear     string `json:"protocol_year,omitempty"`
	Segment          string `json:"segment,omitempty"`
	Court            string `json:"court,omitempty"`
	SourceUnit       string `json:"source_unit,omitempty"`
	ArgNumber        string `json:"arg_number,omitempty"`
	District         string `json:"district,omitempty"`
	UF               string `json:"UF,omitempty"`
	TJ               string `json:"TJ,omitempty"`
}

// AnalyzeCNJ returns the complex struct AnalysisCNJ containing all the useful data from this package
func AnalyzeCNJ(cnj string) (AnalysisCNJ, error) {
	decomposedCNJ, err := DecomposeCNJ(cnj)
	if err != nil {
		return AnalysisCNJ{}, err
	}

	vCNJ, err := ValidateCNJ(cnj)
	if err != nil {
		return AnalysisCNJ{}, err
	}

	segment, err := getSegment(decomposedCNJ.Segment)
	if err != nil {
		return AnalysisCNJ{}, err
	}

	sourceUnit, err := getSourceUnit(decomposedCNJ.SourceUnit, segment)
	if err != nil {
		return AnalysisCNJ{}, err
	}

	originCourt, err := getOriginCourt(decomposedCNJ.Court, segment)
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

// WriteCNJ returns a single string with all organized interpretation of CNJ information
func WriteCNJ(number AnalysisCNJ) (string, error) {

	if number.Detailed.LawsuitNumber == "" {
		return "", errors.New("AnalysisCNJ have no usable data")
	}

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

	return "Processo número: " + lawsuit + ", protocolado n" + preposition + " " + sourceU1 + " de " + sourceU2 + ", no ano " + year + " | " + ct1 + ": " + ct2 + " | " + segment1 + " (" + segment2 + ")", nil

}
