package CSV

import (
	"encoding/csv"
	"github.com/Darklabel91/CNJ_Validate/Error"
	"github.com/Darklabel91/CNJ_Validate/Structs"
	"os"
	"path/filepath"
	"strconv"
)

func create(p string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		return nil, err
	}
	return os.Create(p)
}

func ExportCSV(nameFile string, nameFolder string, cnj []Structs.AnalysisCNJ) {
	var empData [][]string

	head := []string{"CNJ_original", "CNJ_e_válido", "CNJ_válido", "Dígito_Verificador_Válido", "Segmento", "Segmento_short", "Tipo_Unidade_Judiciária", "Número_Unidade_Judiciária", "Nome_Unidade_Judiciária", "Tipo_Região", "Número_Região", "Nome_Região", "Número_Processo", "Dígito_Verificador", "Ano Protocolo", "Poder Judiciário", "Região", "Unidade Judiciária"}
	empData = append(empData, head)

	for i := 0; i < len(cnj); i++ {
		bl := strconv.FormatBool(cnj[i].ValidCNJ)
		final := []string{
			cnj[i].ReceivedCNJ,
			bl,
			cnj[i].CorrectCNJ,
			cnj[i].ValidDigit,
			cnj[i].Segment1,
			cnj[i].Segment2,
			cnj[i].SourceUnit1,
			cnj[i].SourceUnit2,
			cnj[i].Detailed.District,
			cnj[i].Court1,
			cnj[i].Court2,
			cnj[i].Detailed.UF,
			cnj[i].Detailed.LawsuitNumber,
			cnj[i].Detailed.VerifyingDigit,
			cnj[i].Detailed.ProtocolYear,
			cnj[i].Detailed.Segment,
			cnj[i].Detailed.Court,
			cnj[i].Detailed.SourceUnit,
		}
		empData = append(empData, final)
	}

	csvFile, _ := create(nameFolder + "/" + nameFile + ".csv")
	csvWriter := csv.NewWriter(csvFile)

	for _, empRow := range empData {
		_ = csvWriter.Write(empRow)
	}
	csvWriter.Flush()
	err := csvFile.Close()
	Error.CheckError(err)
}