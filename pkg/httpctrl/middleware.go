package httpctrl

import (
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
)

func SetResponseHeader(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if reqID := middleware.GetReqID(r.Context()); reqID != "" {
			w.Header().Add("reqID", reqID)
		}

		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
