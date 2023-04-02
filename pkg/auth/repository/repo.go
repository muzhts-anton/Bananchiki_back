package autrep

import (
	"banana/pkg/domain"
	"banana/pkg/utils/cast"
	"banana/pkg/utils/database"
	"banana/pkg/utils/log"

	"golang.org/x/crypto/bcrypt"
)

type dbAuthRepository struct {
	dbm *database.DBManager
}

func InitAutRep(manager *database.DBManager) domain.AuthRepository {
	return &dbAuthRepository{
		dbm: manager,
	}
}

func (ur *dbAuthRepository) GetUserByEmail(email string) (domain.User, error) {
	resp, err := ur.dbm.Query(queryGetByEmail, email)
	if err != nil {
		log.Warn("{GetByEmail} in query: " + queryGetByEmail)
		log.Error(err)
		return domain.User{}, err
	}

	if len(resp) == 0 {
		return domain.User{}, domain.ErrNoUser
	}

	return domain.User{
		Id:             cast.ToUint64(resp[0][0]),
		Username:       cast.ToString(resp[0][1]),
		Password:       cast.ToString(resp[0][4]),
		Email:          cast.ToString(resp[0][2]),
		Imgsrc:         cast.ToString(resp[0][3]),
		RepeatPassword: cast.ToString(resp[0][4]),
	}, nil
}

func (ur *dbAuthRepository) GetUserById(id uint64) (domain.User, error) {
	resp, err := ur.dbm.Query(queryGetById, id)
	if err != nil {
		log.Warn("{GetById} in query: " + queryGetById)
		log.Error(err)
		return domain.User{}, err
	}

	if len(resp) == 0 {
		return domain.User{}, domain.ErrNoUser
	}

	return domain.User{
		Id:       cast.ToUint64(resp[0][0]),
		Username: cast.ToString(resp[0][1]),
		Email:    cast.ToString(resp[0][2]),
		Imgsrc:   cast.ToString(resp[0][3]),
	}, nil
}

func (ur *dbAuthRepository) CreateUser(us domain.User) (uint64, error) {
	passwordByte, err := bcrypt.GenerateFromPassword([]byte(us.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error(err)
		return 0, err
	}

	resp, err := ur.dbm.Query(queryCreateUser, us.Username, us.Email, passwordByte)
	if err != nil {
		log.Warn("{AddUser} in query: " + queryCreateUser)
		log.Error(err)
		return 0, err
	}

	return cast.ToUint64(resp[0][0]), nil
}
