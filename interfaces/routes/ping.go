package routes

import (
	"github.com/bismarkanes/web-go/application"
  "github.com/bismarkanes/web-go/infrastructure/utils"
  "net/http"
)

type Ping interface {
	GetPing(w http.ResponseWriter, r *http.Request)
}

type ping struct {
	ping application.Ping
}

type data struct {
  Contents string `json:"contents"`
}

func (p *ping) GetPing(w http.ResponseWriter, r *http.Request) {
  d := data{
    Contents: "PONG",
  }

  utils.JSON(w, r, true, "", d)
}

func NewPing(p application.Ping) Ping {
	return &ping{
		ping: p,
	}
}
