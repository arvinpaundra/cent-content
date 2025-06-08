package entity

type Content struct {
	ID          int64
	UserId      int64
	RingtoneUrl string
	QrCode      *QrCode
}

func (e *Content) SetQrCode(qrcode *QrCode) {
	e.QrCode = qrcode
}

func (e *Content) IsQrCodeEmpty() bool {
	return e.QrCode == nil
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
