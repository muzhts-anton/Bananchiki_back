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

func (u *quizUsecase) PollQuizVote(idx uint32, qid uint64, votername string, vid uint64) error {
	if votername == "" && vid == 0 {
		return u.quizRepo.PollQuizVote(idx, qid)
	}

	err := u.quizRepo.PollQuizVoteTracked(idx, qid, vid)
	if err != nil {
		return err
	}

	err = u.quizRepo.CalculatePoints(idx, qid, vid)
	if err != nil {
		return err
	}

	return u.quizRepo.FinishCompetition(qid)
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

	var idxs = make(map[uint64]int)
	for i, it := range prevTop {
		idxs[it.Id] = i
	}

	out := make([]domain.ResultItem, len(currentTop))
	buf := make([]domain.ResultItem, 0)
	for _, toper := range currentTop {
		if idx, ex := idxs[toper.Id]; !ex {
			buf = append(buf, toper)
		} else {
			out[idx] = toper
		}
	}

	var j int
LOOP:
	for _, b := range buf {
		for out[j].Id != 0 {
			j++
			if j >= len(out) {
				break LOOP
			}
		}
		out[j] = b
	}

	return out, nil
}
