package utilities

import (
    "encoding/json"
    "net/http"

    "github.com/bismarkanes/web-go/infrastructure/constants"
)

const (
    jsonParseFailed = "{\"reason\":\"failed to parse json\"}"
)

type ResponsMeta struct {
    Total int `json:"total"`
    Limit int `json:"limit"`
    Skip  int `json:"skip"`
}

// JSONResponse
// Status -> 200, 400, 401
// Message -> error message
// Data -> Payload
// Meta -> Pagination etc
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

    w.Header().Set(constants.HttpHeaderContentType, constants.HttpHeaderApplicationJson)

    if payload, err = json.Marshal(resp); err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte(jsonParseFailed))
        return
    }

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
