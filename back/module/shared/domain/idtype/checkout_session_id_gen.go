// Code generated by codegen/idtype/generate.py; DO NOT EDIT.

package idtype

import "fmt"

type CheckoutSessionID string

func (id CheckoutSessionID) String() string {
	return string(id)
}

func (id CheckoutSessionID) IsZero() bool {
	return id == ""
}

func ParseCheckoutSessionID(s string) (CheckoutSessionID, error) {
	if s == "" {
		return "", fmt.Errorf("failed to parse CheckoutSessionID %v", s)
	}
	return CheckoutSessionID(s), nil
}
