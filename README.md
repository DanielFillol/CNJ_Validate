# CNJ_Validate
Projeto criado para realizar a manipulação de todos os dados possíveis dentro do padrão do CNJ. Todo o projeto tem como base na [Resolução 65 do CNJ](https://atos.cnj.jus.br/files/resolucao_comp_65_16122008_26032019140041.pdf)

## Instal
``` go get -u github.com/Darklabel91/CNJ_Validate ```

## Data Struct
Os dados de retorno podem ser ```string```, ```bool```, ```AnalysisCNJ``` ou ```CNJNumber``` , essas últimas são composta por:

``` 
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
```
### AnalysisCNJ
- ReceivedCNJ: o cnj passado como parâmetro [0001327-64.2018.8.26.0158]
- ValidCNJ: informa '''true''' caso o *ReceivedCNJ* tenha o dígito verificador como válido [true]
- CorrectCNJ: informa o CNJ com o dígito verificador válido [0001327-64.2018.8.26.0158]
- ValidDigit: informa o dígito verificador correto para o *ReceivedCNJ* [64]
- Segment1: informa o seguimento do ramo da justiça correspondente [Justiça dos Estados e do Distrito Federal e Territórios]
- Segment2: informa a sigla do seguimento da justiça correspondente [Justiça Comum]
- SourceUnit1: informa o tipo da unidade de origem [foro]
- SourceUnit2: informa a unidade de origem [0158]
- Court1: informa o tipo de corte de origem [unidade federativa]
- Court2: informa a unidade da corte de origem [26]
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
	"github.com/Darklabel91/CNJ_Validate"
	"github.com/Darklabel91/CNJ_Validate/Error"
	"github.com/Darklabel91/CNJ_Validate/Functions"
)

func main() {

	cnj := "0001327-64.2018.8.26.0158"

	//Returns true if the CNJ is valid and the validation digit that it should have
	bCNJ, _ := Functions.ValidateCNJ(cnj)
	fmt.Println(bCNJ)

	//Returns the valid CNJ in string format
	vCNJ := Functions.ReturnValidCNJ(cnj)
	fmt.Println(vCNJ)

	//Returns a complex struct "AnalysisCNJ" with all the data in CNJ format ; returns error if any step of the verification is faulted
	aCNJ, err := CNJ_Validate.AnalyzeCNJ(cnj)
	Error.CheckError(err)
	fmt.Println(aCNJ)

	//Returns an organized string of all cnj attributes
	cnjW := CNJ_Validate.CNJWrite(aCNJ)
	fmt.Println(cnjW)

	//Returns a CSV File with the structured cnj analysis
	raw := "/Users/Desktop/test.csv"
	sp := ','
	resultF := "TestFolder"

	CNJ_Validate.AnalyzeCNJCSV(raw, sp, resultF)

}

 ```
 Output
 ``` 
true

0001327-64.2018.8.26.0158

{0001327-64.2018.8.26.0158 true 0001327-64.2018.8.26.0158 64 Justiça dos Estados e do Distrito Federal e Territórios Justiça Comum foro 0158 unidade federativa 26 {0001327 64 2018 8 26 0158 00013272018826015800 São Paulo São Paulo}}

Processo número: 0001327, protocolado no foro de São Paulo, no ano 2018 | unidade federativa: São Paulo | Justiça dos Estados e do Distrito Federal e Territórios (Justiça Comum)

Files created
 ```
 
 ## Functions

### Main Function:
- AnalyzeCNJ(cnj string) retorna a estrutura *AnalysisCNJ* necessitando apenas de um CNJ no formato *NNNNNNN-DD.AAAA.J.CT.0000* ou *NNNNNNNDDAAAAJCT0000* retorna erro caso qualquer verificação seja inválida.
- AnalyzeCNJCSV(rawFilePath string, separator rune, nameResultFolder string) retorna um CSV com a estrutura *AnalysisCNJ* necessitando do caminho onde está o arquivo para leitura (devendo ter apenas uma coluna com os números CNJ), o separador (','), e o nome da pasta em que os arquivos devem retornar
- CNJWrite(number Structs.AnalysisCNJ) retorna uma frase para demostrar a organização das informações, necessita de um *AnalysisCNJ*

### Suport Functions:
- ReturnStructCNJ(cnj string): retorna *Structs.CNJNumber*, retorna erro caso o cnj esteja fora do padrão
- OriginCourt(cnj string): retorna o *Court1* e *Court2*, retorna erro caso a sequência seja inválida
- Segment(cnj string): retorna o *Segment1* e *Segment2*, retorna erro caso o a sequência seja inválida
- SourceUnit(cnj string): retorna o *SourceUnit1* e *SourceUnit2*, retorna erro caso a sequência seja inválida
- ValidateCNJ(cnj string): retorna *true* caso o cnj sejá válido e o dígito verificador
- ReturnValidCNJ(cnj string): retorna o CNJ com o dígito verificador correto
- ValidVD(cnj string): retorna o dígito verificador correto

## Considerações
A) Esse projeto foi criado de forma voluntária, você pode contribuir de qualquer modo. Se encontrar uma falha, não hesite em criar um “issue” ou  procure corrigir você mesma(o) o erro e dar um “pull request”.

B) use os dados baixados para agregar valor, como por exemplo, para realizar análises ou publicar artigos, fazer inferências, elaborar recomendações aos poderes públicos etc. Baixar esses dados para reproduzi-los em sua página web é tirar proveito do trabalho alheio, mesmo sendo esses dados públicos.
