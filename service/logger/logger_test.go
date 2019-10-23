package logger

import "testing"

func TestLogger(t *testing.T) {
	GetLogger().Info("Just Test Logger")
}
