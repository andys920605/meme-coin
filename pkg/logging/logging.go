package logging

import (
	"context"
	"os"
	"sort"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/andys920605/meme-coin/pkg/trace"
)

type Logging struct {
	logger *zap.SugaredLogger
}

type Fields map[string]any

func New(options ...Option) *Logging {
	return newLogging(1, options...)
}

func newLogging(callerSkipCount int, options ...Option) *Logging {
	o := []Option{
		WithServiceName("backend-service"),
		WithLevel(DebugLevel),
	}

	o = append(o, options...)
	settings := processOpts(o)

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:       "time",
		LevelKey:      "severity",
		NameKey:       "logger",
		CallerKey:     "caller",
		MessageKey:    "message",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   encodeLevel(),
		EncodeTime:    zapcore.RFC3339TimeEncoder,
		EncodeCaller:  zapcore.ShortCallerEncoder,
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.AddSync(os.Stdout),
		zapcore.Level(settings.Level),
	)

	zapOption := make([]zap.Option, 0)
	zapOption = append(zapOption, zap.AddCallerSkip(callerSkipCount))
	if settings.IsShowCaller {
		zapOption = append(zapOption, zap.AddCaller())
	}

	logger := zap.New(core, zapOption...).Sugar()
	if settings.ServiceName != "" {
		logger = logger.With("service_name", settings.ServiceName)
	}

	return &Logging{
		logger: logger,
	}
}

func encodeLevel() zapcore.LevelEncoder {
	return func(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
		var content string
		switch l {
		case zapcore.DebugLevel:
			content = "DEBUG"
		case zapcore.InfoLevel:
			content = "INFO"
		case zapcore.WarnLevel:
			content = "WARNING"
		case zapcore.ErrorLevel:
			content = "ERROR"
		case zapcore.DPanicLevel:
			content = "CRITICAL"
		case zapcore.PanicLevel:
			content = "ALERT"
		case zapcore.FatalLevel:
			content = "EMERGENCY"
		}
		enc.AppendString(content)
	}
}

func (l *Logging) Sync() {
	_ = l.logger.Sync()
}

func (l *Logging) WithFields(field Fields) *Logging {
	keys := make([]string, 0, len(field))
	for k := range field {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	fields := make([]any, 0, len(field)*2)
	for _, key := range keys {
		fields = append(fields, key, field[key])
	}

	return &Logging{
		logger: l.logger.With(fields...),
	}
}

func (l *Logging) WithTraceID(ctx context.Context) *Logging {
	return l.WithFields(Fields{
		"trace_id": trace.GetTraceIDFromContext(ctx),
	})
}

func (l *Logging) Debug(msg string) {
	l.logger.Debug(msg)
}

func (l *Logging) Debugf(format string, args ...any) {
	l.logger.Debugf(format, args...)
}

func (l *Logging) Info(msg string) {
	l.logger.Info(msg)
}

func (l *Logging) Infof(format string, args ...any) {
	l.logger.Infof(format, args...)
}

func (l *Logging) Warning(msg string) {
	l.logger.Warn(msg)
}

func (l *Logging) Warningf(format string, args ...any) {
	l.logger.Warnf(format, args...)
}

func (l *Logging) Error(msg string) {
	l.logger.Error(msg)
}

func (l *Logging) Errorf(format string, args ...any) {
	l.logger.Errorf(format, args...)
}

// Critical logs are particularly important errors. In development the logger panics after writing the message.
func (l *Logging) Critical(msg string) {
	l.logger.DPanic(msg)
}

func (l *Logging) Criticalf(format string, args ...any) {
	l.logger.DPanicf(format, args...)
}

// Alert logs a message, then panics
func (l *Logging) Alert(msg string) {
	l.logger.Panic(msg)
}

func (l *Logging) Alertf(format string, args ...any) {
	l.logger.Panicf(format, args...)
}

// Emergency logs a message, then calls os.Exit(1).
func (l *Logging) Emergency(msg string) {
	l.logger.Fatal(msg)
}

func (l *Logging) Emergencyf(format string, args ...any) {
	l.logger.Fatalf(format, args...)
}
