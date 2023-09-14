package logger

import (
	"os"
	"testing"

	"github.com/truebluekiwi/utility/slack"
)

func setup() {
	token := os.Getenv("SLACK_TOKEN")
	slack.Init("test", "local", "1.2.2", token)

	Init()
}

func TestInfo(t *testing.T) {
	setup()
	Info("test")
}

func TestError(t *testing.T) {
	setup()
	Error("test")
}
