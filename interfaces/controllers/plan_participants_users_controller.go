package controllers

import (
	"CoopeLunch-api/domain"
	"CoopeLunch-api/interfaces/database"
	"CoopeLunch-api/usecase"
	"encoding/json"
	"fmt"
	"net/http"
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
	fmt.Println(planParticipantUsers)
	if err != nil {
		response(w, err, map[string]interface{}{"error": err})
		return
	}
	id, err := controller.interactor.InsertPlanParticipantUsers(planParticipantUsers)
	if err != nil {
		response(w, err, map[string]interface{}{"error": err})
		return
	}
	response(w, nil, map[string]interface{}{"id": id})
	return
}
