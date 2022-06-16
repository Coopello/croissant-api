package database

import (
	"CoopeLunch-api/domain"
)

type PlanParticipantUsersRepository struct {
	SqlHandler
}

func (repo *PlanParticipantUsersRepository) Insert(planParticipantUsers domain.TPlanParticipantUsersInsert) (id int, err error) {
	exe, err := repo.Execute(
		"INSERT INTO plan_participant_users (UserId, PlanId) VALUES (?, ?)",
		planParticipantUsers.UserId, planParticipantUsers.PlanId,
	)
	if err != nil {
		return
	}
	rawId, err := exe.LastInsertId()
	if err != nil {
		return
	}
	id = int(rawId)
	return
}
