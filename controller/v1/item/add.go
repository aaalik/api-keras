package item

import (
	"encoding/json"
	"net/http"

	"github.com/aaalik/api-keras/bootstrap"
	"github.com/aaalik/api-keras/helper"

	modelItem "github.com/aaalik/api-keras/model/item"
)

func AddItem(w http.ResponseWriter, r *http.Request) {
	helper.Log.Info(r.Method + " " + r.RequestURI)

	item := modelItem.Item{}

	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		helper.JSONResponse(w, helper.HTTPStatusError, http.StatusBadRequest, err.Error())
		return
	}

	result := bootstrap.DB.Create(&item)

	if result.Error != nil {
		helper.JSONResponse(w, helper.HTTPStatusError, http.StatusBadRequest, result.Error.Error())
		return
	}

	helper.JSONResponse(w, helper.HTTPStatusSuccess, http.StatusOK, item)
}
