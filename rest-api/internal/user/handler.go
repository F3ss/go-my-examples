package user

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"rest_api/internal/handlers"
)

const (
	usersURL = "/api/users/"
	userURL  = "/api/users/:uuid"
)

type handler struct {
}

func NewHandler() handlers.Handler {
	return &handler{}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(usersURL, h.GetUsersList)
	router.GET(userURL, h.GetUserByUUID)
	router.POST(usersURL, h.CreateUser)
	router.PUT(userURL, h.UpdateUser)
	router.PATCH(userURL, h.PartialUpdateUser)
	router.DELETE(userURL, h.DeleteUser)
}

func (h *handler) GetUsersList(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(200)
	if _, err := w.Write([]byte("Here is list of users")); err != nil {
		log.Println(err)
	}
}

func (h *handler) GetUserByUUID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(200)
	if _, err := w.Write([]byte("Here is GetUserByUUID")); err != nil {
		log.Fatalln(err)
	}
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(201)
	if _, err := w.Write([]byte("Here is CreateUser")); err != nil {
		log.Println(err)
	}
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(204)
	if _, err := w.Write([]byte("Here is UpdateUser")); err != nil {
		log.Println(err)
	}
}

func (h *handler) PartialUpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(204)
	if _, err := w.Write([]byte("Here is PartialUpdateUser")); err != nil {
		log.Println(err)
	}
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(200)
	if _, err := w.Write([]byte("Here is DeleteUser")); err != nil {
		log.Fatalln(err)
	}
}
