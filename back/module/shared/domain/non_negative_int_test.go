package shareddomain_test

import (
	"testing"

	shareddomain "github.com/twin-te/twin-te/back/module/shared/domain"
)

func TestNonNegativeInt_Int(t *testing.T) {
	nni := shareddomain.NonNegativeInt(5)
	if nni.Int() != 5 {
		t.Errorf("Int() = %v, want 5", nni.Int())
	}
}

func TestNewNonNegativeIntParser(t *testing.T) {
	tests := []struct {
		name    string
		input   int
		want    shareddomain.NonNegativeInt
		wantErr bool
	}{
		{"positive", 5, shareddomain.NonNegativeInt(5), false},
		{"zero", 0, shareddomain.NonNegativeInt(0), false},
		{"negative", -1, 0, true},
	}

	parser := shareddomain.NewNonNegativeIntParser("count")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parser(tt.input)
			if (err != nil) != tt.wantErr {
				t.Fatalf("err = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("got = %v, want %v", got, tt.want)
			}
		})
	}
}
