package domain

type TPlanParticipantUsers struct {
	ID     int `json:"id"`
	UserId int `json:"user_id"`
	PlanId int `json:"plan_id"`
}

type TPlanParticipantUsersInsert struct {
	UserId int `json:"user_id"`
	PlanId int `json:"plan_id"`
}

type PlanParticipantUsersInteractor interface {
	InsertPlanParticipantUsers(TPlanParticipantUsersInsert) (int, error)
	ListPlanParticipantUsersByUserId(int) ([]TPlanWithParticipantUsers, error)
}
