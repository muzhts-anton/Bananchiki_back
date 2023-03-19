package quizusc

import (
	"banana/pkg/domain"
)

type quizUsecase struct {
	quizRepo domain.QuizRepository
}

func InitQuizUsc(qr domain.QuizRepository) domain.QuizUsecase {
	return &quizUsecase{
		quizRepo: qr,
	}
}

func (u *quizUsecase) CreateQuiz(q domain.Quiz, cid, pid uint64) (uint64, error) {
	return u.quizRepo.CreateQuiz(q, pid)
}

func (u *quizUsecase) DeleteQuiz(qid, cid, pid uint64) error {
	return u.quizRepo.DeleteQuiz(qid, pid)
}

func (u *quizUsecase) UpdateQuiz(q domain.Quiz, cid, pid uint64) error {
	return u.quizRepo.UpdateQuiz(q, pid)
}

func (u *quizUsecase) CreateQuizVote(q domain.Vote, qid, cid uint64) error {
	return u.quizRepo.CreateQuizVote(q, qid)
}

func (u *quizUsecase) UpdateQuizVote(q []domain.Vote, qid, cid uint64) (err error) {
	var tmperr error
	for _, vote := range q {
		tmperr = u.quizRepo.UpdateQuizVote(vote, qid)
		if tmperr != nil && err == nil {
			err = tmperr
		}
	}

	return
}

func (u *quizUsecase) DeleteQuizVote(idx uint32, qid, cid uint64) error {
	return u.quizRepo.DeleteQuizVote(idx, qid)
}
