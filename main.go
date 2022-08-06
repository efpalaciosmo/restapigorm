package main

import (
	"github.com/efpalaciosmo/RestApiGorm/db"
	"github.com/efpalaciosmo/RestApiGorm/models"
	"github.com/efpalaciosmo/RestApiGorm/routes"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	db.DBConnection()

	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})

	router := mux.NewRouter()

	router.HandleFunc("/", routes.HomeHandler)
	router.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	router.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	router.HandleFunc("/users", routes.CreateUserHandler).Methods("POST")
	router.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")

	//tasks routs

	router.HandleFunc("/tasks", routes.GetTasksHander).Methods("GET")
	router.HandleFunc("/tasks/{id}", routes.GetTaskHander).Methods("GET")
	router.HandleFunc("/tasks", routes.CreateTaskHander).Methods("POST")
	router.HandleFunc("/tasks/{id}", routes.DeleteTaskHander).Methods("DELETE")

	http.ListenAndServe(":3000", router)
}
