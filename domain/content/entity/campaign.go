package entity

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
