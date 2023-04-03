# CNJ_Validate
Package for manipulation and validation of lawsuit number in CNJ format. The CNJ format is defined in [Resolução 65 do CNJ](https://atos.cnj.jus.br/files/resolucao_comp_65_16122008_26032019140041.pdf)

## Instal
``` go get -u github.com/Darklabel91/CNJ_Validate  ```

## Data Struct
Retuurn data can be: ```string```, ```bool```, ```AnalysisCNJ``` or ```DecomposedCNJ```, the last two:

``` 
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

type DecomposedCNJ struct {
	LawsuitCNJFormat string `json:"lawsuitCNJFormat,omitempty"`
	LawsuitNumber    string `json:"lawsuit_number,omitempty"`
	VerifyingDigit   string `json:"verifying_digit,omitempty"`
	ProtocolYear   	 string `json:"protocol_year,omitempty"`
	Segment        	 string `json:"segment,omitempty"`
	Court          	 string `json:"court,omitempty"`
	SourceUnit     	 string `json:"source_unit,omitempty"`
	ArgNumber      	 string `json:"arg_number,omitempty"`
	District       	 string `json:"district,omitempty"`
	UF             	 string `json:"UF,omitempty"`
}
```
### AnalysisCNJ
- ReceivedCNJ: searched CNJ [0001327-64.2018.8.26.0158]
- ValidCNJ: return '''true''' case *ReceivedCNJ*  has a valid verifying digit [true]
- SegmentName: return competent justice branch [Justiça dos Estados e do Distrito Federal e Territórios]
- SegmentShort: return justice branch short name [Justiça Comum]
- SourceUnitType: return justice unit type [foro]
- SourceUnitNumber: return justice unit number [0158]
- CourtType: origin court type [unidade federativa]
- CourtNumber: origin court number [26]
- Detailed: return CNJNumber

### CNJNumber
- LawsuitCNJFormat [NNNNNNN]-[DD].[AAAA].[J].[CT].[0000]
- LawsuitNumber: NNNNNNN [0001327]
- VerifyingDigit: DD [64]
- ProtocolYear: AAAA [2018]
- Segment: J [8]
- Court: CT [26]
- SourceUnit: 0000 [0158]
- ArgNumber:NNNNNNNAAAAJCT0000 + 00 [00013272018826015800] (need for caculate *VerifyingDigit*)
- District: location where the lawsuit began (city, electoral district, labor district)[São Paulo]
- UF: The state short name where the lawsuit began [SP]


## Example

``` 
package main

import (
	"fmt"
	"github.com/Darklabel91/CNJ_Validate/CNJ"
	"github.com/Darklabel91/CNJ_Validate/CNJCSV"
)

func main() {

	cnj := "0001327-64.2018.8.26.0158"

	//Returns true if the CNJ is valid and the validation digit that it should have
	isValid, err := CNJ.ValidateCNJ(cnj)
	fmt.Println(isValid, err)

	//Returns a decomposed CNJ
	decomposedCNJ, err := CNJ.DecomposeCNJ(cnj)
	fmt.Println(decomposedCNJ, err)

	//Returns a complex struct "AnalysisCNJ" with all the data in CNJ format ; returns error if any step of the verification is faulted
	aCNJ, err := CNJ.AnalyzeCNJ(cnj)
	fmt.Println(aCNJ, err)

	//Returns an organized string of all cnj attributes
	cnjW := CNJ.WriteCNJ(aCNJ)
	fmt.Println(cnjW)

	//Returns a CSV File with the structured cnj analysis
	raw := "CSV/exampleCNJfile.csv"
	sp := ','
	resultF := "TestFolder"

	err = CNJCSV.AnalyzeCNJCSV(raw, sp, resultF)
	if err != nil {
		fmt.Println(err)
	}

}

 ```
 Output
 ``` 
 true <nil>
 
{0001327 64 2018 8 26 0158 00013272018826015800 São Paulo SP} <nil>

{0001327-64.2018.8.26.0158 true Justiça dos Estados e do Distrito Federal e Territórios Comum foro 0158 unidade federativa 8 {0001327 64 2018 8 26 0158 00013272018826015800 São Paulo SP}} <nil>

Processo número: 0001327, protocolado no foro de São Paulo, no ano 2018 | unidade federativa: SP | Justiça dos Estados e do Distrito Federal e Territórios (Comum)

Files created
 ```
 
 ## Functions

### Main Function:
- AnalyzeCNJ(cnj string) return *AnalysisCNJ* need CNJ *NNNNNNN-DD.AAAA.J.CT.0000* or *NNNNNNNDDAAAAJCT0000* can return error.
- CNJWrite(number Structs.AnalysisCNJ) return string explaining the lawsuit, need *AnalysisCNJ*
- AnalyzeCNJCSV(rawFilePath string, separator rune, nameResultFolder string) return CSV with *AnalysisCNJ* on given folder in *nameResultFolder*, neeed .csv file path in *rawFilePath*, must be a single column.


### Suport Functions:
-  DecomposeCNJ(cnj string) returns *[DecomposedCNJ](https://pkg.go.dev/github.com/Darklabel91/CNJ_Validate/CNJ#DecomposedCNJ)* need CNJ *NNNNNNN-DD.AAAA.J.CT.0000* or *NNNNNNNDDAAAAJCT0000*
-  GetOriginCourt(court string, segment Segment) return [OriginCourt](https://pkg.go.dev/github.com/Darklabel91/CNJ_Validate/CNJ#OriginCourt) need court code (CT) also return *[Segment](https://pkg.go.dev/github.com/Darklabel91/CNJ_Validate/CNJ#Segment)* and 
- GetSourceUnit(sourceUnit string, segment Segment) return *[SourceUnit](https://pkg.go.dev/github.com/Darklabel91/CNJ_Validate/CNJ#SourceUnit)* need source unit code [0000] also return *[Segment](https://pkg.go.dev/github.com/Darklabel91/CNJ_Validate/CNJ#Segment)*
- ValidateCNJ(cnj string) return *True* for valid CNJ sequence *NNNNNNN-DD.AAAA.J.CT.0000* or *NNNNNNNDDAAAAJCT0000* 
