package utils

import (
	"time"
)

func GetTimeElapsed(datetime time.Time, elapsed *time.Duration) {
	currentTime := time.Now()
	*elapsed = currentTime.Sub(datetime)
}
