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
	err := ur.dbm.Execute(queryUpdEmotions, pid)
	if err != nil {
		log.Warn("{ReactionsUpd} in query: " + queryUpdEmotions)
		log.Error(err)
	}

	return err
}