# hmetrics

A port of [`heroku/x/hmetrics`](https://github.com/heroku/x/tree/master/hmetrics) to OpenCensus.

## Usage

```go
package main

import (
	"log"
	"net/http"

	"contrib.go.opencensus.io/exporter/prometheus"
	_ "github.com/jonahgeorge/hmetrics/onload"
)

func main() {
	exporter, err := prometheus.NewExporter(prometheus.Options{})
	if err != nil {
		log.Fatal(err)
	}

	// Serve the scrape endpoint on port 9999.
	http.Handle("/metrics", exporter)
	log.Fatal(http.ListenAndServe(":9999", nil))
}
```

```
$ curl localhost:9999/metrics
# HELP go_gc_collections 
# TYPE go_gc_collections counter
go_gc_collections 5
# HELP go_gc_goal 
# TYPE go_gc_goal gauge
go_gc_goal 4.473924e+06
# HELP go_gc_pause_ns 
# TYPE go_gc_pause_ns counter
go_gc_pause_ns 5
# HELP go_memory_heap_bytes 
# TYPE go_memory_heap_bytes gauge
go_memory_heap_bytes 1.003448e+06
# HELP go_memory_heap_objects 
# TYPE go_memory_heap_objects gauge
go_memory_heap_objects 5430
# HELP go_memory_stack_bytes 
# TYPE go_memory_stack_bytes gauge
go_memory_stack_bytes 917504
# HELP go_routines 
# TYPE go_routines gauge
go_routines 3
```
