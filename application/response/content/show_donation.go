package content

type ShowDonation struct {
	UserId  int64   `json:"user_id"`
	UserKey string  `json:"user_key"`
	Amount  float64 `json:"amount"`
	Sender  string  `json:"sender"`
	Message string  `json:"message"`

	MessageSetting *MessageSetting `json:"message_setting"`
}

type MessageSetting struct {
	ID          int64  `json:"id"`
	BgColor     string `json:"bg_color"`
	TextColor   string `json:"text_color"`
	IsTTSActive bool   `json:"is_tts_active"`
}
