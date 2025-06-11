package entity

type Content struct {
	ID          int64
	UserId      int64
	RingtoneUrl string
	QrCode      *QrCode
	Campaign    *Campaign
}

func (e *Content) IsNew() bool {
	return e.ID <= 0
}

func (e *Content) SetQrCode(qrcode *QrCode) {
	e.QrCode = qrcode
}

func (e *Content) SetCampaign(campaign *Campaign) {
	e.Campaign = campaign
}

type QrCode struct {
	ID      int64
	Code    string
	BgColor string
	QrColor string
	Text    *string
}

func (e *QrCode) IsNew() bool {
	return e.ID <= 0
}

func (e *QrCode) IsEmpty() bool {
	return e == nil
}

type Campaign struct {
	ID            int64
	ContentId     int64
	TargetAmount  float64
	CurrentAmount float64
	Text          string
	IsActive      bool
}

func (e *Campaign) IsEmpty() bool {
	return e == nil
}
