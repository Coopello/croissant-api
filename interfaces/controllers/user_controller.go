package controllers

import (
	"CoopeLunch-api/domain"
	"CoopeLunch-api/interfaces/database"
	"CoopeLunch-api/usecase"
	"encoding/json"
	"net/http"
)

type UserController struct {
	interactor domain.UserInteractor
}

func NewUserController(sqlHandler database.SqlHandler) *UserController {
	return &UserController{
		interactor: usecase.NewUserInteractor(
			&database.UserRepository{
				SqlHandler: sqlHandler,
			},
		),
	}
}

func (controller *UserController) SighUpView(w http.ResponseWriter, r *http.Request) {
	var user domain.TUserInsert
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		response(w, err, nil)
		return
	}
	resUser, err := controller.interactor.SighUp(user)
	if err != nil {
		response(w, err, nil)
		return
	}
	response(w, nil, map[string]interface{}{"data": resUser})
	return
}

func (controller *UserController) LoginUserView(w http.ResponseWriter, r *http.Request) {
	var user domain.TLoginUser
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		response(w, err, nil)
		return
	}
	resUser, err := controller.interactor.LoginUser(user)
	if err != nil {
		response(w, err, nil)
		return
	}
	response(w, nil, map[string]interface{}{"data": resUser})
	return
}
