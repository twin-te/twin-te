package sharedport

//go:generate go run golang.org/x/tools/cmd/stringer -type=Lock -trimprefix=Lock -output=lock_string.gen.go
type Lock int

const (
	LockNone Lock = iota
	LockShared
	LockExclusive
)
