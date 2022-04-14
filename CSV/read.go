package CSV

import (
	"encoding/csv"
	"github.com/Darklabel91/CNJ_Validate/Error"
	"os"
)

func ReadCsvFile(filePath string, separator rune) []string {
	var data []string

	csvFile, err := os.Open(filePath)
	Error.CheckError(err)

	defer func(csvFile *os.File) {
		err0 := csvFile.Close()
		Error.CheckError(err0)
	}(csvFile)

	csvLines := csv.NewReader(csvFile)
	csvLines.Comma = separator
	csvData, err1 := csvLines.ReadAll()
	Error.CheckError(err1)

	for _, line := range csvData {
		emp := line[0]
		data = append(data, emp)
	}

	return data
}
