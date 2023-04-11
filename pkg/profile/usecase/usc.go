package profusc

import (
	"banana/pkg/domain"
)

type profUsecase struct {
	profRepo domain.ProfRepository
}

func InitProfUsc(dr domain.ProfRepository) domain.ProfUsecase {
	return &profUsecase{
		profRepo: dr,
	}
}

func (u *profUsecase) GetProfile(usrId uint64) (domain.Profile, error) {
	usrInfo, err := u.profRepo.GetUserInfo(usrId)
	if err != nil {
		return domain.Profile{}, err
	}

	pres, err := u.profRepo.GetAllPres(usrId)
	if err != nil {
		return domain.Profile{}, err
	}

	return domain.Profile{
		Username:      usrInfo.Username,
		Email:         usrInfo.Email,
		Imgsrc:        usrInfo.Imgsrc,
		Presentations: pres,
	}, nil
}

func (u *profUsecase) UpdateProfileAvatar(filename string, usrId uint64) error{
	err := u.profRepo.UpdateAvatar(usrId, filename)
	if err != nil{
		return err
	}
	return nil
}
