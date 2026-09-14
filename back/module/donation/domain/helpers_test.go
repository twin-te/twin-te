package donationdomain_test

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
)

func newUserID() idtype.UserID {
	return idtype.UserID(uuid.New())
}

// assertConstructFnError verifies that a Construct function propagates
// the error returned by its builder function.
func assertConstructFnError[T any](t *testing.T, construct func(fn func(*T) error) (*T, error)) {
	t.Helper()
	wantErr := errors.New("boom")
	_, err := construct(func(v *T) error {
		return wantErr
	})
	if !errors.Is(err, wantErr) {
		t.Errorf("err = %v, want %v", err, wantErr)
	}
}

// assertConstructMissingFields table-drivenly verifies that a Construct
// function returns an error when a required field is left unset by the builder.
func assertConstructMissingFields[T any](t *testing.T, construct func(fn func(*T) error) (*T, error), tests []struct {
	name string
	fn   func(*T)
}) {
	t.Helper()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := construct(func(v *T) error {
				tt.fn(v)
				return nil
			})
			if err == nil {
				t.Error("expected error, got nil")
			}
		})
	}
}
