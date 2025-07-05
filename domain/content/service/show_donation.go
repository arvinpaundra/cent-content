package service

import (
	"context"

	contentcmd "github.com/arvinpaundra/cent/content/application/command/content"
	contentres "github.com/arvinpaundra/cent/content/application/response/content"
	"github.com/arvinpaundra/cent/content/domain/content/repository"
)

type ShowDonation struct {
	contentReader repository.ContentReader
}

func NewShowDonation(
	contentReader repository.ContentReader,
) ShowDonation {
	return ShowDonation{
		contentReader: contentReader,
	}
}

func (s ShowDonation) Exec(ctx context.Context, command contentcmd.ShowDonation) (contentres.ShowDonation, error) {
	content, err := s.contentReader.FindActiveContentByUserId(ctx, command.UserId)
	if err != nil {
		return contentres.ShowDonation{}, err
	}

	message := content.Message

	res := contentres.ShowDonation{
		UserId:  command.UserId,
		UserKey: command.UserKey,
		Amount:  command.Amount,
		Sender:  command.Sender,
		Message: command.Message,
		MessageSetting: &contentres.MessageSetting{
			ID:          message.ID,
			BgColor:     message.BgColor,
			TextColor:   message.TextColor,
			IsTTSActive: message.IsTTSActive,
		},
	}

	return res, nil
}
