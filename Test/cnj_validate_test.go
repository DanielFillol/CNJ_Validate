package Test

import (
	"github.com/Darklabel91/CNJ_Validate"
	"github.com/Darklabel91/CNJ_Validate/Structs"
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
		want    Structs.AnalysisCNJ
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CNJ_Validate.AnalyzeCNJ(tt.args.cnj)
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
			if err := CNJ_Validate.AnalyzeCNJCSV(tt.args.rawFilePath, tt.args.separator, tt.args.nameResultFolder); (err != nil) != tt.wantErr {
				t.Errorf("AnalyzeCNJCSV() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCNJWrite(t *testing.T) {
	type args struct {
		number Structs.AnalysisCNJ
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CNJ_Validate.CNJWrite(tt.args.number); got != tt.want {
				t.Errorf("CNJWrite() = %v, want %v", got, tt.want)
			}
		})
	}
}
