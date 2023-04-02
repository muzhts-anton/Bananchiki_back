package domain

type UserBasic struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	Id             uint64 `json:"ID"`
	Username       string `json:"username"`
	Password       string `json:"password,omitempty"`
	Email          string `json:"email"`
	Imgsrc         string `json:"imgsrc"`
	RepeatPassword string `json:"repeatpassword,omitempty"`
}

func (us *User) ClearPasswords() User {
	us.Password = ""
	us.RepeatPassword = ""
	return *us
}

type SessionInfo struct {
	Status bool `json:"status"`
}

type AuthUsecase interface {
	Register(u User) (User, error)
	Login(u UserBasic) (User, error)
}

type AuthRepository interface {
	GetUserByEmail(e string) (User, error)
	GetUserById(id uint64) (User, error)
	CreateUser(u User) (uint64, error)
}
