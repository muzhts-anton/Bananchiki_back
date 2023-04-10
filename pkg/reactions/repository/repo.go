package reacrep

import (
	"banana/pkg/domain"
	"banana/pkg/utils/cast"
	"banana/pkg/utils/database"
	"banana/pkg/utils/hash"
	"banana/pkg/utils/log"
)

type dbReacRepository struct {
	dbm *database.DBManager
}

func InitReacRep(manager *database.DBManager) domain.ReacRepository {
	return &dbReacRepository{
		dbm: manager,
	}
}

func (ur *dbReacRepository) GetPresIdByHash(h string) (uint64, error) {
	resp, err := ur.dbm.Query(queryGetAllPres)
	if err != nil {
		log.Warn("{GetPresIdByHash} in query: " + queryGetAllPres)
		log.Error(err)
		return 0, err
	}

	for _, pres := range resp {
		if code := cast.ToString(pres[1]); hash.EncodeToHash(code) == h {
			return cast.ToUint64(pres[0]), nil
		}
	}

	return 0, domain.ErrCodeNotFound
}

func (ur *dbReacRepository) ReactionsUpd(pid uint64, emo domain.PresEmotions) error {
	err := ur.dbm.Execute(queryUpdEmotions, emo.Like, emo.Love, emo.Laughter, emo.Surprise, emo.Sad, pid)
	if err != nil {
		log.Warn("{ReactionsUpd} in query: " + queryUpdEmotions)
		log.Error(err)
	}

	return err
}

func (ur *dbReacRepository) QuestionAsk(pid uint64, q domain.Question) error {
	err := ur.dbm.Execute(queryQuestionAsk, pid, q.Idx, q.Option, q.Likes)
	if err != nil {
		log.Warn("{QuestionAsk} in query: " + queryQuestionAsk)
		log.Error(err)
	}

	return err
}

func (ur *dbReacRepository) QuestionLike(pid uint64, idx uint64) error {
	err := ur.dbm.Execute(queryQuestionLike, pid, idx)
	if err != nil {
		log.Warn("{QuestionLike} in query: " + queryQuestionLike)
		log.Error(err)
	}

	return err
}
