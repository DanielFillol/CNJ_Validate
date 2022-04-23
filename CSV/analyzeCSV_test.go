package CSV

import (
	"github.com/Darklabel91/CNJ_Validate/CNJ"
	"os"
	"reflect"
	"testing"
)

func TestAnalyzeCNJCSV(t *testing.T) {
	type args struct {
		rawFilePath      string
		separator        rune
		nameResultFolder string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := AnalyzeCNJCSV(tt.args.rawFilePath, tt.args.separator, tt.args.nameResultFolder); (err != nil) != tt.wantErr {
				t.Errorf("AnalyzeCNJCSV() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_createCSVs(t *testing.T) {
	type args struct {
		raw              []string
		nameResultFolder string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := createCSVs(tt.args.raw, tt.args.nameResultFolder); (err != nil) != tt.wantErr {
				t.Errorf("createCSVs() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_createFile(t *testing.T) {
	type args struct {
		p string
	}
	tests := []struct {
		name    string
		args    args
		want    *os.File
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := createFile(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("createFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createFile() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateHeaders(t *testing.T) {
	tests := []struct {
		name string
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateHeaders(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateHeaders() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateRow(t *testing.T) {
	type args struct {
		cnj CNJ.AnalysisCNJ
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateRow(tt.args.cnj); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateRow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readCsvFile(t *testing.T) {
	type args struct {
		filePath  string
		separator rune
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readCsvFile(tt.args.filePath, tt.args.separator)
			if (err != nil) != tt.wantErr {
				t.Errorf("readCsvFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readCsvFile() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_writeCSV(t *testing.T) {
	type args struct {
		fileName   string
		folderName string
		cnjRows    []CNJ.AnalysisCNJ
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := writeCSV(tt.args.fileName, tt.args.folderName, tt.args.cnjRows); (err != nil) != tt.wantErr {
				t.Errorf("writeCSV() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
