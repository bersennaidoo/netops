package mid

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/kataras/golog"
)

type Middleware struct {
	log *golog.Logger
}

func New(log *golog.Logger) *Middleware {
	return &Middleware{
		log: log,
	}
}

func (m *Middleware) LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m.log.Infof("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI())

		next.ServeHTTP(w, r)
	})
}

func (m *Middleware) RecoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connection", "close")
				trace := fmt.Sprintf("%s\n%s", fmt.Errorf("%s", err), debug.Stack())
				m.log.Error(trace)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
