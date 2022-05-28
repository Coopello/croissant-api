package domain

type TPlan struct {
	ID              int    `json:"id"`
	ShopName        string `json:"shop_name"`
	MeetPlace       string `json:"meet_place"`
	MaxPeopleNumber int    `json:"max_people_number"`
	MinPeopleNumber int    `json:"min_people_number"`
	MeetTime        string    `json:"meet_time"`
	PlanStatus      int    `json:"status"`
	OwnerUserId     int    `json:"owner_user_id"`
}

type TPlanInsert struct {
	ShopName        string `json:"shop_name"`
	MeetPlace       string `json:"meet_place"`
	MaxPeopleNumber int    `json:"max_people_number"`
	MinPeopleNumber int    `json:"min_people_number"`
	MeetTime        string    `json:"meet_time"`
	OwnerUserId     int    `json:"owner_user_id"`
}

type PlanInteractor interface {
	ListPlan() ([]TPlan, error)
	InsertPlan(TPlanInsert) (int, error)
}
