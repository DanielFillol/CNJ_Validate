# CNJ_Validate
Project created for verification of all information contained in CNJ format

## Instal
``` go get -u github.com/Darklabel91/CNJ_Validate ```

## Data Struct
Os dados de retorno podem ser ```A```, ```B```, ```C``` ou ```D``` , essa última é composta por:

``` 
type A struct {
	B
}
```

- A
- B
- C
- D


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

Main Function:
- A
- b

Suport Functions:
- c
- d
- e

## Considerações
A) Esse projeto foi criado de forma voluntária, você pode contribuir de qualquer modo. Se encontrar uma falha, não hesite em criar um “issue” ou  procure corrigir você mesma(o) o erro e dar um “pull request”.

B) use os dados baixados para agregar valor, como por exemplo, para realizar análises ou publicar artigos, fazer inferências, elaborar recomendações aos poderes públicos etc. Baixar esses dados para reproduzi-los em sua página web é tirar proveito do trabalho alheio, mesmo sendo esses dados públicos.
