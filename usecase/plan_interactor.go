package usecase

import "CoopeLunch-api/domain"

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
}

func (interactor *PlanInteractor) ListPlan() (plans []domain.TPlan, err error) {
	plans, err = interactor.repository.All()
	return
}
