package usecase

import (
	"CoopeLunch-api/domain"
)

type UserInteractor struct {
	repository UserRepository
}

func NewUserInteractor(user UserRepository) domain.UserInteractor {
	return &UserInteractor{
		repository: user,
	}
}

type UserRepository interface {
	SighUp(domain.TUserInsert) (domain.TUserResponse, error)
	LoginUser(domain.TLoginUser) (domain.TUserResponse, error)
}

func (interactor *UserInteractor) SighUp(user domain.TUserInsert) (domain.TUserResponse, error) {
	return interactor.repository.SighUp(user)
}

func (interactor *UserInteractor) LoginUser(user domain.TLoginUser) (domain.TUserResponse, error) {
	return interactor.repository.LoginUser(user)
}
