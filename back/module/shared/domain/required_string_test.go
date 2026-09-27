package shareddomain_test

import (
	"testing"

	shareddomain "github.com/twin-te/twin-te/back/module/shared/domain"
)

func TestRequiredString_String(t *testing.T) {
	rs := shareddomain.RequiredString("hello")
	if rs.String() != "hello" {
		t.Errorf("String() = %v, want hello", rs.String())
	}
}

func TestRequiredString_StringPtr(t *testing.T) {
	t.Run("non-nil receiver", func(t *testing.T) {
		rs := shareddomain.RequiredString("hello")
		p := rs.StringPtr()
		if p == nil {
			t.Fatal("expected non-nil pointer")
		}
		if *p != "hello" {
			t.Errorf("*p = %v, want hello", *p)
		}
	})

	t.Run("nil receiver", func(t *testing.T) {
		var rs *shareddomain.RequiredString
		if got := rs.StringPtr(); got != nil {
			t.Errorf("StringPtr() = %v, want nil", got)
		}
	})
}

func TestRequiredString_IsZero(t *testing.T) {
	tests := []struct {
		name string
		rs   shareddomain.RequiredString
		want bool
	}{
		{"empty", shareddomain.RequiredString(""), true},
		{"non-empty", shareddomain.RequiredString("hello"), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.rs.IsZero(); got != tt.want {
				t.Errorf("IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewRequiredStringParser(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    shareddomain.RequiredString
		wantErr bool
	}{
		{"non-empty", "hello", shareddomain.RequiredString("hello"), false},
		{"trims whitespace", "  hello  ", shareddomain.RequiredString("hello"), false},
		{"empty string", "", "", true},
		{"whitespace only", "   ", "", true},
	}

	parser := shareddomain.NewRequiredStringParser("name")

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
