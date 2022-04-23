package CSV

import (
	"fmt"
	"github.com/Darklabel91/CNJ_Validate/CNJ"
)

// AnalyzeCNJCSV creates a csv with every single row with AnalyzeCNJ
//from a given raw file and separator rune in a given folder
func AnalyzeCNJCSV(rawFilePath string, separator rune, nameResultFolder string) error {
	raw, err := readCsvFile(rawFilePath, separator)
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

//execute the AnalyzeCNJ from a []string
func createCSVs(raw []string, nameResultFolder string) error {
	var analyzeCNJCSV []CNJ.AnalysisCNJ

	for _, cnj := range raw {
		dataReturn, _ := CNJ.AnalyzeCNJ(cnj)
		analyzeCNJCSV = append(analyzeCNJCSV, dataReturn)
	}

	err := writeCSV("filesOK", nameResultFolder, analyzeCNJCSV)
	if err != nil {
		return err
	}

	return nil
}
