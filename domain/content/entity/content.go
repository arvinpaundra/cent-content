package entity

type Content struct {
	ID          int64
	UserId      int64
	RingtoneUrl string
	QrCode      *QrCode
	Message     *Message
	Campaign    *Campaign
}

func (e *Content) IsNew() bool {
	return e.ID <= 0
}

func (e *Content) SetQrCode(qrcode *QrCode) {
	e.QrCode = qrcode
}

func (e *Content) SetMessage(message *Message) {
	e.Message = message
}

func (e *Content) SetCampaign(campaign *Campaign) {
	e.Campaign = campaign
}
