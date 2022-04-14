package Structs

type CNJNumber struct {
	LawsuitNumber  string `json:"lawsuit_number,omitempty"`
	VerifyingDigit string `json:"verifying_digit,omitempty"`
	ProtocolYear   string `json:"protocol_year,omitempty"`
	Segment        string `json:"segment,omitempty"`
	Court          string `json:"court,omitempty"`
	SourceUnit     string `json:"source_unit,omitempty"`
	ArgNumber      string `json:"arg_number,omitempty"`
	District       string `json:"district,omitempty"`
	UF             string `json:"UF,omitempty"`
}

type AnalysisCNJ struct {
	ReceivedCNJ string    `json:"received_cnj,omitempty"`
	ValidCNJ    bool      `json:"valid_cnj,omitempty"`
	CorrectCNJ  string    `json:"correct_cnj,omitempty"`
	ValidDigit  string    `json:"valid_digit,omitempty"`
	Segment1    string    `json:"segment1,omitempty"`
	Segment2    string    `json:"segment2,omitempty"`
	SourceUnit1 string    `json:"source_unit1,omitempty"`
	SourceUnit2 string    `json:"source_unit2,omitempty"`
	Court1      string    `json:"court1,omitempty"`
	Court2      string    `json:"court2,omitempty"`
	Detailed    CNJNumber `json:"detailed"`
}

type CNJDistrict struct {
	SemiCNJ    string `json:"semiCNJ,omitempty"`
	SourceUnit string `json:"sourceUnit,omitempty"`
	District   string `json:"district,omitempty"`
	UF         string `json:"UF,omitempty"`
}
