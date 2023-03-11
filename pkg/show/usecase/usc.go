package quizusc

import (
	"banana/pkg/domain"
)

type quizUsecase struct {
	quizRepo domain.QuizRepository
}

func InitAnnUsc(qr domain.QuizRepository) domain.QuizUsecase {
	return &quizUsecase{
		quizRepo: qr,
	}
}

func (u *quizUsecase) CreateQuiz(s *domain.Quiz) error {
	return nil
}

func (r *quizUsecase) VoteQuiz(q *domain.Quiz) error {
	return nil
}

func (r *quizUsecase) ShowQuiz(q *domain.Quiz) error {
	return nil
}
