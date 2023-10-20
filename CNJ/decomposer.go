package CNJ

import (
	"errors"
	"github.com/DanielFillol/CNJ_Validate/CNJDatabase"
	"strconv"
	"strings"
)

const MATH = "00"

// DecomposeCNJ decompose cnj format number into the specifics:
//
//		lawsuitCNJFormat = [NNNNNNN]-[DD].[AAAA].[J].[CT].[0000]
//	 	LawsuitNumber = [NNNNNNN]
//		VerifyingDigit = [DD]
//		ProtocolYear = [AAAA]
//		Segment = [J]
//		Court = [TR]
//		SourceUnit [OOOO]
//		ArgNumber = [LawsuitNumber + ProtocolYear + Segment + Court + SourceUnit + "00"]
//		DistrictInfo = [Name and UF of the city] it is a reading of J.TR.OOOO combination
//		UF = [UF]
//		TJ = [{COURT}+{UF}]
func DecomposeCNJ(cnj string) (DecomposedCNJ, error) {
	if len(cnj) > 25 && len(cnj) < 20 {
		return DecomposedCNJ{}, errors.New("CNJ out of format, it must have 25 or 20 char")
	}

	splitAtVerifyingDigit := strings.Split(cnj, "-")

	if len(splitAtVerifyingDigit) > 1 {
		lawsuitNumber := splitAtVerifyingDigit[0]

		remainCNJ := splitAtVerifyingDigit[1]
		splitRemainingCNJ := strings.Split(remainCNJ, ".")

		verifyingDigit := splitRemainingCNJ[0]
		yearProtocol := splitRemainingCNJ[1]
		segment := splitRemainingCNJ[2]
		court := splitRemainingCNJ[3]
		sourceUnit := splitRemainingCNJ[4]
		argNumber := lawsuitNumber + yearProtocol + segment + court + sourceUnit + MATH
		semiCNJ := segment + "." + court + "." + sourceUnit

		dt, err := CNJDatabase.FetchDistrict(semiCNJ)
		sg, err := getSegment(segment)
		if err != nil {
			return DecomposedCNJ{
				LawsuitCNJFormat: lawsuitNumber + "-" + verifyingDigit + "." + yearProtocol + "." + segment + "." + court + "." + sourceUnit,
				LawsuitNumber:    lawsuitNumber,
				VerifyingDigit:   verifyingDigit,
				ProtocolYear:     yearProtocol,
				Segment:          segment,
				Court:            court,
				SourceUnit:       sourceUnit,
				ArgNumber:        argNumber,
				District:         err.Error(),
				UF:               err.Error(),
				TJ:               err.Error(),
			}, err
		}

		var tj string
		if sg.Number == 1 || sg.Number == 2 || sg.Number == 3 || sg.Number == 4 || sg.Number == 7 {
			courtNumber, err := strconv.Atoi(court)
			if err != nil {
				return DecomposedCNJ{
					LawsuitCNJFormat: lawsuitNumber + "-" + verifyingDigit + "." + yearProtocol + "." + segment + "." + court + "." + sourceUnit,
					LawsuitNumber:    lawsuitNumber,
					VerifyingDigit:   verifyingDigit,
					ProtocolYear:     yearProtocol,
					Segment:          segment,
					Court:            court,
					SourceUnit:       sourceUnit,
					ArgNumber:        argNumber,
					District:         err.Error(),
					UF:               err.Error(),
					TJ:               err.Error(),
				}, err
			}
			tj = sg.Short + strconv.Itoa(courtNumber)
		} else {
			tj = sg.Short + dt.UF
		}

		return DecomposedCNJ{
			LawsuitCNJFormat: lawsuitNumber + "-" + verifyingDigit + "." + yearProtocol + "." + segment + "." + court + "." + sourceUnit,
			LawsuitNumber:    lawsuitNumber,
			VerifyingDigit:   verifyingDigit,
			ProtocolYear:     yearProtocol,
			Segment:          segment,
			Court:            court,
			SourceUnit:       sourceUnit,
			ArgNumber:        argNumber,
			District:         dt.SourceUnit,
			UF:               dt.UF,
			TJ:               tj,
		}, nil
	} else {
		lawsuitNumber := cnj[0:7]

		last13 := cnj[len(cnj)-13:]
		verifyingDigit := last13[0:2]

		last11 := cnj[len(cnj)-11:]
		yearProtocol := last11[0:4]

		last7 := cnj[len(cnj)-7:]
		segment := last7[0:1]

		last6 := cnj[len(cnj)-6:]
		court := last6[0:2]

		sourceUnit := cnj[len(cnj)-4:]
		argNumber := lawsuitNumber + yearProtocol + segment + court + sourceUnit + MATH
		semiCNJ := segment + "." + court + "." + sourceUnit

		dt, err := CNJDatabase.FetchDistrict(semiCNJ)
		sg, err := getSegment(segment)
		if err != nil {
			if err != nil {
				return DecomposedCNJ{
					LawsuitCNJFormat: lawsuitNumber + "-" + verifyingDigit + "." + yearProtocol + "." + segment + "." + court + "." + sourceUnit,
					LawsuitNumber:    lawsuitNumber,
					VerifyingDigit:   verifyingDigit,
					ProtocolYear:     yearProtocol,
					Segment:          segment,
					Court:            court,
					SourceUnit:       sourceUnit,
					ArgNumber:        argNumber,
					District:         err.Error(),
					UF:               err.Error(),
					TJ:               err.Error(),
				}, err
			}
		}

		var tj string
		if sg.Number == 1 || sg.Number == 2 || sg.Number == 3 || sg.Number == 4 || sg.Number == 7 {
			courtNumber, err := strconv.Atoi(court)
			if err != nil {
				return DecomposedCNJ{
					LawsuitCNJFormat: lawsuitNumber + "-" + verifyingDigit + "." + yearProtocol + "." + segment + "." + court + "." + sourceUnit,
					LawsuitNumber:    lawsuitNumber,
					VerifyingDigit:   verifyingDigit,
					ProtocolYear:     yearProtocol,
					Segment:          segment,
					Court:            court,
					SourceUnit:       sourceUnit,
					ArgNumber:        argNumber,
					District:         err.Error(),
					UF:               err.Error(),
					TJ:               err.Error(),
				}, err
			}
			tj = sg.Short + strconv.Itoa(courtNumber)
		} else {
			tj = sg.Short + dt.UF
		}

		return DecomposedCNJ{
			LawsuitCNJFormat: lawsuitNumber + "-" + verifyingDigit + "." + yearProtocol + "." + segment + "." + court + "." + sourceUnit,
			LawsuitNumber:    lawsuitNumber,
			VerifyingDigit:   verifyingDigit,
			ProtocolYear:     yearProtocol,
			Segment:          segment,
			Court:            court,
			SourceUnit:       sourceUnit,
			ArgNumber:        argNumber,
			District:         dt.District,
			UF:               dt.UF,
			TJ:               tj,
		}, nil

	}
}
