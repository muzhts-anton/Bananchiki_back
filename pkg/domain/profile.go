package domain

import (
	"net/http"
)

const(
	AvatarFilePath = "/static/profile/avatars/"
)

type ProfilePresInfo struct {
	Id          uint64 `json:"id"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	Hash        string `json:"hash"`
	QuizNum     uint64 `json:"quizNum"`
	ConSlideNum uint64 `json:"convSlideNum"`
}

type Profile struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Imgsrc   string `json:"imgsrc"`

	Presentations []ProfilePresInfo `json:"presentations"`
}

type ProfUsecase interface {
	GetProfile(usrId uint64) (Profile, error)
	UpdateProfileAvatar(filename string, usrId uint64) error
}

type ProfRepository interface {
	GetUserInfo(userId uint64) (User, error)
	GetAllPres(userId uint64) ([]ProfilePresInfo, error)
	UpdateAvatar(userId uint64, avatarPath string) error
}
