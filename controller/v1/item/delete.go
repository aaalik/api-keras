package item

import (
	"net/http"
	"strconv"

	"github.com/aaalik/api-keras/bootstrap"
	"github.com/aaalik/api-keras/helper"
	"github.com/gorilla/mux"

	modelItem "github.com/aaalik/api-keras/model/item"
)

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	helper.Log.Info(r.Method + " " + r.RequestURI)

	vars := mux.Vars(r)

	var item modelItem.Item

	id, _ := strconv.Atoi(vars["id"])

	item.ID = uint32(id)

	bootstrap.DB.First(&item)

	bootstrap.DB.Delete(&item)

	helper.JSONResponse(w, helper.HTTPStatusSuccess, http.StatusNoContent, "")
}
