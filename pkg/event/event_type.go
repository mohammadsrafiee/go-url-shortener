package event

import (
	"time"
)

type Type int

type Event struct {
	Type
	Message string
	time    time.Time
}

const (
	CREATE Type = iota // EnumIndex = 0
	VIEW               // EnumIndex = 1
)

func (event Type) String() string {
	return [...]string{"CREATE", "VIEW"}[event]
}

func (event Type) EnumIndex() int {
	return int(event)
}

func WriteLog(event Event) {
	// nothing to do
}
