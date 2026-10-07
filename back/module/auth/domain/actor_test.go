package authdomain_test

import (
	"testing"

	authdomain "github.com/twin-te/twin-te/back/module/auth/domain"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
)

func TestUnknown_HasPermission(t *testing.T) {
	u := authdomain.NewUnknown(authdomain.PermissionExecuteBatchJob)

	if !u.HasPermission(authdomain.PermissionExecuteBatchJob) {
		t.Error("expected true for granted permission")
	}
	if u.HasPermission(authdomain.Permission(999)) {
		t.Error("expected false for ungranted permission")
	}
}

func TestUnknown_AuthNUser(t *testing.T) {
	u := authdomain.NewUnknown()
	authNUser, ok := u.AuthNUser()
	if ok || authNUser != nil {
		t.Errorf("expected (nil, false), got (%v, %v)", authNUser, ok)
	}
}

func TestNewUnknown(t *testing.T) {
	u := authdomain.NewUnknown(authdomain.PermissionExecuteBatchJob)
	if len(u.Permissions) != 1 || u.Permissions[0] != authdomain.PermissionExecuteBatchJob {
		t.Errorf("got %+v", u)
	}
}

func TestAuthNUser_HasPermission(t *testing.T) {
	userID := idtype.NewUserID()
	actor := authdomain.NewAuthNUser(userID, authdomain.PermissionExecuteBatchJob)

	if !actor.HasPermission(authdomain.PermissionExecuteBatchJob) {
		t.Error("expected true for granted permission")
	}
	if actor.HasPermission(authdomain.Permission(999)) {
		t.Error("expected false for ungranted permission")
	}
}

func TestAuthNUser_AuthNUser(t *testing.T) {
	userID := idtype.NewUserID()
	actor := authdomain.NewAuthNUser(userID)

	got, ok := actor.AuthNUser()
	if !ok || got != actor {
		t.Errorf("got (%v, %v), want (%v, true)", got, ok, actor)
	}
}

func TestNewAuthNUser(t *testing.T) {
	userID := idtype.NewUserID()
	actor := authdomain.NewAuthNUser(userID, authdomain.PermissionExecuteBatchJob)

	if actor.UserID != userID {
		t.Errorf("UserID = %v, want %v", actor.UserID, userID)
	}
	if len(actor.Permissions) != 1 || actor.Permissions[0] != authdomain.PermissionExecuteBatchJob {
		t.Errorf("got %+v", actor.Permissions)
	}
}

func TestActor_InterfaceSatisfaction(t *testing.T) {
	var actors []authdomain.Actor = []authdomain.Actor{
		authdomain.NewUnknown(),
		authdomain.NewAuthNUser(idtype.NewUserID()),
	}
	if len(actors) != 2 {
		t.Fatalf("expected 2 actors")
	}
}
