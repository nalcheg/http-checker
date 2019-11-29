package types

import (
	"sync"
	"time"
)

type Result struct {
	Wg           *sync.WaitGroup
	Host         string
	ResponseCode int
	ResponseTime float64
	Time         time.Time
}
