package utils

import (
	"context"
	"time"
)

func Tick(interval time.Duration, f func(ctx context.Context)) {
	ticker := time.NewTicker(interval)
	ctx := context.TODO()
	go func() {
		defer ticker.Stop()
		for {
			f(ctx)
			select {
			case <-ticker.C:
				continue
			}
		}
	}()
}
