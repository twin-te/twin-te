package authdomain_test

import (
	"errors"
	"reflect"
	"testing"

	authdomain "github.com/twin-te/twin-te/back/module/auth/domain"
	autherr "github.com/twin-te/twin-te/back/module/auth/err"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
	sharederr "github.com/twin-te/twin-te/back/module/shared/err"

	"github.com/twin-te/twin-te/back/apperr"
)

func TestUser_Clone(t *testing.T) {
	u := newTestUser()

	clone := u.Clone()

	if clone.ID != u.ID || !reflect.DeepEqual(u.Authentications, clone.Authentications) {
		t.Fatalf("clone = %+v, want equal to %+v", clone, u)
	}

	clone.Authentications[0] = authdomain.NewUserAuthentication(authdomain.ProviderApple, "other")
	if u.Authentications[0] == clone.Authentications[0] {
		t.Error("Clone should copy the Authentications slice, not share it")
	}
}

func TestUser_BeforeUpdateHook(t *testing.T) {
	u := newTestUser()
	u.BeforeUpdateHook()

	before, ok := u.BeforeUpdated.Get()
	if !ok {
		t.Fatal("expected BeforeUpdated to be set")
	}
	if before.ID != u.ID {
		t.Errorf("BeforeUpdated.ID = %v, want %v", before.ID, u.ID)
	}
}

func TestUser_AddAuthentication(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		u := newTestUser()
		newAuth := authdomain.NewUserAuthentication(authdomain.ProviderApple, "apple-id")

		err := u.AddAuthentication(newAuth)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(u.Authentications) != 2 {
			t.Fatalf("got %d authentications, want 2", len(u.Authentications))
		}
	})

	t.Run("duplicate provider", func(t *testing.T) {
		u := newTestUser()
		dup := authdomain.NewUserAuthentication(authdomain.ProviderGoogle, "another-social-id")

		err := u.AddAuthentication(dup)
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !apperr.Is(err, autherr.CodeUserHasAtMostOneAuthenticationFromSameProvider) {
			t.Errorf("err = %v, want code %v", err, autherr.CodeUserHasAtMostOneAuthenticationFromSameProvider)
		}
	})
}

func TestUser_DeleteAuthentication(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		u := newTestUser()
		u.Authentications = append(u.Authentications, authdomain.NewUserAuthentication(authdomain.ProviderApple, "apple-id"))

		err := u.DeleteAuthentication(authdomain.ProviderApple)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(u.Authentications) != 1 {
			t.Fatalf("got %d authentications, want 1", len(u.Authentications))
		}
		if u.Authentications[0].Provider != authdomain.ProviderGoogle {
			t.Errorf("remaining authentication = %+v", u.Authentications[0])
		}
	})

	t.Run("last authentication cannot be deleted", func(t *testing.T) {
		u := newTestUser()

		err := u.DeleteAuthentication(authdomain.ProviderGoogle)
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !apperr.Is(err, autherr.CodeUserHasAtLeastOneAuthentication) {
			t.Errorf("err = %v, want code %v", err, autherr.CodeUserHasAtLeastOneAuthentication)
		}
	})

	t.Run("provider not associated", func(t *testing.T) {
		u := newTestUser()
		u.Authentications = append(u.Authentications, authdomain.NewUserAuthentication(authdomain.ProviderApple, "apple-id"))

		err := u.DeleteAuthentication(authdomain.ProviderTwitter)
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !apperr.Is(err, sharederr.CodeInvalidArgument) {
			t.Errorf("err = %v, want code %v", err, sharederr.CodeInvalidArgument)
		}
	})
}

func TestConstructUser(t *testing.T) {
	id := idtype.NewUserID()
	auth := newTestUserAuthentication()

	t.Run("success", func(t *testing.T) {
		u, err := authdomain.ConstructUser(func(u *authdomain.User) error {
			u.ID = id
			u.Authentications = []authdomain.UserAuthentication{auth}
			return nil
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if u.ID != id || !reflect.DeepEqual(u.Authentications, []authdomain.UserAuthentication{auth}) {
			t.Errorf("got %+v", u)
		}
	})

	t.Run("fn returns error", func(t *testing.T) {
		wantErr := errors.New("boom")
		_, err := authdomain.ConstructUser(func(u *authdomain.User) error {
			return wantErr
		})
		if !errors.Is(err, wantErr) {
			t.Errorf("err = %v, want %v", err, wantErr)
		}
	})

	t.Run("missing authentications", func(t *testing.T) {
		_, err := authdomain.ConstructUser(func(u *authdomain.User) error {
			u.ID = id
			return nil
		})
		if err == nil {
			t.Error("expected error, got nil")
		}
	})

	t.Run("missing id", func(t *testing.T) {
		_, err := authdomain.ConstructUser(func(u *authdomain.User) error {
			u.Authentications = []authdomain.UserAuthentication{auth}
			return nil
		})
		if err == nil {
			t.Error("expected error, got nil")
		}
	})
}
