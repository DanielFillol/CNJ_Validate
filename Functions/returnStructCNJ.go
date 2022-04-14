package Functions

import (
	"errors"
	"github.com/Darklabel91/CNJ_Validate/Database"
	"github.com/Darklabel91/CNJ_Validate/Structs"
)

func ReturnStructCNJ(cnj string) (Structs.CNJNumber, error) {
	var err error

	math := "00"
	err = nil

	lawsuitNumber := ""
	verifyingDigit := ""
	protocolYear := ""
	segment := ""
	court := ""
	sourceUnit := ""

	if len(cnj) == 25 {
		last17 := cnj[len(cnj)-17:]
		last9 := cnj[len(cnj)-9:]
		last7 := cnj[len(cnj)-7:]
		last14 := cnj[len(cnj)-14:]

		lawsuitNumber = cnj[0:7]
		verifyingDigit = last17[0:2]
		protocolYear = last14[0:4]
		segment = last9[0:1]
		court = last7[0:2]
		sourceUnit = cnj[len(cnj)-4:]

	} else if len(cnj) == 20 {
		last13 := cnj[len(cnj)-13:]
		last7 := cnj[len(cnj)-7:]
		last6 := cnj[len(cnj)-6:]
		last11 := cnj[len(cnj)-11:]

		lawsuitNumber = cnj[0:7]
		verifyingDigit = last13[0:2]
		protocolYear = last11[0:4]
		segment = last7[0:1]
		court = last6[0:2]
		sourceUnit = cnj[len(cnj)-4:]
	} else {
		err = errors.New("CNJ out of format, it must have 25 or 20 char")
	}

	args := lawsuitNumber + protocolYear + segment + court + sourceUnit + math
	SemiCNJ := segment + "." + court + "." + sourceUnit

	dt, uf := matchDatabase(SemiCNJ)

	return Structs.CNJNumber{
		LawsuitNumber:  lawsuitNumber,
		VerifyingDigit: verifyingDigit,
		ProtocolYear:   protocolYear,
		Segment:        segment,
		Court:          court,
		SourceUnit:     sourceUnit,
		ArgNumber:      args,
		District:       dt,
		UF:             uf,
	}, err
}

func matchDatabase(semiCNJ string) (string, string) {
	db := Database.ReturnDatabase()
	district := ""
	uf := ""

	for i := 0; i < len(db); i++ {
		if semiCNJ == db[i].SemiCNJ {
			district = db[i].SourceUnit
			uf = db[i].UF
		}
	}

	return district, uf
}
