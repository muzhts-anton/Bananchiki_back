package presrep

import (
	"banana/pkg/domain"
)

type dbPresRepository struct {
}

func InitPresRep() domain.PresRepository {
	return &dbPresRepository{}
}

func (r *dbPresRepository) CreatePres(q *domain.Presentation) error {
	return nil
}

func (r *dbPresRepository) GetPres(q *domain.Presentation) error {
	return nil
}
