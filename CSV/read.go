package CSV

import (
	"encoding/csv"
	"github.com/Darklabel91/CNJ_Validate/Error"
	"os"
)

func ReadCsvFile(filePath string, separator rune) ([]string, error) {
	csvFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer func(csvFile *os.File) {
		err0 := csvFile.Close()
		Error.CheckError(err0)
	}(csvFile)

	csvR := csv.NewReader(csvFile)
	csvR.Comma = separator

	csvData, err := csvR.ReadAll()
	if err != nil {
		return nil, err
	}

	var data []string
	for _, line := range csvData {
		newLine := line[0]
		data = append(data, newLine)
	}

	return data, nil
}
