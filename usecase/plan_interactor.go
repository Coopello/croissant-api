package usecase

import (
	"CoopeLunch-api/domain"
)

type PlanInteractor struct {
	repository PlanRepository
}

func NewPlanInteractor(plan PlanRepository) domain.PlanInteractor {
	return &PlanInteractor{
		repository: plan,
	}
}

type PlanRepository interface {
	All() ([]domain.TPlan, error)
	GetByUserId(userId int) ([]domain.TPlan, error)
	Insert(domain.TPlanInsert) (int, error)
	MakePlanEnd(planId int) (bool, error)
}

func (interactor *PlanInteractor) ListPlan() (plans []domain.TPlan, err error) {
	plans, err = interactor.repository.All()
	return
}

func (interactor *PlanInteractor) ListPlanByUserId(userId int) (plans []domain.TPlan, err error) {
	plans, err = interactor.repository.GetByUserId(userId)
	return
}

func (interactor *PlanInteractor) InsertPlan(plan domain.TPlanInsert) (id int, err error) {
	id, err = interactor.repository.Insert(plan)
	return
}
