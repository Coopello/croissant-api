package database

import (
	"CoopeLunch-api/domain"
	"fmt"
)

type PlanRepository struct {
	SqlHandler
}

func (repo *PlanRepository) All() (plans []domain.TPlan, err error) {
	rows, err := repo.Query(
		"SELECT * FROM plan",
	)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var p domain.TPlan

		err := rows.Scan(
			&p.ID, &p.ShopName, &p.MeetPlace, &p.MaxPeopleNumber, &p.MinPeopleNumber, &p.MeetTime, &p.PlanStatus, &p.OwnerUserId,
		)
		if err != nil {
			panic(err.Error())
		}
		plans = append(plans, p)
	}
	return
}

func (repo *PlanRepository) Insert(plan domain.TPlanInsert) (id int, err error) {
	fmt.Println("repo Insert", plan)
	exe, err := repo.Execute(
		"INSERT INTO plan(ShopName, MeetPlace, MaxPeopleNumber, MinPeopleNumber, MeetTime, OwnerUserId) VALUES(?, ?, ?, ?, ?, ?)",
		plan.ShopName, plan.MeetPlace, plan.MaxPeopleNumber, plan.MinPeopleNumber, plan.MeetTime, plan.OwnerUserId,
	)
	if err != nil {
		return id, err
	}

	rawId, err := exe.LastInsertId()
	if err != nil {
		return id, err
	}
	id = int(rawId)
	return
}
