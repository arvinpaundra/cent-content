package service

import (
	"context"

	"github.com/arvinpaundra/cent/content/application/command/content"
	"github.com/arvinpaundra/cent/content/domain/content/entity"
	"github.com/arvinpaundra/cent/content/domain/content/repository"
)

type SetupContentHandler struct {
	contentWriter repository.ContentWriter
	unitOfWork    repository.UnitOfWork
}

func NewSetupContentHandler(
	contentWriter repository.ContentWriter,
	unitOfWork repository.UnitOfWork,
) SetupContentHandler {
	return SetupContentHandler{
		contentWriter: contentWriter,
		unitOfWork:    unitOfWork,
	}
}

func (s SetupContentHandler) Handle(ctx context.Context, command content.CreateSetupContent) error {
	tx, err := s.unitOfWork.Begin()
	if err != nil {
		return err
	}

	content := entity.Content{
		UserId:      command.UserId,
		RingtoneUrl: "",
	}

	// setup qrcode metadata
	qrcode := entity.QrCode{
		Code:    command.UserSlug,
		BgColor: "",
		QrColor: "",
	}

	content.SetQrCode(&qrcode)

	err = tx.ContentWriter().Save(ctx, &content)
	if err != nil {
		if uowErr := tx.Rollback(); uowErr != nil {
			return uowErr
		}

		return err
	}

	if uowErr := tx.Commit(); uowErr != nil {
		return uowErr
	}

	return nil
}
