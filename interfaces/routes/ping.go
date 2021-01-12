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

func (p *ping) GetPing(w http.ResponseWriter, r *http.Request) {
  utils.JSON(w, r, true, "", p.ping.GetValue())
}

func NewPing(p application.Ping) Ping {
	return &ping{
		ping: p,
	}
}
