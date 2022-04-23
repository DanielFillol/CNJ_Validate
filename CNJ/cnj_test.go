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
		court   string
		segment Segment
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
			got, err := GetOriginCourt(tt.args.court, tt.args.segment)
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

func Test_getMax(t *testing.T) {
	type args struct {
		segment Segment
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getMax(tt.args.segment)
			if (err != nil) != tt.wantErr {
				t.Errorf("getMax() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getMax() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isCourtValid(t *testing.T) {
	type args struct {
		segment Segment
		court   int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isCourtValid(tt.args.segment, tt.args.court); got != tt.want {
				t.Errorf("isCourtValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isFirstDigit9(t *testing.T) {
	type args struct {
		sourceUnit string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isFirstDigit9(tt.args.sourceUnit); got != tt.want {
				t.Errorf("isFirstDigit9() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isZeroSequence(t *testing.T) {
	type args struct {
		sourceUnit string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isZeroSequence(tt.args.sourceUnit); got != tt.want {
				t.Errorf("isZeroSequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ocError(t *testing.T) {
	type args struct {
		segment Segment
		court   int
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
			if err := ocError(tt.args.segment, tt.args.court); (err != nil) != tt.wantErr {
				t.Errorf("ocError() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_parseCourt00(t *testing.T) {
	type args struct {
		segment Segment
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
			got, err := parseCourt00(tt.args.segment)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseCourt00() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseCourt00() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseCourt90(t *testing.T) {
	type args struct {
		segment Segment
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
			got, err := parseCourt90(tt.args.segment)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseCourt90() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseCourt90() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseCourtOther(t *testing.T) {
	type args struct {
		segment    Segment
		orginCourt string
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
			got, err := parseCourtOther(tt.args.segment, tt.args.orginCourt)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseCourtOther() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseCourtOther() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseMartialCourt(t *testing.T) {
	type args struct {
		oc int
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
			got, err := parseMartialCourt(tt.args.oc)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseMartialCourt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseMartialCourt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseSourceUnit(t *testing.T) {
	type args struct {
		segment    Segment
		sourceUnit string
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
			got, err := parseSourceUnit(tt.args.segment, tt.args.sourceUnit)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseSourceUnit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseSourceUnit() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validVD(t *testing.T) {
	type args struct {
		argNumber string
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
			got, err := validVD(tt.args.argNumber)
			if (err != nil) != tt.wantErr {
				t.Errorf("validVD() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("validVD() got = %v, want %v", got, tt.want)
			}
		})
	}
}
