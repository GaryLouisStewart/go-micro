package middleware

import (
	"log/slog"
	"net/http"
	"os"
	"time"
)


var Logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))


func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

	wrappedWriter := &statusRecorder{ResponseWriter: w, StatusCode: http.StatusOK}
	
	next.ServeHTTP(wrappedWriter, r)

	duration := time.Since(start)

	Logger.Info("request handled",
			slog.String("method", r.Method),
			slog.String("path", r.URL.Path),
			slog.Int("status", wrappedWriter.StatusCode),
			slog.String("duration", duration.String()),
			slog.String("user_agent", r.UserAgent()),
	)
 })
}

type statusRecorder struct {
	http.ResponseWriter
	StatusCode int
}


func (rec *statusRecorder) Writeheader(code int) {
	rec.StatusCode = code
	rec.ResponseWriter.WriteHeader(code)
}

