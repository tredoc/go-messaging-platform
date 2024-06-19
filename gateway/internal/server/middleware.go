package server

import (
	"log/slog"
	"net/http"
	"time"
)

type ResponseRecorder struct {
	http.ResponseWriter
	StatusCode int
	Body       []byte
}

func (rec *ResponseRecorder) WriteHeader(statusCode int) {
	rec.StatusCode = statusCode
	rec.ResponseWriter.WriteHeader(statusCode)
}

func (rec *ResponseRecorder) Write(body []byte) (int, error) {
	rec.Body = body
	return rec.ResponseWriter.Write(body)
}

func HttpLogger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		startTime := time.Now()
		rec := &ResponseRecorder{
			ResponseWriter: res,
			StatusCode:     http.StatusOK,
		}

		handler.ServeHTTP(rec, req)
		duration := time.Since(startTime)

		logger := slog.Info
		if rec.StatusCode != http.StatusOK {
			logger = slog.With(slog.Any("res_body", string(rec.Body))).Error
		}

		logger("HTTP call",
			slog.String("protocol", "http"),
			slog.String("method", req.Method),
			slog.String("path", req.RequestURI),
			slog.Int("status_code", rec.StatusCode),
			slog.String("status_text", http.StatusText(rec.StatusCode)),
			slog.Duration("duration", duration))
	})
}
