package database

import "CoopeLunch-api/domain"

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
