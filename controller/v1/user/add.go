package user

import (
	"encoding/json"
	"net/http"

	"github.com/aaalik/api-keras/bootstrap"
	"github.com/aaalik/api-keras/helper"
	"golang.org/x/crypto/bcrypt"

	modelUser "github.com/aaalik/api-keras/model/user"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
	helper.Log.Info(r.Method + " " + r.RequestURI)

	user := modelUser.User{}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		helper.JSONResponse(w, helper.HTTPStatusError, http.StatusBadRequest, err.Error())
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		helper.JSONResponse(w, helper.HTTPStatusError, http.StatusBadRequest, err.Error())
		return
	}

	user.Password = string(hashedPassword)

	result := bootstrap.DB.Create(&user)

	if result.Error != nil {
		helper.JSONResponse(w, helper.HTTPStatusError, http.StatusBadRequest, result.Error.Error())
		return
	}

	helper.JSONResponse(w, helper.HTTPStatusSuccess, http.StatusOK, user)
}
