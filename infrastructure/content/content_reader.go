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

	if content.Campaign != nil {
		campaign := &entity.Campaign{
			ID:            content.Campaign.ID,
			ContentId:     content.Campaign.ContentId,
			TargetAmount:  content.Campaign.TargetAmount,
			CurrentAmount: content.Campaign.CurrentAmount,
			Text:          content.Campaign.Text,
			IsActive:      content.Campaign.IsActive,
		}

		content.SetCampaign(campaign)
	}

	return &content, nil
}
