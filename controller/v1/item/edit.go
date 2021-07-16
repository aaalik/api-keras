package item

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/aaalik/api-keras/bootstrap"
	"github.com/aaalik/api-keras/helper"
	"github.com/gorilla/mux"

	modelItem "github.com/aaalik/api-keras/model/item"
)

func EditItem(w http.ResponseWriter, r *http.Request) {
	helper.Log.Info(r.Method + " " + r.RequestURI)

	vars := mux.Vars(r)

	item := modelItem.Item{}

	id, _ := strconv.Atoi(vars["id"])
	item.ID = uint32(id)

	result := bootstrap.DB.First(&item)

	if result.Error != nil {
		response := helper.ErrorResponse{
			Title:  "item not found",
			Detail: "The requested item for given item id is not found",
			Code:   101,
		}

		helper.JSONResponse(w, helper.HTTPStatusError, http.StatusBadRequest, response)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		helper.JSONResponse(w, helper.HTTPStatusError, http.StatusBadRequest, err.Error())
		return
	}

	bootstrap.DB.Model(&item).Updates(item)

	helper.JSONResponse(w, helper.HTTPStatusSuccess, http.StatusOK, item)
}
