package handlers

import (
    "net/http"

    "encoding/json"
    "strconv"

    "github.com/bismarkanes/web-go/application"
    "github.com/bismarkanes/web-go/domain/request"
    "github.com/bismarkanes/web-go/infrastructure/utilities"
    "github.com/go-chi/chi"
)

type User interface {
    GetUser(w http.ResponseWriter, r *http.Request)
    GetAllUser(w http.ResponseWriter, r *http.Request)
    CreateUser(w http.ResponseWriter, r *http.Request)
    UpdateUser(w http.ResponseWriter, r *http.Request)
    DeleteUser(w http.ResponseWriter, r *http.Request)
}

type users struct {
    userApp application.User
}

func NewUserHandler(up application.User) User {
    us := users{userApp: up}
    return &us
}

func (us *users) GetUser(w http.ResponseWriter, r *http.Request) {
    sUserID := chi.URLParam(r, "userID")

    userID, err := strconv.Atoi(sUserID)
    if err != nil {
        utilities.JSONError(w, r, "ERR_USERID")
        return
    }

    resp := us.userApp.GetUser(userID)

    utilities.JSON(w, r, nil, resp)
}

func (us *users) GetAllUser(w http.ResponseWriter, r *http.Request) {
    resp := us.userApp.GetAllUser()

    utilities.JSON(w, r, nil, resp)
}

func (us *users) CreateUser(w http.ResponseWriter, r *http.Request) {
    b := request.User{}

    err := json.NewDecoder(r.Body).Decode(&b)
    if err != nil {
        utilities.JSONError(w, r, http.StatusText(http.StatusBadRequest))
        return
    }

    resp := us.userApp.CreateUser(b.Name, b.Email)

    utilities.JSON(w, r, nil, resp)
}

func (us *users) UpdateUser(w http.ResponseWriter, r *http.Request) {
    sUserID := chi.URLParam(r, "userID")

    userID, err := strconv.Atoi(sUserID)
    if err != nil {
        utilities.JSONError(w, r, "ERR_USERID")
        return
    }

    u := request.User{}

    err = json.NewDecoder(r.Body).Decode(&u)
    if err != nil {
        utilities.JSONError(w, r, http.StatusText(http.StatusBadRequest))
        return
    }

    resp := us.userApp.UpdateUser(userID, u.Name, u.Email)

    utilities.JSON(w, r, nil, resp)
}

func (us *users) DeleteUser(w http.ResponseWriter, r *http.Request) {
    sUserID := chi.URLParam(r, "userID")

    userID, err := strconv.Atoi(sUserID)
    if err != nil {
        utilities.JSONError(w, r, "ERR_USERID")
        return
    }

    err = us.userApp.DeleteUser(userID)
    if err != nil {
        utilities.JSONError(w, r, "ERR_DELETE")
        return
    }

    utilities.JSON(w, r, nil, nil)
}
