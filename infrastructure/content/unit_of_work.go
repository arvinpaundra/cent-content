package content

import (
	"github.com/arvinpaundra/cent/content/domain/content/repository"
	"gorm.io/gorm"
)

var _ repository.UnitOfWork = (*UnitOfWork)(nil)
var _ repository.UnitOfWorkProcessor = (*UnitOfWorkProcessor)(nil)

type UnitOfWork struct {
	db *gorm.DB
}

func NewUnitOfWork(db *gorm.DB) repository.UnitOfWork {
	return UnitOfWork{db: db}
}

func (r UnitOfWork) Begin() (repository.UnitOfWorkProcessor, error) {
	tx := r.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	return UnitOfWorkProcessor{tx: tx}, nil
}

type UnitOfWorkProcessor struct {
	tx *gorm.DB
}

func (r UnitOfWorkProcessor) ContentWriter() repository.ContentWriter {
	return NewContentWriterRepository(r.tx)
}

func (r UnitOfWorkProcessor) Rollback() error {
	return r.tx.Rollback().Error
}

func (r UnitOfWorkProcessor) Commit() error {
	return r.tx.Commit().Error
}
