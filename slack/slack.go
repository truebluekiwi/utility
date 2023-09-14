package slack

import (
	"fmt"
	"log"

	"github.com/slack-go/slack"
)

var (
	svc          *slack.Client
	channel      = ""
	msg          = ""
	importantMsg = ""
)

type LogLevel int

const (
	LogLevelWarning LogLevel = iota + 1
	LogLevelError
	LogLevelLog
)

type Writer struct {
	LogLevel LogLevel
}

func (sw *Writer) Write(b []byte) (n int, err error) {
	switch sw.LogLevel {
	case LogLevelWarning:
		SendWarning(string(b))
	case LogLevelError:
		SendError(string(b))
	case LogLevelLog:
		SendLog(string(b))
	}
	return len(b), nil
}

func (sw *Writer) Sync() error {
	return nil
}

func Init(appName, appEnv, appVersion, token string) {
	svc = slack.New(token)
	channel = appName + "-monitoring"
	msg = fmt.Sprintf("StockAlerts *%s[%s]* `v%s`:\n", appName, appEnv, appVersion)
	importantMsg = fmt.Sprintf("StockAlerts *%s[%s]* `v%s`:\n<!channel> ", appName, appEnv, appVersion)
}

func SendWarning(message string) {
	_, _, err := svc.PostMessage(channel, slack.MsgOptionText(importantMsg+message, false))
	if err != nil {
		log.Printf("Error: failed to SendWarning: %+v", err)
	}
}

func SendError(message string) {
	if _, _, err := svc.PostMessage(channel, slack.MsgOptionText(importantMsg+message, false)); err != nil {
		log.Printf("Error: failed to SendError: %+v", err)
	}
}

func SendLog(message string) {
	if _, _, err := svc.PostMessage("logs-monitoring", slack.MsgOptionText(msg+message, false)); err != nil {
		log.Printf("Error: failed to SendLog: %+v", err)
	}
}
