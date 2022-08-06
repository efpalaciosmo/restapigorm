package routes

import (
	"encoding/json"
	"github.com/efpalaciosmo/RestApiGorm/db"
	"github.com/efpalaciosmo/RestApiGorm/models"
	"github.com/gorilla/mux"
	"net/http"
)

func GetUserHandler(writer http.ResponseWriter, r *http.Request) {
	var user models.User
	//use to get the params send on request
	params := mux.Vars(r)
	db.DB.First(&user, params["id"])
	if user.ID == 0 {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte("User not found"))
		return
	}

	db.DB.Model(&user).Association("Tasks").Find(&user.Tasks)

	json.NewEncoder(writer).Encode(&user)
}

func GetUsersHandler(writer http.ResponseWriter, r *http.Request) {
	var users []models.User
	db.DB.Find(&users)
	json.NewEncoder(writer).Encode(&users)
}

func CreateUserHandler(writer http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	createdUser := db.DB.Create(&user)
	err := createdUser.Error
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest) //400
		writer.Write([]byte(err.Error()))
	}

	json.NewEncoder(writer).Encode(&user)
}

func DeleteUserHandler(writer http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)
	db.DB.First(&user, params["id"])
	if user.ID == 0 {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte("User not found"))
		return
	}
	//without unescoped you really don't delete the user, only put the date where the query has been sent it,
	// this change is on delete_at column
	db.DB.Unscoped().Delete(&user)
	writer.Write([]byte("User deleted"))
}
