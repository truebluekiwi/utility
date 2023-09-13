package slack

import (
	"fmt"

	"github.com/slack-go/slack"
)

var (
	svc          *slack.Client
	channel      = ""
	msg          = ""
	importantMsg = ""
)

func Configure(appName, appEnv, appVersion, token string) {
	svc = slack.New(token)
	channel = appName + "-monitoring"
	msg = fmt.Sprintf("StockAlerts *%s[%s]* `v%s`:\n", appName, appEnv, appVersion)
	importantMsg = fmt.Sprintf("StockAlerts *%s[%s]* `v%s`:\n<!channel> ", appName, appEnv, appVersion)
}

func Send(message string, isImportant bool) {
	m := msg + message
	if isImportant {
		m = importantMsg + message
	}
	svc.PostMessage(channel, slack.MsgOptionText(m, false))
}
