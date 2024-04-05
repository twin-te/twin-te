package interceptor

import (
	"context"
	"net/http"

	"github.com/bufbuild/connect-go"
	"github.com/samber/lo"
	"github.com/twin-te/twin-te/back/appenv"
	"github.com/twin-te/twin-te/back/apperr"
	authmodule "github.com/twin-te/twin-te/back/module/auth"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
)

func NewErrorInterceptor() connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			res, err := next(ctx, req)
			if aerr, ok := apperr.As(err); ok {
				if connectErrorCode, ok := AppErrorCodeToConnectErrorCode[aerr.Code]; ok {
					err = connect.NewError(connectErrorCode, err)
				}
			}
			return res, err
		})
	}
	return connect.UnaryInterceptorFunc(interceptor)
}

func NewAuthInterceptor(accessController authmodule.AccessController) connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			sessionID, ok := getSessionIDFromHeader(req.Header())
			ctx, err := accessController.WithActor(ctx, lo.Ternary(ok, &sessionID, nil))
			if err != nil {
				return nil, err
			}
			return next(ctx, req)
		})
	}
	return connect.UnaryInterceptorFunc(interceptor)
}

func getSessionIDFromHeader(header http.Header) (id idtype.SessionID, ok bool) {
	request := http.Request{Header: header}
	cookie, err := request.Cookie(appenv.COOKIE_SESSION_NAME)
	if err != nil {
		return
	}
	id, err = idtype.ParseSessionID(cookie.Value)
	if err != nil {
		return
	}
	return id, true
}
