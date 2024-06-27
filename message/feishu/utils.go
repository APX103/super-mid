package feishu

import (
	"encoding/json"

	"github.com/sirupsen/logrus"
)

func BuildSimpleFeishuCardWithText(title string, txt string) string {
	cardContent := &CardContent{
		Header: &CardHeader{
			Title: &CardText{
				Tag:     "plain_text",
				Content: title,
			},
		},
		Config: &CardConfig{
			WideScreenMode: true,
			EnableForward:  true,
		},
		Elements: []interface{}{
			CardText{
				Tag:     "markdown",
				Content: txt,
			},
		},
	}

	content, err := json.Marshal(cardContent)
	if err != nil {
		logrus.Error("marshal card content error.")
	}
	return string(content)
}
