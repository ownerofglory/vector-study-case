package logging

import (
	"github.com/golang/mock/gomock"
	"testing"
)

func TestDebug(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockLogger := NewMockLogger(ctrl)
	currentLogger = mockLogger

	mockLogger.EXPECT().Debug(gomock.Any())
	Debug("args")
}

func TestDebugf(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockLogger := NewMockLogger(ctrl)
	currentLogger = mockLogger

	mockLogger.EXPECT().Debugf(gomock.Any())
	Debugf("args")
}

func TestInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockLogger := NewMockLogger(ctrl)
	currentLogger = mockLogger

	mockLogger.EXPECT().Info(gomock.Any())
	Info("args")
}

func TestInfof(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockLogger := NewMockLogger(ctrl)
	currentLogger = mockLogger

	mockLogger.EXPECT().Infof(gomock.Any())
	Infof("args")
}

func TestWarn(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockLogger := NewMockLogger(ctrl)
	currentLogger = mockLogger

	mockLogger.EXPECT().Warn(gomock.Any())
	Warn("args")
}

func TestWarnf(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockLogger := NewMockLogger(ctrl)
	currentLogger = mockLogger

	mockLogger.EXPECT().Warnf(gomock.Any())
	Warnf("args")
}

func TestError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockLogger := NewMockLogger(ctrl)
	currentLogger = mockLogger

	mockLogger.EXPECT().Error(gomock.Any())
	Error("args")
}

func TestErrorf(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockLogger := NewMockLogger(ctrl)
	currentLogger = mockLogger

	mockLogger.EXPECT().Errorf(gomock.Any())
	Errorf("args")
}
