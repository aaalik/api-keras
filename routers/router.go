package routers

import (
	"net/http"
	"os"

	oauth "github.com/aaalik/api-keras/controller/oauth"
	v1Item "github.com/aaalik/api-keras/controller/v1/item"
	v1User "github.com/aaalik/api-keras/controller/v1/user"
	"github.com/aaalik/api-keras/helper"
	"github.com/aaalik/api-keras/middleware"
	"github.com/gorilla/mux"
)

func SetupRouter() {
	jwtMiddleware := middleware.InitMiddleware()
	r := mux.NewRouter()
	r.StrictSlash(true)

	// oauth router
	r.HandleFunc("/oauth/token", oauth.Token).Methods("POST")

	// item router
	s := r.PathPrefix("/v1/items").Subrouter()
	s.HandleFunc("/{id}", v1Item.GetItem).Methods("GET")
	s.HandleFunc("/{id}", v1Item.EditItem).Methods("PUT")
	s.HandleFunc("/{id}", v1Item.DeleteItem).Methods("DELETE")
	s.HandleFunc("", v1Item.AddItem).Methods("POST")
	s.HandleFunc("", v1Item.GetItems).Methods("GET")

	// user router
	s = r.PathPrefix("/v1/users").Subrouter()
	s.HandleFunc("/{id}", v1User.GetUser).Methods("GET")
	s.HandleFunc("/{id}", v1User.EditUser).Methods("PUT")
	s.HandleFunc("/{id}", v1User.DeleteUser).Methods("DELETE")
	s.HandleFunc("", v1User.AddUser).Methods("POST")
	s.HandleFunc("", v1User.GetUsers).Methods("GET")

	s.Use(jwtMiddleware.Handler)

	helper.Log.Info("Server started at port " + os.Getenv("SERVER_PORT"))
	http.ListenAndServe(":"+os.Getenv("SERVER_PORT"), r)
}

func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://"+r.Host+r.URL.String(), http.StatusPermanentRedirect)
}
