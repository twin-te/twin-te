package donationdomain_test

import (
	"testing"
	"time"

	donationdomain "github.com/twin-te/twin-te/back/module/donation/domain"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
)

func TestSubscriptionPlan_Clone(t *testing.T) {
	sp := &donationdomain.SubscriptionPlan{
		ID:     idtype.SubscriptionPlanID("id"),
		Name:   "plan",
		Amount: 500,
	}

	clone := sp.Clone()

	if *clone != *sp {
		t.Fatalf("clone = %+v, want equal to %+v", clone, sp)
	}

	clone.Name = "changed"
	if sp.Name == clone.Name {
		t.Error("Clone should copy the struct, not share it")
	}
}

func TestConstructSubscriptionPlan(t *testing.T) {
	id := idtype.SubscriptionPlanID("id")

	t.Run("success", func(t *testing.T) {
		sp, err := donationdomain.ConstructSubscriptionPlan(func(sp *donationdomain.SubscriptionPlan) error {
			sp.ID = id
			sp.Name = "plan"
			sp.Amount = 500
			return nil
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if sp.ID != id || sp.Name != "plan" || sp.Amount != 500 {
			t.Errorf("got %+v", sp)
		}
	})

	t.Run("fn returns error", func(t *testing.T) {
		assertConstructFnError(t, donationdomain.ConstructSubscriptionPlan)
	})

	tests := []struct {
		name string
		fn   func(sp *donationdomain.SubscriptionPlan)
	}{
		{"missing ID", func(sp *donationdomain.SubscriptionPlan) {
			sp.Name = "plan"
			sp.Amount = 500
		}},
		{"missing Name", func(sp *donationdomain.SubscriptionPlan) {
			sp.ID = id
			sp.Amount = 500
		}},
		{"non-positive Amount", func(sp *donationdomain.SubscriptionPlan) {
			sp.ID = id
			sp.Name = "plan"
			sp.Amount = 0
		}},
	}
	assertConstructMissingFields(t, donationdomain.ConstructSubscriptionPlan, tests)
}

func TestConstructSubscription(t *testing.T) {
	id := idtype.SubscriptionID("id")
	paymentUserID := idtype.PaymentUserID("pu-id")
	planID := idtype.SubscriptionPlanID("plan-id")
	now := time.Now()

	t.Run("success", func(t *testing.T) {
		s, err := donationdomain.ConstructSubscription(func(s *donationdomain.Subscription) error {
			s.ID = id
			s.PaymentUserID = paymentUserID
			s.PlanID = planID
			s.IsActive = true
			s.CreatedAt = now
			return nil
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if s.ID != id || s.PaymentUserID != paymentUserID || s.PlanID != planID ||
			!s.IsActive || !s.CreatedAt.Equal(now) {
			t.Errorf("got %+v", s)
		}
	})

	t.Run("fn returns error", func(t *testing.T) {
		assertConstructFnError(t, donationdomain.ConstructSubscription)
	})

	tests := []struct {
		name string
		fn   func(s *donationdomain.Subscription)
	}{
		{"missing ID", func(s *donationdomain.Subscription) {
			s.PaymentUserID = paymentUserID
			s.PlanID = planID
			s.CreatedAt = now
		}},
		{"missing PaymentUserID", func(s *donationdomain.Subscription) {
			s.ID = id
			s.PlanID = planID
			s.CreatedAt = now
		}},
		{"missing PlanID", func(s *donationdomain.Subscription) {
			s.ID = id
			s.PaymentUserID = paymentUserID
			s.CreatedAt = now
		}},
		{"missing CreatedAt", func(s *donationdomain.Subscription) {
			s.ID = id
			s.PaymentUserID = paymentUserID
			s.PlanID = planID
		}},
	}
	assertConstructMissingFields(t, donationdomain.ConstructSubscription, tests)
}
