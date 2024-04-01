package sharederr

import (
	"fmt"

	"github.com/twin-te/twinte-back/apperr"
)

func NewInvalidArgument(format string, a ...any) *apperr.Error {
	return apperr.New(
		CodeInvalidArgument,
		fmt.Sprintf(format, a...),
	)
}
