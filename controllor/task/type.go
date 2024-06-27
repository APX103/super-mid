package task

type WebPostTaskRequest struct {
	FeishuChatID       string `json:"chat_id,omitempty"`
	FeishuSenderID     string `json:"sender_id,omitempty"`
	FeishuSenderIDType string `json:"sender_id_type,omitempty"`
	UserName           string `json:"user_name,omitempty"`
	CommandLine        string `json:"command_line"`
}

type RunnerWebResponse struct {
	Code       string `json:"code"`
	Message    string `json:"message"`
	HasContent bool   `json:"has_content"`
	Content    string `json:"content,omitempty"`
}
