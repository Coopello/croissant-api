package domain

type TUser struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type TUserInsert struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type TLoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TLoginUserDBResponse struct {
	ID           int    `json:"id"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}

type TUserResponse struct {
	ID    int    `json:"id"`
	Token string `json:"token"`
}

type UserInteractor interface {
	SighUp(TUserInsert) (TUserResponse, error)
	LoginUser(TLoginUser) (TUserResponse, error)
}
