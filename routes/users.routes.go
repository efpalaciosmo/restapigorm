package routes

import (
	"encoding/json"
	"github.com/efpalaciosmo/RestApiGorm/db"
	"github.com/efpalaciosmo/RestApiGorm/models"
	"net/http"
)

func GetUserHandler(writer http.ResponseWriter, r *http.Request) {
	writer.Write([]byte("get user"))
}

func GetUsersHandler(writer http.ResponseWriter, r *http.Request) {
	var users []models.User
	db.DB.Find(&users)
	json.NewEncoder(writer).Encode(&users)
}

func PostUserHandler(writer http.ResponseWriter, r *http.Request) {
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
	writer.Write([]byte("delete"))
}
