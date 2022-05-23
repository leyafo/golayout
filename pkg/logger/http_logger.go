package logger

import (
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"

	// _ "go.uber.org/zap/zapcore"
	"fmt"
	"net/http"
	"time"
)

type StructuredLogger struct {
	logger *zap.Logger
}

func NewStructuredLogger() func(next http.Handler) http.Handler {
	return middleware.RequestLogger(&StructuredLogger{zapLogger})
}

func (l *StructuredLogger) NewLogEntry(r *http.Request) middleware.LogEntry {
	var logFields []zap.Field

	// logFields = append(logFields, zap.Time("ts", time.Now().UTC()))
	var reqID string
	if reqID = middleware.GetReqID(r.Context()); reqID != "" {
		logFields = append(logFields, zap.String("req_id", reqID))
	}

	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	logFields = append(logFields, zap.String("http_scheme", scheme))
	logFields = append(logFields, zap.String("http_proto", r.Proto))
	logFields = append(logFields, zap.String("http_method", r.Method))
	logFields = append(logFields, zap.String("remote_addr", r.RemoteAddr))

	logFields = append(logFields, zap.String("user_agent", r.UserAgent()))

	logFields = append(logFields, zap.String("uri", fmt.Sprintf("%s://%s%s", scheme, r.Host, r.RequestURI)))

	entry := &StructuredLoggerEntry{Logger: l.logger, ReqID: reqID}
	l.logger.Info("request started", logFields...)

	return entry
}

type StructuredLoggerEntry struct {
	Logger *zap.Logger
	ReqID  string
}

func (l *StructuredLoggerEntry) Write(status, bytes int,
	header http.Header, elapsed time.Duration, extra interface{}) {

	l.Logger.Info("request complete",
		zap.String("req_id", l.ReqID),
		zap.Int("resp_status", status),
		zap.Int("resp_bytes_length", bytes),
		zap.Float64("resp_elapsed_ms", float64(elapsed.Nanoseconds())/1000000.0))
}

func (l *StructuredLoggerEntry) Panic(v interface{}, stack []byte) {
	l.Logger.Error("panic",
		zap.String("req_id", l.ReqID),
		zap.ByteString("stack", stack),
		zap.String("panic", fmt.Sprintf("%+v", v)),
	)
}
