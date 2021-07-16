package item

import (
	"net/http"

	"github.com/aaalik/api-keras/bootstrap"
	"github.com/aaalik/api-keras/helper"
	"github.com/gorilla/mux"

	modelItem "github.com/aaalik/api-keras/model/item"
)

func GetItems(w http.ResponseWriter, r *http.Request) {
	helper.Log.Info(r.Method + " " + r.RequestURI)

	var items []modelItem.Item

	bootstrap.DB.Find(&items)

	if len(items) == 0 {
		response := helper.ErrorResponse{
			Title:  "item empty",
			Detail: "The items list are empty",
			Code:   100,
		}

		helper.JSONResponse(w, helper.HTTPStatusError, http.StatusBadRequest, response)
		return
	}

	helper.JSONResponse(w, helper.HTTPStatusSuccess, http.StatusOK, items)
}

func GetItem(w http.ResponseWriter, r *http.Request) {
	helper.Log.Info(r.Method + " " + r.RequestURI)

	vars := mux.Vars(r)

	var item modelItem.Item

	bootstrap.DB.Find(&item, vars["id"])

	if (item == modelItem.Item{}) {
		response := helper.ErrorResponse{
			Title:  "item not found",
			Detail: "The requested item for given item id is not found",
			Code:   101,
		}

		helper.JSONResponse(w, helper.HTTPStatusError, http.StatusBadRequest, response)
		return
	}

	helper.JSONResponse(w, helper.HTTPStatusSuccess, http.StatusOK, item)
}
