package handlers

import (
    "github.com/bismarkanes/web-go/infrastructure/constants"
    "github.com/bismarkanes/web-go/infrastructure/utilities"
    "net/http"
)

func GetPing(w http.ResponseWriter, r *http.Request) {
    utilities.JSONSuccess(w, r, constants.PONG)
}
