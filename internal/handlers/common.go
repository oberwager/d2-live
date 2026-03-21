package handlers

import (
	"log/slog"
	"net/http"
	"strings"
	"time"
)

type Controller struct {
	Logger  *slog.Logger
	Version string
}

type responseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriter) WriteHeader(status int) {
	rw.status = status
	rw.ResponseWriter.WriteHeader(status)
}

func (c *Controller) LoggingMiddleware(f http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		start := time.Now()
		lrw := &responseWriter{ResponseWriter: rw, status: http.StatusOK}
		f(lrw, req)

		path := strings.Split(req.URL.Path, "/")[1]
		if path == "" {
			path = "svg"
		}

		c.Logger.Info("request",
			"endpoint", path,
			"duration_ms", time.Since(start).Milliseconds(),
			"status", lrw.status,
		)
	}
}
