package constants

import (
	"time"
)

type Tick struct {
	Duration time.Duration
}

const TickDuration = time.Second * 1

type Dimensions struct {
	Width  int
	Height int
}
