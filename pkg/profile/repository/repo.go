package profrep

import (
	"banana/pkg/domain"
	"banana/pkg/utils/cast"
	"banana/pkg/utils/database"
	"banana/pkg/utils/hash"
	"banana/pkg/utils/log"
)

type dbProfRepository struct {
	dbm *database.DBManager
}

func InitProfRep(manager *database.DBManager) domain.ProfRepository {
	return &dbProfRepository{
		dbm: manager,
	}
}

func (ur *dbProfRepository) GetUserInfo(usrId uint64) (domain.User, error) {
	resp, err := ur.dbm.Query(queryGetUserInfo, usrId)
	if err != nil {
		log.Warn("{GetUserInfo} in query: " + queryGetUserInfo)
		log.Error(err)
		return domain.User{}, err
	}
	if len(resp) == 0 {
		return domain.User{}, domain.ErrNoUser
	}

	return domain.User{
		Id:       usrId,
		Username: cast.ToString(resp[0][0]),
		Email:    cast.ToString(resp[0][1]),
		Imgsrc:   cast.ToString(resp[0][2]),
	}, err
}

func (ur *dbProfRepository) GetAllPres(usrId uint64) ([]domain.ProfilePresInfo, error) {
	resp, err := ur.dbm.Query(queryGetAllPres, usrId)
	if err != nil {
		log.Warn("{GetAllPres} in query: " + queryGetAllPres)
		log.Error(err)
		return nil, err
	}

	out := make([]domain.ProfilePresInfo, 0)
	for i, p := range resp {
		out = append(out, domain.ProfilePresInfo{
			Id:          cast.ToUint64(p[0]),
			Name:        cast.ToString(p[1]),
			Code:        cast.ToString(p[2]),
			QuizNum:     cast.ToUint64(p[3]),
			ConSlideNum: cast.ToUint64(p[4]),
		})
		out[i].Hash = hash.EncodeToHash(out[i].Code)
	}

	return out, nil
}

func (ur *dbProfRepository) UpdateAvatar(userId uint64, avatarPath string) error{
	err := ur.dbm.Execute(queryUpdateAvatar, avatarPath, userId)
	if err != nil {
		log.Warn("{UpdateAvatar} in query: " + queryUpdateAvatar)
		log.Error(err)
		return err
	}
	
	return nil
}
