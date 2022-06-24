package main

import (
	"CoopeLunch-api/infrastructure"
	"CoopeLunch-api/interfaces/controllers"
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	sqlHandler := infrastructure.NewSqlHandler()

	// User
	userController := controllers.NewUserController(sqlHandler)
	router.Handle("/signUp", http.HandlerFunc(userController.SighUpView))
	router.Handle("/login", http.HandlerFunc(userController.LoginUserView))

	// Plan
	planController := controllers.NewPlanController(sqlHandler)
	router.Handle("/plan/", http.HandlerFunc(planController.PlanListView))
	router.Handle("/plan/users/", http.HandlerFunc(planController.PlanListByUserIdView))
	router.Handle("/plan/create/", http.HandlerFunc(planController.PlanInsertView))

	// PlanParticipantUsers
	planParticipantUsersController := controllers.NewPlanParticipantUsersController(sqlHandler)
	router.Handle("/plan-participant-users/", http.HandlerFunc(planParticipantUsersController.PlanParticipantUsersInsertView))
	router.Handle("/plan-participant-users/users/", http.HandlerFunc(planParticipantUsersController.PlanParticipantUsersListByUserIdView))
	router.Handle("/plan-participant-users/histories/", http.HandlerFunc(planParticipantUsersController.PlanParticipantUsersListHistoriesView))

	srv := http.Server{
		Addr:    ":8000",
		Handler: router,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
