package calendardomain

import (
	"fmt"

	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
	sharederr "github.com/twin-te/twin-te/back/module/shared/err"
	sharedhelper "github.com/twin-te/twin-te/back/module/shared/helper"
)

// IcalSubscriptionMode は TargetTagIDs が付与されたコースを iCal フィードでどう出力するかを制御する。
type IcalSubscriptionMode string

const (
	// IcalSubscriptionModeSync はすべてのコースを通常出力する。TargetTagIDs は空でなければならない。
	IcalSubscriptionModeSync IcalSubscriptionMode = "sync"
	// IcalSubscriptionModeExclude は TargetTagIDs のいずれかのタグを持つコースを出力しない。
	IcalSubscriptionModeExclude IcalSubscriptionMode = "exclude"
	// IcalSubscriptionModeTransparent は TargetTagIDs のいずれかのタグを持つコースに TRANSP:TRANSPARENT を付与する。
	IcalSubscriptionModeTransparent IcalSubscriptionMode = "transparent"
)

func ParseIcalSubscriptionMode(s string) (IcalSubscriptionMode, error) {
	switch IcalSubscriptionMode(s) {
	case IcalSubscriptionModeSync, IcalSubscriptionModeExclude, IcalSubscriptionModeTransparent:
		return IcalSubscriptionMode(s), nil
	default:
		return "", sharederr.NewInvalidArgument("invalid ical subscription mode %q", s)
	}
}

func (m IcalSubscriptionMode) String() string {
	return string(m)
}

// IcalSubscription はユーザーごとの iCal フィード設定。
type IcalSubscription struct {
	ID           idtype.IcalSubscriptionID
	UserID       idtype.UserID
	Mode         IcalSubscriptionMode
	TargetTagIDs []idtype.TagID
}

func ConstructIcalSubscription(fn func(s *IcalSubscription) (err error)) (*IcalSubscription, error) {
	s := new(IcalSubscription)
	if err := fn(s); err != nil {
		return nil, err
	}

	if s.ID.IsZero() || s.UserID.IsZero() {
		return nil, fmt.Errorf("failed to construct %+v", s)
	}

	if _, err := ParseIcalSubscriptionMode(s.Mode.String()); err != nil {
		return nil, err
	}

	if err := s.validateTargetTagIDs(); err != nil {
		return nil, err
	}

	return s, nil
}

// validateTargetTagIDs は TargetTagIDs の不変条件を検証する:
//   - 重複が無いこと
//   - Mode が sync のときは空であること
//
// タグの所有者（UserID に属すること）の検証はユースケース層で行う。
func (s *IcalSubscription) validateTargetTagIDs() error {
	if err := sharedhelper.ValidateDuplicates(s.TargetTagIDs); err != nil {
		return err
	}
	if s.Mode == IcalSubscriptionModeSync && len(s.TargetTagIDs) != 0 {
		return sharederr.NewInvalidArgument("target tag ids must be empty when mode is sync")
	}
	return nil
}

// IsTransparent は指定したタグ ID を持つコースを TRANSP:TRANSPARENT にすべきかを返す。
func (s *IcalSubscription) IsTransparent(courseTagIDs []idtype.TagID) bool {
	return s.Mode == IcalSubscriptionModeTransparent && s.intersectsTargets(courseTagIDs)
}

// IsExcluded は指定したタグ ID を持つコースをフィードから除外すべきかを返す。
func (s *IcalSubscription) IsExcluded(courseTagIDs []idtype.TagID) bool {
	return s.Mode == IcalSubscriptionModeExclude && s.intersectsTargets(courseTagIDs)
}

func (s *IcalSubscription) intersectsTargets(courseTagIDs []idtype.TagID) bool {
	if len(s.TargetTagIDs) == 0 || len(courseTagIDs) == 0 {
		return false
	}
	targets := make(map[idtype.TagID]struct{}, len(s.TargetTagIDs))
	for _, t := range s.TargetTagIDs {
		targets[t] = struct{}{}
	}
	for _, t := range courseTagIDs {
		if _, ok := targets[t]; ok {
			return true
		}
	}
	return false
}
