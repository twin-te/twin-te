package sharedport

import "errors"

//go:generate go run golang.org/x/tools/cmd/stringer -type=Lock -trimprefix=Lock -output=lock_string.gen.go
type Lock int

const (
	LockNone Lock = iota
	LockShared
	LockExclusive
)

var ErrNotFound = errors.New("not found in repository") // For Repository.Find method
