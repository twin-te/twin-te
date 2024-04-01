package sharedhelper

import (
	"github.com/samber/lo"
	sharederr "github.com/twin-te/twinte-back/module/shared/err"
)

func ValidateDuplicates[T comparable](collection []T) error {
	if duplicates := lo.FindDuplicates(collection); len(duplicates) != 0 {
		return sharederr.NewInvalidArgument("found duplicates %+v", duplicates)
	}
	return nil
}

func ValidateDifference[T comparable](expected, actual []T) error {
	left, right := lo.Difference(expected, actual)
	if len(left) != 0 || len(right) != 0 {
		return sharederr.NewInvalidArgument("expected %+v, but got %+v", expected, actual)
	}
	return nil
}
