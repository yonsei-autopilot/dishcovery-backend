package middleware

import (
	"bytes"
	"io"
	"log/slog"
	"net/http"
	"time"

	"github.com/yonsei-autopilot/smart-menu-backend/internal/common/util"
)

type wrappedWriter struct {
	http.ResponseWriter
	statusCode int
	body       bytes.Buffer
}

func (w *wrappedWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

func (w *wrappedWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := util.GetKstNow()
		reqBodyString := getReqeustBodyString(r)
		wrapped := &wrappedWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(wrapped, r)

		logApi(r, start, reqBodyString, *wrapped)
	})
}

func getReqeustBodyString(r *http.Request) string {
	reqBodyBytes, _ := io.ReadAll(r.Body)
	reqBodyString := string(reqBodyBytes)
	r.Body = io.NopCloser(bytes.NewReader(reqBodyBytes))
	return reqBodyString
}

func logApi(r *http.Request, start time.Time, reqBodyString string, wrapped wrappedWriter) {
	requestGroup := slog.Group(
		"request",
		slog.String("method", r.Method),
		slog.String("path", r.URL.Path),
		slog.String("body", reqBodyString),
		slog.String("accessToken", r.Header.Get("Authorization")),
	)
	responseGroup := slog.Group(
		"response",
		slog.Int("status", wrapped.statusCode),
		slog.String("duration", time.Since(start).String()),
		slog.String("body", wrapped.body.String()),
	)
	slog.Info("req and res",
		requestGroup,
		responseGroup,
	)
}
