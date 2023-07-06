package utils

import (
	"go.uber.org/zap"
)

func Error(err, msg string) {
	zap.L().Error(msg, zap.Field{})
}
