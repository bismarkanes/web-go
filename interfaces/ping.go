package routes

import (
	"github.com/bismarkanes/web-go/application"
	"net/http"
)

type Ping interface {
	GetPing(w http.ResponseWriter, r *http.Request)
}

type ping struct {
	ping application.Ping
}

func (p *ping) GetPing(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("PONG"))
}

func NewPing(p application.Ping) Ping {
	return &ping{
		ping: p,
	}
}
