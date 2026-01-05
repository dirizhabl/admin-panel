package middleware

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

func Logger(logger *logrus.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			rw := &responseWriter{ResponseWriter: w}
			next.ServeHTTP(rw, r)

			var level logrus.Level
			switch {
			case rw.code >= 500:
				level = logrus.ErrorLevel
			case rw.code >= 400:
				level = logrus.WarnLevel
			default:
				level = logrus.InfoLevel
			}
			logger.Logf(
				level,
				`%s - "%s %s %s" %d %s %s`,
				r.RemoteAddr,
				r.Method,
				r.RequestURI,
				r.Proto,
				rw.code,
				http.StatusText(rw.code),
				time.Since(start),
			)
		})
	}
}
