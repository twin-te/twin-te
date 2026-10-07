package authdomain_test

import (
	"testing"

	authdomain "github.com/twin-te/twin-te/back/module/auth/domain"
)

func TestPermission_String(t *testing.T) {
	tests := []struct {
		name string
		p    authdomain.Permission
		want string
	}{
		{"known permission", authdomain.PermissionExecuteBatchJob, "ExecuteBatchJob"},
		{"unknown permission", authdomain.Permission(999), "Permission(999)"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.String(); got != tt.want {
				t.Errorf("String() = %q, want %q", got, tt.want)
			}
		})
	}
}
