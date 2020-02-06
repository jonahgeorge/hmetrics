// Package metrics implements a simple Go runtime collector
package metrics

import (
	"context"
	"runtime"

	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
)

var (
	mGCCollections = stats.Float64("go_gc_collections", "", stats.UnitDimensionless)
	mGCPauseNS     = stats.Float64("go_gc_pause_ns", "", stats.UnitDimensionless)
	mHeapBytes     = stats.Float64("go_memory_heap_bytes", "", stats.UnitDimensionless)
	mStackBytes    = stats.Float64("go_memory_stack_bytes", "", stats.UnitDimensionless)
	mHeapObjects   = stats.Float64("go_memory_heap_objects", "", stats.UnitDimensionless)
	mGCGoal        = stats.Float64("go_gc_goal", "", stats.UnitDimensionless)
	mRoutines      = stats.Float64("go_routines", "", stats.UnitDimensionless)
)

var RuntimeViews = []*view.View{
	{Name: "go_gc_collections", Measure: mGCCollections, Aggregation: view.Count()},
	{Name: "go_gc_pause_ns", Measure: mGCPauseNS, Aggregation: view.Count()},
	{Name: "go_memory_heap_bytes", Measure: mHeapBytes, Aggregation: view.LastValue()},
	{Name: "go_memory_stack_bytes", Measure: mStackBytes, Aggregation: view.LastValue()},
	{Name: "go_memory_heap_objects", Measure: mHeapObjects, Aggregation: view.LastValue()},
	{Name: "go_gc_goal", Measure: mGCGoal, Aggregation: view.LastValue()},
	{Name: "go_routines", Measure: mRoutines, Aggregation: view.LastValue()},
}

var ms = new(runtime.MemStats)

func Collect(ctx context.Context, prevPauseTotalNs uint64, prevNumGC uint32) (uint64, uint32) {
	runtime.ReadMemStats(ms)

	stats.Record(ctx,
		mGCCollections.M(float64(ms.NumGC-prevNumGC)),
		mGCPauseNS.M(float64(ms.PauseTotalNs-prevPauseTotalNs)),
		mHeapBytes.M(float64(ms.Alloc)),
		mStackBytes.M(float64(ms.StackInuse)),
		mHeapObjects.M(float64(ms.Mallocs-ms.Frees)),
		mGCGoal.M(float64(ms.NextGC)),
		mRoutines.M(float64(runtime.NumGoroutine())),
	)

	return ms.PauseTotalNs, ms.NumGC
}
