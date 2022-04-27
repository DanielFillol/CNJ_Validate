# CNJ_Validate
Projeto criado para realizar a manipulação de todos os dados possíveis dentro do padrão do CNJ. Todo o projeto tem como base na [Resolução 65 do CNJ](https://atos.cnj.jus.br/files/resolucao_comp_65_16122008_26032019140041.pdf)

## Instal
``` go get -u github.com/Darklabel91/CNJ_Validate ```

## Data Struct
Os dados de retorno podem ser ```string```, ```bool```, ```AnalysisCNJ``` ou ```DecomposedCNJ``` , essas últimas são composta por:

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
```
### AnalysisCNJ
- ReceivedCNJ: o cnj passado como parâmetro [0001327-64.2018.8.26.0158]
- ValidCNJ: informa '''true''' caso o *ReceivedCNJ* tenha o dígito verificador como válido [true]
- SegmentName: informa o seguimento do ramo da justiça correspondente [Justiça dos Estados e do Distrito Federal e Territórios]
- SegmentShort: informa a sigla do seguimento da justiça correspondente [Justiça Comum]
- SourceUnitType: informa o tipo da unidade de origem [foro]
- SourceUnitNumber: informa a unidade de origem [0158]
- CourtType: informa o tipo de corte de origem [unidade federativa]
- CourtNumber: informa a unidade da corte de origem [26]
- Detailed: retorna a estrutura do CNJNumber

### CNJNumber
- LawsuitNumber: NNNNNNN [0001327]
- VerifyingDigit: DD [64]
- ProtocolYear: AAAA [2018]
- Segment: J [8]
- Court: CT [26]
- SourceUnit: 0000 [0158]
- ArgNumber:NNNNNNNAAAAJCT0000 + 00 [00013272018826015800] (utilizado pelo cálculo do *VerifyingDigit*)
- District: o foro em que a ação foi protocolada (normalmente a cidade, zona eleitoral, vara trabalhista) [São Paulo]
- UF: A unidade federativa que o foro em que a ação foi protocolada pertence [São Paulo] 


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
- AnalyzeCNJ(cnj string) retorna a estrutura *AnalysisCNJ* necessitando apenas de um CNJ no formato *NNNNNNN-DD.AAAA.J.CT.0000* ou *NNNNNNNDDAAAAJCT0000* retorna erro caso qualquer verificação seja inválida.
- CNJWrite(number Structs.AnalysisCNJ) retorna uma frase para demostrar a organização das informações, necessita de um *AnalysisCNJ*
- AnalyzeCNJCSV(rawFilePath string, separator rune, nameResultFolder string) retorna um CSV com a estrutura *AnalysisCNJ* necessitando do caminho onde está o arquivo para leitura (devendo ter apenas uma coluna com os números CNJ), o separador (','), e o nome da pasta em que os arquivos devem retornar


### Suport Functions:
-  DecomposeCNJ(cnj string) retorna *[DecomposedCNJ](https://pkg.go.dev/github.com/Darklabel91/CNJ_Validate/CNJ#DecomposedCNJ)* necessitando de um CNJ no formato *NNNNNNN-DD.AAAA.J.CT.0000* ou *NNNNNNNDDAAAAJCT0000*
-  GetOriginCourt(court string, segment Segment) retorna [OriginCourt](https://pkg.go.dev/github.com/Darklabel91/CNJ_Validate/CNJ#OriginCourt) necessitando do códgigo do tribunal (composto por dois dígitos) e um *[Segment](https://pkg.go.dev/github.com/Darklabel91/CNJ_Validate/CNJ#Segment)*
- GetSourceUnit(sourceUnit string, segment Segment) retorna *[SourceUnit](https://pkg.go.dev/github.com/Darklabel91/CNJ_Validate/CNJ#SourceUnit)* necessitando do código da unidade (composto por 4 dígitos) e um *[Segment](https://pkg.go.dev/github.com/Darklabel91/CNJ_Validate/CNJ#Segment)*
- ValidateCNJ(cnj string) retorna um *True* para uma sequência CNJ válida no formato *NNNNNNN-DD.AAAA.J.CT.0000* ou *NNNNNNNDDAAAAJCT0000* 

## Considerações
A) Esse projeto foi criado de forma voluntária, você pode contribuir de qualquer modo. Se encontrar uma falha, não hesite em criar um “issue” ou  procure corrigir você mesma(o) o erro e dar um “pull request”.

B) use os dados baixados para agregar valor, como por exemplo, para realizar análises ou publicar artigos, fazer inferências, elaborar recomendações aos poderes públicos etc. Baixar esses dados para reproduzi-los em sua página web é tirar proveito do trabalho alheio, mesmo sendo esses dados públicos.
