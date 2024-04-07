package calendarv1beta

import (
	"net/http"
)

var _ http.Handler = (*impl)(nil)

func (h *impl) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

type impl struct{}

func New() *impl {
	return &impl{}
}
