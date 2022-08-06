package routes

import (
	"encoding/json"
	"github.com/efpalaciosmo/RestApiGorm/db"
	"github.com/efpalaciosmo/RestApiGorm/models"
	"github.com/gorilla/mux"
	"net/http"
)

func GetTasksHander(writer http.ResponseWriter, r *http.Request) {
	var tasks []models.Task
	db.DB.Find(&tasks)
	json.NewEncoder(writer).Encode(&tasks)
}

func GetTaskHander(writer http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var task models.Task
	db.DB.First(&task, params["id"])
	if task.ID == 0 {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte("Task not found"))
		return
	}
	json.NewEncoder(writer).Encode(&task)
}

func CreateTaskHander(writer http.ResponseWriter, r *http.Request) {
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)
	createTask := db.DB.Create(&task)
	error := createTask.Error
	if error != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte(error.Error()))
		return
	}

	json.NewEncoder(writer).Encode(&task)
}

func DeleteTaskHander(writer http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var task models.Task
	db.DB.First(&task, params["id"])

	if task.ID == 0 {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte("Task not found"))
		return
	}
	db.DB.Unscoped().Delete(&task)
	writer.Write([]byte("User deleted"))
}
