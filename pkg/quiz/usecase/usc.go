package quizusc

import (
	"banana/pkg/domain"
	// "banana/pkg/utils"
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

func (u *quizUsecase) PollQuizVote(idx uint32, qid uint64) error {
	return u.quizRepo.PollQuizVote(idx, qid)
}


func (u *quizUsecase) CompetitionStart(quizId uint64, presId uint64) error{
	return u.quizRepo.CompetitionStart(quizId, presId)
}

func (u *quizUsecase) CompetitionStop(quizId uint64, presId uint64) error{
	return u.quizRepo.CompetitionStop(quizId, presId)
}

func (u *quizUsecase) CompetitionVoterRegister(name string, hash string) (uint64, error){
	presId, err := u.quizRepo.GetPresIdByHash(hash)
	if err != nil {
		return 0, err
	}

	return u.quizRepo.CompetitionVoterRegister(name, presId)
}

func contains(s []domain.ResultItem, str domain.ResultItem) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func (u *quizUsecase) GetCompetitionResult(presId uint64) ([]domain.ResultItem, error){
	prevResults, err := u.quizRepo.GetPrevCompetitionResult(presId)
	if err != nil {
		return []domain.ResultItem{}, err
	}

	currentResults, err := u.quizRepo.GetCurrentCompetitionResult(presId)
	if err != nil {
		return []domain.ResultItem{}, err
	}

	var newResults []domain.ResultItem

	for i, _ := range prevResults{
		if contains(prevResults, currentResults[i]){
			newResults = append(newResults, currentResults[i])
		}
	}

	for i, _ := range prevResults{
		if !contains(newResults, currentResults[i]){
			newResults = append(newResults, currentResults[i])
		}
	}

	return newResults, nil
}