package event

import (
	"time"
)

type EventType int

type Event struct {
	EventType
	Message string
	time    time.Time
}

const (
	CREATE EventType = iota // EnumIndex = 0
	VIEW                    // EnumIndex = 1
)

func (event EventType) String() string {
	return [...]string{"CREATE", "VIEW"}[event]
}

func (event EventType) EnumIndex() int {
	return int(event)
}

func WriteLog(event Event) {
	// nothing to do
}
