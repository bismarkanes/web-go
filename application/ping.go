package application

type Ping interface{}

type ping struct {
	value string
}

func NewPing() Ping {
	return &ping{
		value: "PONG",
	}
}
