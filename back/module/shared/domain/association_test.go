package shareddomain_test

import (
	"testing"

	shareddomain "github.com/twin-te/twin-te/back/module/shared/domain"
)

func TestAssociation_ZeroValue(t *testing.T) {
	var a shareddomain.Association[int]

	if !a.IsAbsent() {
		t.Error("zero value Association should be absent")
	}
	if a.IsPresent() {
		t.Error("zero value Association should not be present")
	}

	if _, ok := a.Get(); ok {
		t.Error("Get() on absent Association should return ok=false")
	}
}

func TestAssociation_Set(t *testing.T) {
	var a shareddomain.Association[string]
	a.Set("hello")

	if !a.IsPresent() {
		t.Error("expected IsPresent() = true after Set")
	}
	if a.IsAbsent() {
		t.Error("expected IsAbsent() = false after Set")
	}

	got, ok := a.Get()
	if !ok || got != "hello" {
		t.Errorf("Get() = (%v, %v), want (hello, true)", got, ok)
	}

	if a.MustGet() != "hello" {
		t.Errorf("MustGet() = %v, want hello", a.MustGet())
	}
}

func TestAssociation_MustGet_Panics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected MustGet() to panic on absent Association")
		}
	}()

	var a shareddomain.Association[int]
	a.MustGet()
}
