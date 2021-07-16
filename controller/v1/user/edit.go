package user

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/aaalik/api-keras/bootstrap"
	"github.com/aaalik/api-keras/helper"
	"github.com/gorilla/mux"

	modelUser "github.com/aaalik/api-keras/model/user"
)

func EditUser(w http.ResponseWriter, r *http.Request) {
	helper.Log.Info(r.Method + " " + r.RequestURI)

	vars := mux.Vars(r)

	user := modelUser.User{}

	id, _ := strconv.Atoi(vars["id"])
	user.ID = uint32(id)

	result := bootstrap.DB.First(&user)

	if result.Error != nil {
		response := helper.ErrorResponse{
			Title:  "user not found",
			Detail: "The requested user for given user id is not found",
			Code:   201,
		}

		helper.JSONResponse(w, helper.HTTPStatusError, http.StatusBadRequest, response)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		helper.JSONResponse(w, helper.HTTPStatusError, http.StatusBadRequest, err.Error())
		return
	}

	bootstrap.DB.Model(&user).Updates(user)

	helper.JSONResponse(w, helper.HTTPStatusSuccess, http.StatusOK, user)
}
