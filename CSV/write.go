package CSV

import (
	"encoding/csv"
	"github.com/Darklabel91/CNJ_Validate/CNJ"
	"os"
	"path/filepath"
	"strconv"
)

//ExportCSV exports a csv to a given folder, with a given name from a collection of AnalysisCNJ
func writeCSV(fileName string, folderName string, cnjRows []CNJ.AnalysisCNJ) error {
	var rows [][]string

	rows = append(rows, generateHeaders())

	for _, cnj := range cnjRows {
		rows = append(rows, generateRow(cnj))
	}

	cf, err := createFile(folderName + "/" + fileName + ".csv")
	if err != nil {
		return err
	}

	defer cf.Close()

	w := csv.NewWriter(cf)

	err = w.WriteAll(rows)
	if err != nil {
		return err
	}

	return nil
}

// create csv file from operating system
func createFile(p string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		return nil, err
	}
	return os.Create(p)
}

// generate the necessary headers for csv file
func generateHeaders() []string {
	return []string{
		"CNJ_original",
		"CNJ_é_válido",
		"Segmento",
		"Segmento_short",
		"Tipo_Unidade_Judiciária",
		"Número_Unidade_Judiciária",
		"Nome_Unidade_Judiciária",
		"Tipo_Região",
		"Número_Região",
		"Nome_Região",
		"Número_Processo",
		"Dígito_Verificador",
		"Ano Protocolo",
		"Poder Judiciário",
		"Região",
		"Unidade Judiciária",
	}
}

// returns a []string that compose the row in the csv file
func generateRow(cnj CNJ.AnalysisCNJ) []string {
	return []string{
		cnj.ReceivedCNJ,
		strconv.FormatBool(cnj.ValidCNJ),
		cnj.SegmentName,
		cnj.SegmentShort,
		cnj.SourceUnitType,
		cnj.SourceUnitNumber,
		cnj.Detailed.District,
		cnj.CourtType,
		cnj.CourtNumber,
		cnj.Detailed.UF,
		cnj.Detailed.LawsuitNumber,
		cnj.Detailed.VerifyingDigit,
		cnj.Detailed.ProtocolYear,
		cnj.Detailed.Segment,
		cnj.Detailed.Court,
		cnj.Detailed.SourceUnit,
	}
}
