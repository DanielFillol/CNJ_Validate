package Test

import (
	"github.com/DanielFillol/CNJ_Validate/CNJ"
	"github.com/DanielFillol/CNJ_Validate/CNJCSV"
	"reflect"
	"testing"
)

func TestAnalyzeCNJ(t *testing.T) {
	type args struct {
		cnj string
	}
	tests := []struct {
		name    string
		args    args
		want    CNJ.AnalysisCNJ
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CNJ.AnalyzeCNJ(tt.args.cnj)
			if (err != nil) != tt.wantErr {
				t.Errorf("AnalyzeCNJ() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AnalyzeCNJ() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecomposeCNJ(t *testing.T) {
	type args struct {
		cnj string
	}
	tests := []struct {
		name    string
		args    args
		want    CNJ.DecomposedCNJ
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CNJ.DecomposeCNJ(tt.args.cnj)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecomposeCNJ() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecomposeCNJ() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateCNJ(t *testing.T) {
	type args struct {
		cnj string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CNJ.ValidateCNJ(tt.args.cnj)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateCNJ() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ValidateCNJ() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWriteCNJ(t *testing.T) {
	type args struct {
		number CNJ.AnalysisCNJ
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CNJ.WriteCNJ(tt.args.number)
			if (err != nil) != tt.wantErr {
				t.Errorf("WriteCNJ() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("WriteCNJ() got = %v, want %v", got, tt.want)
			}
		})
	}
}

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
			if err := CNJCSV.AnalyzeCNJCSV(tt.args.rawFilePath, tt.args.separator, tt.args.nameResultFolder); (err != nil) != tt.wantErr {
				t.Errorf("AnalyzeCNJCSV() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
