package Functions

import "github.com/Darklabel91/CNJ_Validate/Error"

func ReturnValidCNJ(cnj string) string {
	var rCNJ string

	flag, nvd := ValidateCNJ(cnj)

	if flag == false {
		nCNJ, err := ReturnStructCNJ(cnj)
		Error.CheckError(err)

		rCNJ = nCNJ.LawsuitNumber + "-" + nvd + "." + nCNJ.ProtocolYear + "." + nCNJ.Segment + "." + nCNJ.Court + "." + nCNJ.SourceUnit
	} else {
		rCNJ = cnj
	}

	return rCNJ
}
