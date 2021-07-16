package user

import (
	"net/http"

	"github.com/aaalik/api-keras/bootstrap"
	"github.com/aaalik/api-keras/helper"
	"github.com/gorilla/mux"

	modelUser "github.com/aaalik/api-keras/model/user"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	helper.Log.Info(r.Method + " " + r.RequestURI)

	var users []modelUser.User

	bootstrap.DB.Find(&users)

	if len(users) == 0 {
		response := helper.ErrorResponse{
			Title:  "user empty",
			Detail: "The users list are empty",
			Code:   200,
		}

		helper.JSONResponse(w, helper.HTTPStatusError, http.StatusBadRequest, response)
		return
	}

	helper.JSONResponse(w, helper.HTTPStatusSuccess, http.StatusOK, users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	helper.Log.Info(r.Method + " " + r.RequestURI)

	vars := mux.Vars(r)

	var user modelUser.User

	bootstrap.DB.Find(&user, vars["id"])

	if (user == modelUser.User{}) {
		response := helper.ErrorResponse{
			Title:  "user not found",
			Detail: "The requested user for given user id is not found",
			Code:   201,
		}

		helper.JSONResponse(w, helper.HTTPStatusError, http.StatusBadRequest, response)
		return
	}

	helper.JSONResponse(w, helper.HTTPStatusSuccess, http.StatusOK, user)
}
