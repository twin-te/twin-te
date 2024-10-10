// Code generated by codegen/idtype/generate.py; DO NOT EDIT.

package idtype

import "fmt"

type SubscriptionID string

func (id SubscriptionID) String() string {
	return string(id)
}

func (id SubscriptionID) IsZero() bool {
	return id == ""
}

func (id *SubscriptionID) Scan(src interface{}) error {
	switch src := src.(type) {
	case nil:
		return nil
	case string:
		*id = SubscriptionID(src)
		return nil
	default:
		return fmt.Errorf("Scan: unable to scan type %T into SubscriptionID", src)
	}
}

func ParseSubscriptionID(s string) (SubscriptionID, error) {
	if s == "" {
		return "", fmt.Errorf("failed to parse SubscriptionID %#v", s)
	}
	return SubscriptionID(s), nil
}
