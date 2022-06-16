package usecase

import (
	"CoopeLunch-api/domain"
)

type PlanParticipantUsersInteractor struct {
	repository PlanParticipantUsersRepository
}

func NewPlanParticipantUsersInteractor(planParticipantUsers PlanParticipantUsersRepository) domain.PlanParticipantUsersInteractor {
	return &PlanParticipantUsersInteractor{
		repository: planParticipantUsers,
	}
}

type PlanParticipantUsersRepository interface {
	Insert(domain.TPlanParticipantUsersInsert) (id int, err error)
	GetByUserId(userId int) ([]domain.TPlanWithParticipantUsers, error)
}

func (interactor *PlanParticipantUsersInteractor) InsertPlanParticipantUsers(planParticipantUsers domain.TPlanParticipantUsersInsert) (id int, err error) {
	id, err = interactor.repository.Insert(planParticipantUsers)
	return
}

func (interactor *PlanParticipantUsersInteractor) ListPlanParticipantUsersByUserId(userId int) (plans []domain.TPlanWithParticipantUsers, err error) {
	plans, err = interactor.repository.GetByUserId(userId)
	return
}
