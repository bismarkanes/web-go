package routes

import (
	"github/web-go/src/application"
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

func NewPing(ping application.Ping) Ping {
	return &ping{
		ping: ping,
	}
}
