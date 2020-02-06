// Package onload automatically starts a goroutine running the simple Go runtime collector
package onload

import (
	"context"
	"log"
	"time"

	metrics "github.com/jonahgeorge/hmetrics"
	"go.opencensus.io/stats/view"
)

const interval = 20 * time.Second

var (
	pauseTotalNS uint64
	numGC        uint32
	ticker       = time.NewTicker(interval)
)

func init() {
	if err := view.Register(metrics.RuntimeViews...); err != nil {
		log.Fatalf("failed to register view: %v", err)
	}

	go func() {
		for range ticker.C {
			pauseTotalNS, numGC = metrics.Collect(context.Background(), pauseTotalNS, numGC)
		}
	}()
}
