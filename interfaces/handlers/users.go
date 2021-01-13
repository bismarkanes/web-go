package handlers

import (
	"github.com/bismarkanes/web-go/application"
  "github.com/bismarkanes/web-go/infrastructure/utils"
  "net/http"
)

type Users interface {
  GetUser(w http.ResponseWriter, r *http.Request)
}

type users struct {
  userApp application.Users
}

func NewUsersHandler(up application.Users) Users {
  us := users{userApp: up}
  return &us
}

func (us *users) GetUser(w http.ResponseWriter, r *http.Request) {
  // ID := 1

  resp := us.userApp.GetUser(1)
  // resp := "USERS"

  utils.JSON(w, r, true, "", resp)
}
