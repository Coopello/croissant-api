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

func (repo *PlanParticipantUsersRepository) GetByUserId(userId int) (plans []domain.TPlanWithParticipantUsers, err error) {
	rows, err := repo.Query(
		"SELECT PlanId FROM plan_participant_users WHERE UserId = ?",
		userId,
	)
	if err != nil {
		panic(err.Error())
	}
	var planId int
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&planId)
		if err != nil {
			panic(err.Error())
		}
	}

	rows, err = repo.Query(
		"SELECT COUNT(*) FROM plan_participant_users WHERE PlanId = ?",
		planId,
	)
	var participantUsers int
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&participantUsers)
		if err != nil {
			panic(err.Error())
		}
	}

	rows, err = repo.Query(
		"SELECT * FROM plans WHERE ID = ? AND PlanStatus <= 2",
		planId,
	)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var p domain.TPlanWithParticipantUsers
		err := rows.Scan(
			&p.ID, &p.ShopName, &p.MeetPlace, &p.MaxPeopleNumber, &p.MinPeopleNumber, &p.MeetTime, &p.PlanStatus, &p.OwnerUserId,
		)
		if err != nil {
			panic(err.Error())
		}
		p.ParticipantUsers = participantUsers
		plans = append(plans, p)
	}

	return
}
