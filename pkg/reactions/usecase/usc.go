package reacusc

import (
	"banana/pkg/domain"
)

type reacUsecase struct {
	reacRepo domain.ReacRepository
}

func InitReacUsc(dr domain.ReacRepository) domain.ReacUsecase {
	return &reacUsecase{
		reacRepo: dr,
	}
}

func (du reacUsecase) ReactionsUpd(emo domain.ReactionsApi) error {
	presId, err := du.reacRepo.GetPresIdByHash(emo.PresHash)
	if err != nil {
		return err
	}

	return du.reacRepo.ReactionsUpd(presId, emo.Emotions)
}

func (du reacUsecase) QuestionAsk(h string, q domain.Question) error {
	presId, err := du.reacRepo.GetPresIdByHash(h)
	if err != nil {
		return err
	}

	return du.reacRepo.QuestionAsk(presId, q)
}
