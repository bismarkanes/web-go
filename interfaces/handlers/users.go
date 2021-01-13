package handlers

import (
  "net/http"

  "encoding/json"
  "github.com/bismarkanes/web-go/application"
  "github.com/bismarkanes/web-go/infrastructure/utils"
  "github.com/go-chi/chi"
  "strconv"
)

type Users interface {
  GetUser(w http.ResponseWriter, r *http.Request)
  GetAllUser(w http.ResponseWriter, r *http.Request)
  CreateUser(w http.ResponseWriter, r *http.Request)
  UpdateUser(w http.ResponseWriter, r *http.Request)
  DeleteUser(w http.ResponseWriter, r *http.Request)
}

type users struct {
  userApp application.Users
}

func NewUsersHandler(up application.Users) Users {
  us := users{userApp: up}
  return &us
}

func (us *users) GetUser(w http.ResponseWriter, r *http.Request) {
  sUserID := chi.URLParam(r, "userID")

  userID, err := strconv.Atoi(sUserID)
  if err != nil {
    utils.JSONError(w, r, "ERR_USERID")
    return
  }

  resp := us.userApp.GetUser(userID)

  utils.JSON(w, r, true, "", resp)
}

func (us *users) GetAllUser(w http.ResponseWriter, r *http.Request) {
  resp := us.userApp.GetAllUser()

  utils.JSON(w, r, true, "", resp)
}

type createBody struct {
  Name string `json:"name"`
}

func (us *users) CreateUser(w http.ResponseWriter, r *http.Request) {
  b := createBody{}

  err := json.NewDecoder(r.Body).Decode(&b)
  if err != nil {
    utils.JSONError(w, r, http.StatusText(http.StatusBadRequest))
    return
  }

  utils.JSON(w, r, true, "", "CREATE")
}

type updateBody struct {
  ID int `json:"id"`
  createBody
}

func (us *users) UpdateUser(w http.ResponseWriter, r *http.Request) {
  sUserID := chi.URLParam(r, "userID")

  userID, err := strconv.Atoi(sUserID)
  if err != nil {
    utils.JSONError(w, r, "ERR_USERID")
    return
  }

  u := updateBody{}

  err = json.NewDecoder(r.Body).Decode(&u)
  if err != nil {
    utils.JSONError(w, r, http.StatusText(http.StatusBadRequest))
    return
  }

  resp := us.userApp.GetUser(userID)

  utils.JSON(w, r, true, "", resp)
}

func (us *users) DeleteUser(w http.ResponseWriter, r *http.Request) {
  sUserID := chi.URLParam(r, "userID")

  userID, err := strconv.Atoi(sUserID)
  if err != nil {
    utils.JSONError(w, r, "ERR_USERID")
    return
  }

  resp := us.userApp.GetUser(userID)

  utils.JSON(w, r, true, "", resp)
}
