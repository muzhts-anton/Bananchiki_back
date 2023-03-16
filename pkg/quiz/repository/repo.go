package quizrep

import (
	"banana/pkg/domain"
)

type dbQuizRepository struct {
	data []domain.Quiz
}

func InitQuizRep() domain.QuizRepository {
	return &dbQuizRepository{
		data: make([]domain.Quiz, 1),
	}
}

func (r *dbQuizRepository) CreateQuiz(q *domain.Quiz) error {
	return nil
}

func (r *dbQuizRepository) VoteQuiz(q *domain.Quiz) error {
	return nil
}

func (r *dbQuizRepository) ShowQuiz(q *domain.Quiz) error {
	return nil
}
