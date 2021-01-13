package handlers

import (
  "net/http"

  "github.com/bismarkanes/web-go/infrastructure/constants"
  "github.com/bismarkanes/web-go/infrastructure/utils"
)

func GetPing(w http.ResponseWriter, r *http.Request) {
  utils.JSON(w, r, true, "", constants.PONG)
}
