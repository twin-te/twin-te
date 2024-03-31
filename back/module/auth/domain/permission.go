package authdomain

// Permission is required to execute the specific use cases.
//
//go:generate go run golang.org/x/tools/cmd/stringer -type=Permission -trimprefix=Permission -output=permission_string.gen.go
type Permission int

const (
	PermissionExecuteBatchJob Permission = iota + 1
)
