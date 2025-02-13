package logging

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Level zapcore.Level

const (
	InfoLevel      = Level(zap.InfoLevel) // default level
	WarningLevel   = Level(zap.WarnLevel)
	ErrorLevel     = Level(zap.ErrorLevel)
	CriticalLevel  = Level(zap.DPanicLevel) // used in development log
	AlertLevel     = Level(zap.PanicLevel)  // logs a message, then panics
	EmergencyLevel = Level(zap.FatalLevel)  // logs a message, then calls os.Exit(1).
	DebugLevel     = Level(zap.DebugLevel)
)

func (l Level) String() string {
	switch l {
	case InfoLevel:
		return "info"
	case WarningLevel:
		return "warning"
	case ErrorLevel:
		return "error"
	case CriticalLevel:
		return "critical"
	case AlertLevel:
		return "alert"
	case EmergencyLevel:
		return "emergency"
	case DebugLevel:
		return "debug"
	default:
		return "unknown"
	}
}

func ParserLevel(level string) (Level, error) {
	switch level {
	case "info":
		return InfoLevel, nil
	case "warning":
		return WarningLevel, nil
	case "error":
		return ErrorLevel, nil
	case "critical":
		return CriticalLevel, nil
	case "alert":
		return AlertLevel, nil
	case "emergency":
		return EmergencyLevel, nil
	case "debug":
		return DebugLevel, nil
	}

	return 0, fmt.Errorf("unknown level: %s", level)
}
