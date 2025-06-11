package content

import (
	"context"
	"errors"

	"github.com/arvinpaundra/cent/content/domain/content/entity"
	"github.com/arvinpaundra/cent/content/domain/content/repository"
	"github.com/arvinpaundra/cent/content/model"
	"github.com/guregu/null/v6"
	"gorm.io/gorm"
)

var _ repository.ContentWriter = (*ContentWriterRepository)(nil)

type ContentWriterRepository struct {
	db *gorm.DB
}

func NewContentWriterRepository(db *gorm.DB) ContentWriterRepository {
	return ContentWriterRepository{db: db}
}

func (r ContentWriterRepository) Save(ctx context.Context, content *entity.Content) error {
	if content.IsNew() {
		return r.insert(ctx, content)
	}

	return errors.New("unsupported database operation")
}

func (r ContentWriterRepository) insert(ctx context.Context, content *entity.Content) error {
	contentModel := model.Content{
		UserId:      content.UserId,
		RingtoneUrl: content.RingtoneUrl,
	}

	err := r.db.WithContext(ctx).
		Model(&model.Content{}).
		Create(&contentModel).
		Error

	if err != nil {
		return err
	}

	content.ID = contentModel.ID

	if !content.QrCode.IsEmpty() {
		qrcode := content.QrCode

		qrcodeModel := model.QrCode{
			ContentId: contentModel.ID,
			Code:      qrcode.Code,
			BgColor:   qrcode.BgColor,
			QrColor:   qrcode.QrColor,
			Text:      null.StringFromPtr(qrcode.Text),
		}

		err := r.db.WithContext(ctx).
			Model(&model.QrCode{}).
			Create(&qrcodeModel).
			Error

		if err != nil {
			return err
		}

		qrcode.ID = qrcodeModel.ID
	}

	if !content.Campaign.IsEmpty() {
	}

	return nil
}
