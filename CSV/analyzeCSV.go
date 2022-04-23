package CSV

import (
	"fmt"
	"github.com/Darklabel91/CNJ_Validate/CNJ"
)

// AnalyzeCNJCSV creates a csv with every single row with AnalyzeCNJ
//from a given raw file and separator rune in a given folder
func AnalyzeCNJCSV(rawFilePath string, separator rune, nameResultFolder string) error {
	raw, err := ReadCsvFile(rawFilePath, separator)
	if err != nil {
		return err
	}

	err = createCSVs(raw, nameResultFolder)
	if err != nil {
		return err
	}

	fmt.Println("Files created")
	return nil
}

func createCSVs(raw []string, nameResultFolder string) error {
	var analyzeCNJCSV []CNJ.AnalysisCNJ

	for i := 0; i < len(raw); i++ {
		dataReturn, _ := CNJ.AnalyzeCNJ(raw[i])
		analyzeCNJCSV = append(analyzeCNJCSV, dataReturn)
	}

	err := writeCSV("filesOK", nameResultFolder, analyzeCNJCSV)
	if err != nil {
		return err
	}

	return nil
}
