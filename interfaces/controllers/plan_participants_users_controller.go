package controllers

import (
	"CoopeLunch-api/domain"
	"CoopeLunch-api/interfaces/database"
	"CoopeLunch-api/usecase"
	"encoding/json"
	"net/http"
	"strconv"
)

type PlanParticipantUsersController struct {
	interactor domain.PlanParticipantUsersInteractor
}

func NewPlanParticipantUsersController(sqlHandler database.SqlHandler) *PlanParticipantUsersController {
	return &PlanParticipantUsersController{
		interactor: usecase.NewPlanParticipantUsersInteractor(
			&database.PlanParticipantUsersRepository{
				SqlHandler: sqlHandler,
			},
		),
	}
}

func (controller *PlanParticipantUsersController) PlanParticipantUsersInsertView(w http.ResponseWriter, r *http.Request) {
	var planParticipantUsers domain.TPlanParticipantUsersInsert
	err := json.NewDecoder(r.Body).Decode(&planParticipantUsers)
	if err != nil {
		response(w, err, map[string]interface{}{"error": err.Error()})
		return
	}
	id, err := controller.interactor.InsertPlanParticipantUsers(planParticipantUsers)
	if err != nil {
		response(w, err, map[string]interface{}{"error": err.Error()})
		return
	}
	response(w, nil, map[string]interface{}{"id": id})
	return
}

func (controller *PlanParticipantUsersController) PlanParticipantUsersListByUserIdView(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(r.URL.Query().Get("userId"))
	if err != nil {
		response(w, err, map[string]interface{}{"error": err.Error()})
		return
	}
	planParticipantUsers, err := controller.interactor.ListPlanParticipantUsersByUserId(userId)
	if len(planParticipantUsers) == 0 {
		planParticipantUsers = make([]domain.TPlan, 0)
	}
	response(w, err, map[string]interface{}{"data": planParticipantUsers})
	return
}

func (controller *PlanParticipantUsersController) PlanParticipantUsersListHistoriesView(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(r.URL.Query().Get("userId"))
	if err != nil {
		response(w, err, map[string]interface{}{"error": err.Error()})
		return
	}
	planParticipantUsers, err := controller.interactor.ListPlanHistoriesByUserId(userId)
	// 配列が空の場合にnullではなく、空の配列がレスポンスになるように
	if len(planParticipantUsers) == 0 {
		planParticipantUsers = make([]domain.TPlan, 0)
	}
	response(w, err, map[string]interface{}{"data": planParticipantUsers})
	return
}
