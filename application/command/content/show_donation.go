package content

type ShowDonation struct {
	UserId  int64   `json:"user_id"`
	UserKey string  `json:"user_key"`
	Amount  float64 `json:"amount"`
	Sender  string  `json:"sender"`
	Message string  `json:"message"`
}
