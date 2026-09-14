package donationdomain_test

import (
	"testing"

	"github.com/samber/mo"
	donationdomain "github.com/twin-te/twin-te/back/module/donation/domain"
	shareddomain "github.com/twin-te/twin-te/back/module/shared/domain"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
)

func TestParseDisplayName(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		got, err := donationdomain.ParseDisplayName("name")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got.String() != "name" {
			t.Errorf("got %v, want %v", got, "name")
		}
	})

	t.Run("empty string", func(t *testing.T) {
		_, err := donationdomain.ParseDisplayName("")
		if err == nil {
			t.Error("expected error, got nil")
		}
	})
}

func TestLink_String(t *testing.T) {
	l := donationdomain.Link("https://example.com")
	if l.String() != "https://example.com" {
		t.Errorf("got %v", l.String())
	}
}

func TestParseLink(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{"valid absolute url", "https://example.com", false},
		{"empty string", "", true},
		{"relative path", "/path", true},
		{"malformed", "://bad-url", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := donationdomain.ParseLink(tt.input)
			if (err != nil) != tt.wantErr {
				t.Fatalf("err = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && got.String() != tt.input {
				t.Errorf("got %v, want %v", got, tt.input)
			}
		})
	}
}

func TestPaymentUser_Clone(t *testing.T) {
	displayName, _ := donationdomain.ParseDisplayName("name")
	pu := &donationdomain.PaymentUser{
		ID:          idtype.PaymentUserID("id"),
		UserID:      newUserID(),
		DisplayName: mo.Some(displayName),
		Link:        mo.Some(donationdomain.Link("https://example.com")),
	}

	clone := pu.Clone()

	if clone.ID != pu.ID || clone.UserID != pu.UserID {
		t.Fatalf("clone = %+v, want equal to %+v", clone, pu)
	}
	if clone.DisplayName.MustGet() != pu.DisplayName.MustGet() {
		t.Errorf("clone.DisplayName = %v, want %v", clone.DisplayName, pu.DisplayName)
	}
	if clone.Link.MustGet() != pu.Link.MustGet() {
		t.Errorf("clone.Link = %v, want %v", clone.Link, pu.Link)
	}

	clone.DisplayName = mo.None[shareddomain.RequiredString]()
	if pu.DisplayName.IsAbsent() {
		t.Error("Clone should copy the struct, not share state")
	}
}

func TestPaymentUser_BeforeUpdateHook(t *testing.T) {
	pu := &donationdomain.PaymentUser{
		ID:     idtype.PaymentUserID("id"),
		UserID: newUserID(),
	}

	pu.BeforeUpdateHook()

	before, ok := pu.BeforeUpdated.Get()
	if !ok {
		t.Fatal("expected BeforeUpdated to be set")
	}
	if before.ID != pu.ID || before.UserID != pu.UserID {
		t.Errorf("before = %+v, want equal to %+v", before, pu)
	}
}

func TestPaymentUser_Update(t *testing.T) {
	displayName, _ := donationdomain.ParseDisplayName("name")
	newDisplayName, _ := donationdomain.ParseDisplayName("new name")
	link := donationdomain.Link("https://example.com")

	t.Run("updates provided fields", func(t *testing.T) {
		pu := &donationdomain.PaymentUser{
			ID:          idtype.PaymentUserID("id"),
			UserID:      newUserID(),
			DisplayName: mo.Some(displayName),
			Link:        mo.Some(link),
		}

		pu.Update(donationdomain.PaymentUserDataToUpdate{
			DisplayName: mo.Some(mo.Some(newDisplayName)),
		})

		if pu.DisplayName.MustGet() != newDisplayName {
			t.Errorf("DisplayName = %v, want %v", pu.DisplayName, newDisplayName)
		}
		if pu.Link.MustGet() != link {
			t.Errorf("Link should remain unchanged, got %v", pu.Link)
		}
	})

	t.Run("clears field when set to None", func(t *testing.T) {
		pu := &donationdomain.PaymentUser{
			ID:          idtype.PaymentUserID("id"),
			UserID:      newUserID(),
			DisplayName: mo.Some(displayName),
			Link:        mo.Some(link),
		}

		pu.Update(donationdomain.PaymentUserDataToUpdate{
			Link: mo.Some(mo.None[donationdomain.Link]()),
		})

		if pu.Link.IsPresent() {
			t.Errorf("Link should be cleared, got %v", pu.Link)
		}
	})

	t.Run("no-op when data fields absent", func(t *testing.T) {
		pu := &donationdomain.PaymentUser{
			ID:          idtype.PaymentUserID("id"),
			UserID:      newUserID(),
			DisplayName: mo.Some(displayName),
			Link:        mo.Some(link),
		}

		pu.Update(donationdomain.PaymentUserDataToUpdate{})

		if pu.DisplayName.MustGet() != displayName || pu.Link.MustGet() != link {
			t.Errorf("fields should remain unchanged, got %+v", pu)
		}
	})
}

func TestConstructPaymentUser(t *testing.T) {
	id := idtype.PaymentUserID("id")
	userID := newUserID()

	t.Run("success", func(t *testing.T) {
		pu, err := donationdomain.ConstructPaymentUser(func(pu *donationdomain.PaymentUser) error {
			pu.ID = id
			pu.UserID = userID
			return nil
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if pu.ID != id || pu.UserID != userID {
			t.Errorf("got %+v", pu)
		}
	})

	t.Run("fn returns error", func(t *testing.T) {
		assertConstructFnError(t, donationdomain.ConstructPaymentUser)
	})

	tests := []struct {
		name string
		fn   func(pu *donationdomain.PaymentUser)
	}{
		{"missing ID", func(pu *donationdomain.PaymentUser) {
			pu.UserID = userID
		}},
		{"missing UserID", func(pu *donationdomain.PaymentUser) {
			pu.ID = id
		}},
	}
	assertConstructMissingFields(t, donationdomain.ConstructPaymentUser, tests)
}
