package internal

import "go.uber.org/zap"

var logger *zap.Logger

func Test() {
	logger.Info("This Test function")
}
