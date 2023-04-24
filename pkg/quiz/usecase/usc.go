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

func (u *quizUsecase) PollQuizVote(idx uint32, qid uint64, votername string, vid uint64) (err error) {
	if votername == "" && vid == 0 {
		err = u.quizRepo.PollQuizVote(idx, qid)
	} else {
		err = u.quizRepo.PollQuizVoteTracked(idx, qid, vid)
	}
	if err != nil {
		return
	}

	return u.quizRepo.CalculatePoints(idx, qid, vid)
}

func (u *quizUsecase) CompetitionStart(quizId uint64, presId uint64) error {
	return u.quizRepo.CompetitionStart(quizId, presId)
}

func (u *quizUsecase) CompetitionStop(quizId uint64, presId uint64) error {
	return u.quizRepo.CompetitionStop(quizId, presId)
}

func (u *quizUsecase) CompetitionVoterRegister(name string, hash string) (uint64, error) {
	presId, err := u.quizRepo.GetPresIdByHash(hash)
	if err != nil {
		return 0, err
	}

	return u.quizRepo.CompetitionVoterRegister(name, presId)
}

func (u *quizUsecase) GetCompetitionResult(presId uint64) ([]domain.ResultItem, error) {
	prevTop, err := u.quizRepo.GetCompetitionResult(presId)
	if err != nil {
		return nil, err
	}

	err = u.quizRepo.SetCompetitionResult(presId)
	if err != nil {
		return nil, err
	}

	currentTop, err := u.quizRepo.GetCompetitionResult(presId)
	if err != nil {
		return nil, err
	}

	out := make([]domain.ResultItem, len(prevTop))
	copy(out, prevTop)
	type elem struct {
		ri  domain.ResultItem
		idx int
	}
	buf := make([]elem, 0)
	for i := range prevTop {
		for j, cv := range currentTop {
			if out[i].Id == cv.Id {
				out[i] = cv
				break
			} else if j == len(prevTop) {
				buf = append(buf, elem{ri: cv, idx: i})
			}
		}
	}
	for _, e := range buf {
		out[e.idx] = e.ri
	}

	return out, nil
}
