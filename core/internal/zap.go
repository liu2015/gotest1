package internal

import (
	"ginserver/global"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Zap = new(_zap)

type _zap struct{}

func (z *_zap) GetEncoder() zapcore.Encoder {
	if global.GVA_CONFIG.Zap.Format == "json" {
		return zapcore.NewJSONEncoder(z.GetEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(z.GetEncoderConfig())
}

func (z *_zap) GetEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{

		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  global.GVA_CONFIG.Zap.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    global.GVA_CONFIG.Zap.ZapEncodeLevel(),
		EncodeTime:     z.CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
}

func (z *_zap) GetEncoderCore(l zapcore.Level, level zap.LevelEnablerFunc) zapcore.Core {
	writer := FileRotatelogs.GetWriteSyncer(l.String())
	return zapcore.NewCore(z.GetEncoder(), writer, level)
}

func (z *_zap) CustomTimeEncoder(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
	encoder.AppendString(global.GVA_CONFIG.Zap.Prefix + t.Format("2006/01/02 - 15:04:05.000"))
}

func (z *_zap) GetZapCores() []zapcore.Core {
	cores := make([]zapcore.Core, 0, 7)
	for level := global.GVA_CONFIG.Zap.TransportLevel(); level <= zapcore.FatalLevel; level++ {
		cores = append(cores, z.GetEncoderCore(level, z.GetLevelPriority(level)))
	}
	return cores

}
func (z *_zap) GetLevelPriority(level zapcore.Level) zap.LevelEnablerFunc {
	switch level {
	case zapcore.DebugLevel:
		return func(l zapcore.Level) bool {
			return level == zap.DebugLevel
		}
	case zapcore.InfoLevel:
		return func(l zapcore.Level) bool {
			return level == zap.InfoLevel
		}
	case zapcore.WarnLevel:
		return func(l zapcore.Level) bool {
			return level == zap.WarnLevel
		}
	case zapcore.ErrorLevel:
		return func(l zapcore.Level) bool {
			return level == zap.ErrorLevel
		}
	case zapcore.DPanicLevel:
		return func(l zapcore.Level) bool {
			return level == zap.DPanicLevel
		}
	case zap.PanicLevel:
		return func(l zapcore.Level) bool {
			return level == zap.PanicLevel
		}
	case zapcore.FatalLevel:
		return func(l zapcore.Level) bool {
			return level == zap.FatalLevel
		}
	default:
		return func(l zapcore.Level) bool {
			return level == zap.DebugLevel
		}

	}
}
