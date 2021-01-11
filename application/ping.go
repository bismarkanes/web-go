package application

import (
	"github.com/bismarkanes/web-go/infrastructure/constants"
)

type Ping interface {
	GetValue() string
}

type ping struct {
	value string
}

func (p *ping) GetValue() string {
	return p.value
}

func NewPing() Ping {
	return &ping{
		value: constants.PONG,
	}
}
