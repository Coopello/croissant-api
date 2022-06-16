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
}

func (interactor *PlanParticipantUsersInteractor) InsertPlanParticipantUsers(planParticipantUsers domain.TPlanParticipantUsersInsert) (id int, err error) {
	id, err = interactor.repository.Insert(planParticipantUsers)
	return
}
