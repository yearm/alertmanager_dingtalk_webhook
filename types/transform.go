package types

import (
	"bytes"
	"fmt"
)

func AlertToDingTalk(alertMsg AlertMessage) *MarkdownMessage {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("**version**: %s\n\n", alertMsg.Version))
	buffer.WriteString(fmt.Sprintf("**groupKey**: %s\n\n", alertMsg.GroupKey))
	buffer.WriteString(fmt.Sprintf("**receiver**: %s\n\n", alertMsg.Receiver))
	buffer.WriteString(fmt.Sprintf("**status**: %s\n\n", alertMsg.Status))
	buffer.WriteString(fmt.Sprintf("**externalUrl**: [%s](%s)\n\n", alertMsg.ExternalURL, alertMsg.ExternalURL))
	buffer.WriteString("### alerts:\n")
	for _, alert := range alertMsg.Alerts {
		buffer.WriteString("----------------------\n\n")
		buffer.WriteString("#### labels:\n")
		buffer.WriteString(fmt.Sprintf("- **status**: %s\n", alert.Labels.Status))
		buffer.WriteString(fmt.Sprintf("- **alertname**: %s\n", alert.Labels.Alertname))
		buffer.WriteString(fmt.Sprintf("- **group**: %s\n", alert.Labels.Group))
		buffer.WriteString(fmt.Sprintf("- **instance**: %s\n", alert.Labels.Instance))
		buffer.WriteString(fmt.Sprintf("- **job**: %s\n", alert.Labels.Job))
		buffer.WriteString("#### annotations:\n")
		buffer.WriteString(fmt.Sprintf("- **description**: %s\n", alert.Annotations.Description))
		buffer.WriteString(fmt.Sprintf("- **summary**: %s\n\n", alert.Annotations.Summary))
		buffer.WriteString(fmt.Sprintf("**startsAt**: %s\n\n", alert.StartsAt.Format("2006-01-02 15:04:05")))
		buffer.WriteString(fmt.Sprintf("**endsAt**: %s\n\n", alert.EndsAt.Format("2006-01-02 15:04:05")))
		buffer.WriteString(fmt.Sprintf("**generatorUrl**: [%s](%s)\n\n", alert.GeneratorURL, alert.GeneratorURL))
		buffer.WriteString("----------------------\n\n")
	}

	return &MarkdownMessage{
		MsgType: "markdown",
		Markdown: &Markdown{
			Title: "Prometheus告警信息",
			Text:  buffer.String(),
		},
		At: &At{
			AtMobiles: []string{},
			IsAtAll:   false,
		},
	}
}
