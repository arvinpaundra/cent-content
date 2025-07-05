package entity

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
