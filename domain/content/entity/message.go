package entity

type Message struct {
	ID          int64
	ContentId   int64
	BgColor     string
	TextColor   string
	IsTTSActive bool
}

func (e *Message) IsEmpty() bool {
	return e == nil
}
