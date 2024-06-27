package feishu

// Feishu msg temp in MongoDB
type FeishuMssageTemplate struct {
	TaskCmdPath  string   `json:"task_cmd_path" bson:"task_cmd_path"`
	TaskStep     string   `json:"task_step" bson:"task_step"`
	TemplateKeys []string `json:"template_keys" bson:"template_keys"`
	Template     string   `json:"template" bson:"template"`
}

// mr meeseeks feishu message struct: Header
type Header struct {
	EventID    string `json:"event_id,omitempty"`
	EventType  string `json:"event_type"`
	CreateTime string `json:"create_time,omitempty"`
	Token      string `json:"token,omitempty"`
	AppID      string `json:"app_id,omitempty"`
	TenantKey  string `json:"tenant_key,omitempty"`
}

// mr meeseeks feishu message struct: SendMessageEvent
type SendMessageEvent struct {
	Message Message `json:"message"`
}

// mr meeseeks feishu message struct: Message
type Message struct {
	MessageID   string     `json:"message_id,omitempty"`
	RootID      string     `json:"root_id,omitempty"`
	ParentID    string     `json:"parent_id,omitempty"`
	CreateTime  string     `json:"create_time,omitempty"`
	ChatID      string     `json:"chat_id,omitempty"`
	ChatType    string     `json:"chat_type,omitempty"`
	MessageType string     `json:"message_type"`
	Content     string     `json:"content"`
	Mentions    []*Mention `json:"mentions,omitempty"`
}

// mr meeseeks feishu message struct: Mention
type Mention struct {
	Key       string  `json:"key,omitempty"`
	ID        *UserID `json:"id,omitempty"`
	Name      string  `json:"name,omitempty"`
	TenantKey string  `json:"tenant_key,omitempty"`
}

// mr meeseeks feishu message struct: UserID
type UserID struct {
	UserID  string `json:"user_id,omitempty"`
	OpenID  string `json:"open_id,omitempty"`
	UnionID string `json:"union_id,omitempty"`
}

/*
Message card, May not use.
*/
type CardContent struct {
	Config   *CardConfig   `json:"config,omitempty"`
	Header   *CardHeader   `json:"header,omitempty"`
	Elements []interface{} `json:"elements,omitempty"`
}

type CardConfig struct {
	WideScreenMode bool `json:"wide_screen_mode,omitempty"`
	EnableForward  bool `json:"enable_forward,omitempty"`
}

type CardHeader struct {
	Title    *CardText `json:"title,omitempty"`
	Template string    `json:"template,omitempty"`
}

type CardElement struct {
	Tag          string       `json:"tag"`
	Text         *CardText    `json:"text,omitempty"`
	Fields       []*CardField `json:"fields,omitempty"`
	ImgKey       string       `json:"img_key,omitempty"`
	Mode         string       `json:"mode,omitempty"`
	Alt          *CardText    `json:"alt,omitempty"`
	CustomWidth  int          `json:"custom_width,omitempty"`
	CompactWidth int          `json:"compact_width,omitempty"`
	Preview      bool         `json:"preview,omitempty"`
	Title        *CardText    `json:"title,omitempty"`
}

type CardNote struct {
	Tag      string        `json:"tag,omitempty"`
	Elements []interface{} `json:"elements,omitempty"`
}

type CardActionBlock struct {
	Tag     string        `json:"tag,omitempty"`
	Layout  string        `json:"layout,omitempty"`
	Actions []interface{} `json:"actions,omitempty"`
}

type CardText struct {
	Tag     string `json:"tag,omitempty"`
	Content string `json:"content,omitempty"`
	Lines   int    `json:"lines,omitempty"`
}

type CardField struct {
	IsShort bool      `json:"is_short,omitempty"`
	Text    *CardText `json:"text,omitempty"`
}

type CardButton struct {
	Tag     string            `json:"tag,omitempty"`
	Text    *CardText         `json:"text,omitempty"`
	Type    string            `json:"type,omitempty"`
	Value   map[string]string `json:"value,omitempty"`
	Confirm *CardConfirm      `json:"confirm,omitempty"`
}

type CardConfirm struct {
	Title *CardText `json:"title,omitempty"`
	Text  *CardText `json:"text,omitempty"`
}

type CardSelectMenu struct {
	Tag           string            `json:"tag"`
	PlaceHolder   *CardText         `json:"placeholder,omitempty"`
	InitialOption string            `json:"initial_option,omitempty"`
	Options       []*CardOption     `json:"options,omitempty"`
	Value         map[string]string `json:"value,omitempty"`
	Confirm       []*CardConfirm    `json:"confirm,omitempty"`
}

type CardOption struct {
	Text     *CardText `json:"text,omitempty"`
	Value    string    `json:"value"`
	URL      string    `json:"url,omitempty"`
	MultiURL *CardURL  `json:"multi_url,omitempty"`
}

type CardURL struct {
	URL        string `json:"url"`
	AndroidURL string `json:"android_url"`
	IosURL     string `json:"ios_url"`
	PcURL      string `json:"pc_url"`
}

type CardSplitLine struct {
	Tag string `json:"tag"`
}

/*
Message card end
*/

type FeishuRequestBody struct {
	Header    Header           `json:"header"`
	Event     SendMessageEvent `json:"event"`
	GroupName string           `json:"groupname,omitempty"`
}
