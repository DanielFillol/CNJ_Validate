# CNJ_Validate
Project created for verification of all information contained in CNJ format

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

	cnj := "1004144-88.2020.8.01.0037"

	//Returns true if the CNJ is valid and the validation digit that it should have
	bCNJ, nvd := Functions.ValidateCNJ(cnj)

	//Returns the valid CNJ in string format
	vCNJ := Functions.ReturnValidCNJ(cnj)

	//Returns a complex struct "AnalysisCNJ" with all the data in CNJ format; returns error if any step of the verification is faulted
	aCNJ, err := CNJ_Validate.AnalyzeCNJ(cnj)
	Error.CheckError(err)

	//Returns an organized string of all cnj attributes
	cnjW := CNJ_Validate.CNJWrite(aCNJ)

	fmt.Println("O CNJ existe?", bCNJ)
	fmt.Println("O dígito verificador correto:", nvd)
	fmt.Println("O CNJ correto seria:", vCNJ)
	fmt.Println("O cnj analisado:", aCNJ)
	fmt.Println(cnjW)
	
}

 ```

