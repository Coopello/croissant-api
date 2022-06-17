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

func (repo *PlanParticipantUsersRepository) GetByUserId(userId int) (plans []domain.TPlan, err error) {
	rows, err := repo.Query(
		"SELECT PlanId as ID, ShopName, MeetPlace, MaxPeopleNumber, MinPeopleNumber, MeetTime, PlanStatus, OwnerUserId, ParticipantUsersCount FROM plans INNER JOIN plan_participant_users ON plans.ID = plan_participant_users.PlanId WHERE plan_participant_users.UserId = ? AND plans.PlanStatus <= 2",
		userId,
	)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var p domain.TPlan
		err := rows.Scan(
			&p.ID, &p.ShopName, &p.MeetPlace, &p.MaxPeopleNumber, &p.MinPeopleNumber, &p.MeetTime, &p.PlanStatus, &p.OwnerUserId, &p.ParticipantUsersCount,
		)
		if err != nil {
			panic(err.Error())
		}
		plans = append(plans, p)
	}

	return
}

func (repo *PlanParticipantUsersRepository) GetHistoriesByUserId(userId int) (plans []domain.TPlan, err error) {
	rows, err := repo.Query(
		"SELECT PlanId as ID, ShopName, MeetPlace, MaxPeopleNumber, MinPeopleNumber, MeetTime, PlanStatus, OwnerUserId, ParticipantUsersCount FROM plans INNER JOIN plan_participant_users ON plans.ID = plan_participant_users.PlanId WHERE plan_participant_users.UserId = ? AND plans.PlanStatus = 3",
		userId,
	)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var p domain.TPlan
		err := rows.Scan(
			&p.ID, &p.ShopName, &p.MeetPlace, &p.MaxPeopleNumber, &p.MinPeopleNumber, &p.MeetTime, &p.PlanStatus, &p.OwnerUserId, &p.ParticipantUsersCount,
		)
		if err != nil {
			panic(err.Error())
		}
		plans = append(plans, p)
	}

	return
}
