package oauth

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/aaalik/api-keras/bootstrap"
	"github.com/aaalik/api-keras/helper"
	"golang.org/x/crypto/bcrypt"

	"github.com/aaalik/api-keras/middleware"
	modelUser "github.com/aaalik/api-keras/model/user"
)

func Token(w http.ResponseWriter, r *http.Request) {
	helper.Log.Info(r.Method + " " + r.RequestURI)

	var user modelUser.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		helper.JSONResponse(w, helper.HTTPStatusError, http.StatusBadRequest, err.Error())
		return
	}

	password := user.Password

	bootstrap.DB.Where("email = ?", user.Email).First(&user)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		helper.JSONResponse(w, helper.HTTPStatusError, http.StatusUnauthorized, errors.New("email or password not match").Error())
		return
	}

	token, err := middleware.CreateToken(int64(user.ID))
	if err != nil {
		helper.JSONResponse(w, helper.HTTPStatusError, http.StatusUnauthorized, err.Error())
		return
	}

	helper.JSONResponse(w, helper.HTTPStatusSuccess, http.StatusOK, token)
}
