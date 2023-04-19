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

	if (emo.Emotions.Laughter != 0 && emo.Emotions.Laughter != 1){
		return domain.ErrWrongEmotions
	}

	if (emo.Emotions.Surprise != 0 && emo.Emotions.Surprise != 1){
		return domain.ErrWrongEmotions
	}

	if (emo.Emotions.Love != 0 && emo.Emotions.Love != 1){
		return domain.ErrWrongEmotions
	}

	if (emo.Emotions.Like != 0 && emo.Emotions.Like != 1){
		return domain.ErrWrongEmotions
	}

	if (emo.Emotions.Sad != 0 && emo.Emotions.Sad != 1){
		return domain.ErrWrongEmotions
	}

	if (emo.Emotions.Laughter + emo.Emotions.Like + emo.Emotions.Love + emo.Emotions.Surprise + emo.Emotions.Sad != 1){
		return domain.ErrWrongEmotions
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

func (du reacUsecase) QuestionLike(h string, idx uint64) error {
	presId, err := du.reacRepo.GetPresIdByHash(h)
	if err != nil {
		return err
	}

	return du.reacRepo.QuestionLike(presId, idx)
}
