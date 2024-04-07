package middleware

import (
	"net/http"

	"github.com/twin-te/twin-te/back/apperr"
	autherr "github.com/twin-te/twin-te/back/module/auth/err"
	sharederr "github.com/twin-te/twin-te/back/module/shared/err"
)

var AppErrorCodeToHttpStatusCode = map[apperr.Code]int{
	autherr.CodeUserAuthenticationAlreadyExists:                http.StatusConflict,
	autherr.CodeUserHasAtMostOneAuthenticationFromSameProvider: http.StatusPreconditionFailed,

	sharederr.CodeInvalidArgument: http.StatusBadRequest,
	sharederr.CodeUnauthenticated: http.StatusUnauthorized,
	sharederr.CodeUnauthorized:    http.StatusForbidden,
}
