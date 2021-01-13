package utils

import (
    "encoding/json"
    "net/http"
)

const (
    jsonParseFailed = "{\"reason\":\"failed to parse json\"}"
)

type JSONResponse struct {
    Success bool        `json:"success"`
    Data    interface{} `json:"data"`
    Message interface{} `json:"message"`
}

// Complete JSON response
func JSON(w http.ResponseWriter, r *http.Request, success bool, message *string, data interface{}) {
    var payload []byte

    var err error

    resp := JSONResponse{
        Success: success,
        Data:    data,
        Message: message,
    }

    if payload, err = json.Marshal(resp); err != nil {
        w.Header().Set("Content-Type", "application/json; charset=utf-8")
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte(jsonParseFailed))
        return
    }

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")

    w.Write(payload)
}

// Simple JSON response success
func JSONSuccess(w http.ResponseWriter, r *http.Request, data interface{}) {
    JSON(w, r, true, nil, data)
}

// Simple JSON response error
func JSONError(w http.ResponseWriter, r *http.Request, message string) {
    JSON(w, r, false, &message, nil)
}
