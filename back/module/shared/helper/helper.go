package sharedhelper

import (
	"fmt"

	"github.com/samber/lo"
	"github.com/twin-te/twinte-back/apperr"
	sharederr "github.com/twin-te/twinte-back/module/shared/err"
)

func ValidateDuplicates[T comparable](collection []T) error {
	if duplicates := lo.FindDuplicates(collection); len(duplicates) != 0 {
		return apperr.New(sharederr.CodeInvalidArgument, fmt.Sprintf("found duplicates %+v", collection))
	}
	return nil
}

func ValidateDifference[T comparable](expected, actual []T) error {
	left, right := lo.Difference(expected, actual)
	if len(left) != 0 || len(right) != 0 {
		return apperr.New(sharederr.CodeInvalidArgument, fmt.Sprintf("expected %+v, but got %+v", expected, actual))
	}
	return nil
}
