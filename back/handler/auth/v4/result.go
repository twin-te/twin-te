package authv4

import (
	"fmt"
	"log"

	"github.com/samber/lo"
	"github.com/twin-te/twin-te/back/apperr"
	autherr "github.com/twin-te/twin-te/back/module/auth/err"
	sharederr "github.com/twin-te/twin-te/back/module/shared/err"
)

// oauth2Result represents the successful result of connecting an authentication.
// It is passed to the front end as the query parameter "connect_result".
type oauth2Result string

const (
	oauth2ResultConnected oauth2Result = "connected"

	// The authentication has already been added to the user.
	oauth2ResultAlreadyConnected oauth2Result = "already_connected"
)

// oauth2ErrorCode represents the reason why the OAuth 2.0 flow failed.
// It is passed to the front end as the query parameter "connect_error" or "login_error".
type oauth2ErrorCode string

const (
	oauth2ErrorCodeCancelled       oauth2ErrorCode = "cancelled"
	oauth2ErrorCodeInvalidProvider oauth2ErrorCode = "invalid_provider"
	oauth2ErrorCodeInvalidState    oauth2ErrorCode = "invalid_state"
	oauth2ErrorCodeUnauthenticated oauth2ErrorCode = "unauthenticated"

	// The authentication has already been added to another user.
	oauth2ErrorCodeAlreadyUsedByAnotherUser oauth2ErrorCode = "already_used_by_another_user"

	// Another authentication from the same provider has already been added to the user.
	oauth2ErrorCodeProviderAlreadyConnected oauth2ErrorCode = "provider_already_connected"

	oauth2ErrorCodeFailed oauth2ErrorCode = "failed"
)

type oauth2Error struct {
	code oauth2ErrorCode
	err  error
}

func (e *oauth2Error) Error() string {
	return fmt.Sprintf("[%s] %v", e.code, e.err)
}

func (e *oauth2Error) Unwrap() error {
	return e.err
}

func newOAuth2Error(code oauth2ErrorCode, err error) *oauth2Error {
	return &oauth2Error{code: code, err: err}
}

// getOAuth2ErrorCodeFromProviderError returns the error code corresponding to the error returned by the provider.
func getOAuth2ErrorCodeFromProviderError(providerError string) oauth2ErrorCode {
	switch providerError {
	// access_denied is defined in RFC 6749, and user_cancelled_authorize is returned by Apple.
	case "access_denied", "user_cancelled_authorize":
		return oauth2ErrorCodeCancelled
	default:
		return oauth2ErrorCodeFailed
	}
}

// getOAuth2ErrorCode returns the error code to be passed to the front end.
// Unexpected errors are logged, since they are not returned as the response.
func getOAuth2ErrorCode(err error) oauth2ErrorCode {
	code := oauth2ErrorCodeFailed

	if oerr, ok := lo.ErrorsAs[*oauth2Error](err); ok {
		code = oerr.code
	} else if aerr, ok := apperr.As(err); ok {
		switch aerr.Code {
		case autherr.CodeUserAuthenticationAlreadyExists:
			code = oauth2ErrorCodeAlreadyUsedByAnotherUser
		case autherr.CodeUserHasAtMostOneAuthenticationFromSameProvider:
			code = oauth2ErrorCodeProviderAlreadyConnected
		case sharederr.CodeUnauthenticated:
			code = oauth2ErrorCodeUnauthenticated
		}
	}

	if code == oauth2ErrorCodeFailed {
		log.Printf("failed to complete oauth2 flow: %+v", err)
	}

	return code
}
