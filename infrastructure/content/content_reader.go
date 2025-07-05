package content

import (
	"context"

	"github.com/arvinpaundra/cent/content/domain/content/entity"
	"github.com/arvinpaundra/cent/content/model"
	"gorm.io/gorm"
)

type ContentReaderRepository struct {
	db *gorm.DB
}

func NewContentReaderRepository(db *gorm.DB) ContentReaderRepository {
	return ContentReaderRepository{db: db}
}

func (r ContentReaderRepository) FindActiveContentByUserId(ctx context.Context, userId int64) (*entity.Content, error) {
	var contentModel model.Content

	err := r.db.WithContext(ctx).
		Model(&model.Content{}).
		Preload("Campaign", "is_active = true").
		Preload("Message").
		Where("user_id = ? AND deleted_at IS NULL", userId).
		First(&contentModel).
		Error

	if err != nil {
		return nil, err
	}

	content := entity.Content{
		ID:          contentModel.ID,
		UserId:      userId,
		RingtoneUrl: contentModel.RingtoneUrl,
	}

	if contentModel.Campaign != nil {
		campaignModel := contentModel.Campaign

		campaign := &entity.Campaign{
			ID:            campaignModel.ID,
			ContentId:     campaignModel.ContentId,
			TargetAmount:  campaignModel.TargetAmount,
			CurrentAmount: campaignModel.CurrentAmount,
			Text:          campaignModel.Text,
			IsActive:      campaignModel.IsActive,
		}

		content.SetCampaign(campaign)
	}

	if contentModel.Message != nil {
		messageModel := contentModel.Message

		message := &entity.Message{
			ID:          messageModel.ID,
			ContentId:   messageModel.ContentId,
			BgColor:     messageModel.BgColor,
			TextColor:   messageModel.TextColor,
			IsTTSActive: messageModel.IsTtsActive,
		}

		content.SetMessage(message)
	}

	return &content, nil
}
