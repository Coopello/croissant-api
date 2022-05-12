package controllers

import (
	"CoopeLunch-api/domain"
	"CoopeLunch-api/interfaces/database"
	"CoopeLunch-api/usecase"
	"net/http"
)

type PlanController struct {
	interactor domain.PlanInteractor
}

func NewPlanController(sqlHandler database.SqlHandler) *PlanController {
	return &PlanController{
		interactor: usecase.NewPlanInteractor(
			&database.PlanRepository{
				SqlHandler: sqlHandler,
			},
		),
	}
}

func (controller *PlanController) PlanListView(w http.ResponseWriter, r *http.Request) {
	plans, err := controller.interactor.ListPlan()
	response(w, err, map[string]interface{}{"data": plans})
	return
}
