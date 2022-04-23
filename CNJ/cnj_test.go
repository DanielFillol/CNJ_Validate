package CNJ

import (
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
		want    AnalysisCNJ
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AnalyzeCNJ(tt.args.cnj)
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
		want    DecomposedCNJ
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DecomposeCNJ(tt.args.cnj)
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

func TestGetOriginCourt(t *testing.T) {
	type args struct {
		cnj string
	}
	tests := []struct {
		name    string
		args    args
		want    OriginCourt
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetOriginCourt(tt.args.cnj)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOriginCourt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOriginCourt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSegment(t *testing.T) {
	type args struct {
		segment string
	}
	tests := []struct {
		name    string
		args    args
		want    Segment
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetSegment(tt.args.segment)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSegment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSegment() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSourceUnit(t *testing.T) {
	type args struct {
		sourceUnit string
		segment    Segment
	}
	tests := []struct {
		name    string
		args    args
		want    SourceUnit
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetSourceUnit(tt.args.sourceUnit, tt.args.segment)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSourceUnit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSourceUnit() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidVD(t *testing.T) {
	type args struct {
		cnj string
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
			got, err := ValidVD(tt.args.cnj)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidVD() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ValidVD() got = %v, want %v", got, tt.want)
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
			got, err := ValidateCNJ(tt.args.cnj)
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
		number AnalysisCNJ
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
			if got := WriteCNJ(tt.args.number); got != tt.want {
				t.Errorf("WriteCNJ() = %v, want %v", got, tt.want)
			}
		})
	}
}
