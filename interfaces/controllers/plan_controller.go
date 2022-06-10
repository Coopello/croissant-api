package controllers

import (
	"CoopeLunch-api/domain"
	"CoopeLunch-api/interfaces/database"
	"CoopeLunch-api/tools"
	"CoopeLunch-api/usecase"
	"encoding/json"
	"net/http"
	"strconv"
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
	// 配列が空の場合にnullではなく、空の配列がレスポンスになるように
	if len(plans) == 0 {
		plans = make([]domain.TPlan, 0)
	}
	response(w, err, map[string]interface{}{"data": plans})
	return
}

func (controller *PlanController) PlanListByUserIdView(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(r.URL.Query().Get("userId"))
	if err != nil {
		response(w, err, nil)
		return
	}
	plans, err := controller.interactor.ListPlanByUserId(userId)
	// 配列が空の場合にnullではなく、空の配列がレスポンスになるように
	if len(plans) == 0 {
		plans = make([]domain.TPlan, 0)
	}
	response(w, err, map[string]interface{}{"data": plans})
	return
}

func (controller *PlanController) PlanInsertView(w http.ResponseWriter, r *http.Request) {
	var plan domain.TPlanInsert
	err := json.NewDecoder(r.Body).Decode(&plan)
	if err != nil {
		response(w, err, nil)
		return
	}
	if plan.MinPeopleNumber == 1 {
		plan.PlanStatus = tools.PLAN_STATUS["ESTABLISHED"]
	}
	id, err := controller.interactor.InsertPlan(plan)
	if err != nil {
		response(w, err, nil)
		return
	}
	response(w, nil, map[string]interface{}{"data": id})
	return
}
