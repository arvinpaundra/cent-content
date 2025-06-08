package repository

type UnitOfWork interface {
	Begin() (UnitOfWorkProcessor, error)
}

type UnitOfWorkProcessor interface {
	ContentWriter() ContentWriter

	Rollback() error
	Commit() error
}
