package handlers

import (
  "github.com/bismarkanes/web-go/infrastructure"
  "github.com/bismarkanes/web-go/infrastructure/utils"
  "net/http"
)

func GetPing(w http.ResponseWriter, r *http.Request) {
  utils.JSONSuccess(w, r, constants.PONG)
}
