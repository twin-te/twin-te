package donationdomain_test

import (
	"testing"
	"time"

	donationdomain "github.com/twin-te/twin-te/back/module/donation/domain"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
)

func TestPaymentType_IsZero(t *testing.T) {
	if !donationdomain.PaymentType(0).IsZero() {
		t.Error("zero value should be zero")
	}
	if donationdomain.PaymentTypeOneTime.IsZero() {
		t.Error("PaymentTypeOneTime should not be zero")
	}
}

func TestPaymentStatus_IsZero(t *testing.T) {
	if !donationdomain.PaymentStatus(0).IsZero() {
		t.Error("zero value should be zero")
	}
	if donationdomain.PaymentStatusPending.IsZero() {
		t.Error("PaymentStatusPending should not be zero")
	}
}

func TestConstructPaymentHistory(t *testing.T) {
	id := idtype.PaymentHistoryID("id")
	now := time.Now()

	t.Run("success", func(t *testing.T) {
		ph, err := donationdomain.ConstructPaymentHistory(func(ph *donationdomain.PaymentHistory) error {
			ph.ID = id
			ph.Type = donationdomain.PaymentTypeOneTime
			ph.Status = donationdomain.PaymentStatusSucceeded
			ph.Amount = 100
			ph.CreatedAt = now
			return nil
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if ph.ID != id || ph.Type != donationdomain.PaymentTypeOneTime ||
			ph.Status != donationdomain.PaymentStatusSucceeded || ph.Amount != 100 || !ph.CreatedAt.Equal(now) {
			t.Errorf("got %+v", ph)
		}
	})

	t.Run("fn returns error", func(t *testing.T) {
		assertConstructFnError(t, donationdomain.ConstructPaymentHistory)
	})

	tests := []struct {
		name string
		fn   func(ph *donationdomain.PaymentHistory)
	}{
		{"missing ID", func(ph *donationdomain.PaymentHistory) {
			ph.Type = donationdomain.PaymentTypeOneTime
			ph.Status = donationdomain.PaymentStatusSucceeded
			ph.Amount = 100
			ph.CreatedAt = now
		}},
		{"missing Type", func(ph *donationdomain.PaymentHistory) {
			ph.ID = id
			ph.Status = donationdomain.PaymentStatusSucceeded
			ph.Amount = 100
			ph.CreatedAt = now
		}},
		{"missing Status", func(ph *donationdomain.PaymentHistory) {
			ph.ID = id
			ph.Type = donationdomain.PaymentTypeOneTime
			ph.Amount = 100
			ph.CreatedAt = now
		}},
		{"non-positive Amount", func(ph *donationdomain.PaymentHistory) {
			ph.ID = id
			ph.Type = donationdomain.PaymentTypeOneTime
			ph.Status = donationdomain.PaymentStatusSucceeded
			ph.Amount = 0
			ph.CreatedAt = now
		}},
		{"missing CreatedAt", func(ph *donationdomain.PaymentHistory) {
			ph.ID = id
			ph.Type = donationdomain.PaymentTypeOneTime
			ph.Status = donationdomain.PaymentStatusSucceeded
			ph.Amount = 100
		}},
	}
	assertConstructMissingFields(t, donationdomain.ConstructPaymentHistory, tests)
}
