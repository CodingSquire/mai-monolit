package httplogger

import (
	"errors"
	"gitlab.com/CodingSquire/mai_monolit/pkg/logger"
	"net/http"
	"runtime/debug"


)

var err = errors.New(http.StatusText(http.StatusInternalServerError))

func Recoverer(internalServerError func(w http.ResponseWriter, r *http.Request, err error)) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if info := recover(); info != nil {
					logger.Ctx(r.Context()).Error().Interface("recover_info", info).Bytes("debug_stack", debug.Stack()).Msg("panic_on_request")
					internalServerError(w, r, err)
				}
			}()
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
