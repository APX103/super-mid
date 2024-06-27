package feishu

import (
	"context"
	"net/http"

	"apx103.com/super-mid/utils/config"
)

type FeishuMessageSender struct {
	Url    string
	Ctx    context.Context
	Client *http.Client
}

func NewFeishuMessageSender(conf config.BaseConfig) *FeishuMessageSender {
	ctx := context.Background()
	client := &http.Client{}
	return &FeishuMessageSender{
		Url:    conf.MrMeeseeksUrl,
		Ctx:    ctx,
		Client: client,
	}
}

func (fms *FeishuMessageSender) SendMessage() {

}
