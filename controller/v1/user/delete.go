package user

import (
	"net/http"
	"strconv"

	"github.com/aaalik/api-keras/bootstrap"
	"github.com/aaalik/api-keras/helper"
	"github.com/gorilla/mux"

	modelUser "github.com/aaalik/api-keras/model/user"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	helper.Log.Info(r.Method + " " + r.RequestURI)

	vars := mux.Vars(r)

	var user modelUser.User

	id, _ := strconv.Atoi(vars["id"])

	user.ID = uint32(id)

	bootstrap.DB.First(&user)

	bootstrap.DB.Delete(&user)

	helper.JSONResponse(w, helper.HTTPStatusSuccess, http.StatusNoContent, "")
}
