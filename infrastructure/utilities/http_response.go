package utilities

import (
    "encoding/json"
    "net/http"
)

const (
    jsonParseFailed = "{\"reason\":\"failed to parse json\"}"
)

type ResponsMeta struct {
    Total int `json:"total"`
    Limit int `json:"limit"`
    Skip  int `json:"skip"`
}

type JSONResponse struct {
    Status  int         `json:"status"`
    Type    interface{} `json:"type"`
    Data    interface{} `json:"data"`
    Message interface{} `json:"message"`
    Meta    interface{} `json:"meta"`
}

// Complete JSON response
func JSON(w http.ResponseWriter, r *http.Request, errorMessage *string, data interface{}) {
    var payload []byte

    var err error

    resp := JSONResponse{
        Data:    data,
        Message: errorMessage,
        Status:  http.StatusOK,
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
    JSON(w, r, nil, data)
}

// Simple JSON response error
func JSONError(w http.ResponseWriter, r *http.Request, message string) {
    JSON(w, r, &message, nil)
}
