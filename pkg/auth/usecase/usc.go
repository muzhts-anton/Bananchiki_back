package authusc

import (
	"banana/pkg/domain"

	"golang.org/x/crypto/bcrypt"
)

type authUsecase struct {
	authRepo domain.AuthRepository
}

func InitAuthUsc(ar domain.AuthRepository) domain.AuthUsecase {
	return &authUsecase{
		authRepo: ar,
	}
}

func (au authUsecase) Register(us domain.User) (domain.User, error) {
	trimCredentials(&us.Email, &us.Username, &us.Password, &us.RepeatPassword)

	if us.Email == "" || us.Username == "" || us.Password == "" || us.RepeatPassword == "" {
		return domain.User{}, domain.ErrEmptyField
	}

	if err := validateEmail(us.Email); err != nil {
		return domain.User{}, err
	}

	if err := validateUsername(us.Username); err != nil {
		return domain.User{}, err
	}

	if err := validatePassword(us.Password); err != nil {
		return domain.User{}, err
	}

	if us.Password != us.RepeatPassword {
		return domain.User{}, domain.ErrUnmatchedPasswords
	}

	if _, err := au.authRepo.GetUserByEmail(us.Email); err == nil {
		return domain.User{}, domain.ErrEmailExists
	}

	idupd, err := au.authRepo.CreateUser(us)
	if err != nil {
		return domain.User{}, err
	}

	out, _ := au.authRepo.GetUserById(idupd)

	return out.ClearPasswords(), nil
}

func (au authUsecase) Login(ub domain.UserBasic) (domain.User, error) {
	if ub.Email == "" || ub.Password == "" {
		return domain.User{}, domain.ErrEmptyField
	}

	usr, err := au.authRepo.GetUserByEmail(ub.Email)
	if err != nil {
		return domain.User{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(ub.Password)); err != nil {
		return domain.User{}, domain.ErrBadPassword
	}

	return usr.ClearPasswords(), nil
}
