package slack

import (
	"testing"
)

func init() {
	Configure("feed-retriever", "test", "1.0.0", "xoxb-5406250867265-5393115934967-03fSjM3RBgiGH5zAodz7DYqM")
}

func TestSendSlackMonitoringMessage(t *testing.T) {
	Send("please ignore, test message", false)
	Send("please ignore, test important message", true)
}
