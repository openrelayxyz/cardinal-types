package metrics

import (
  "github.com/rcrowley/go-metrics"
)

var (
  MajorRegistry = metrics.NewPrefixedRegistry("cardinal")
  MinorRegistry = metrics.NewPrefixedRegistry("cardinal.minor")
)

func NewMajorGauge(name string) metrics.Gauge {
  return metrics.NewRegisteredGauge(name, MajorRegistry)
}
func NewMajorCounter(name string) metrics.Counter {
  return metrics.NewRegisteredCounter(name, MajorRegistry)
}
func NewMajorTimer(name string) metrics.Timer {
  return metrics.NewRegisteredTimer(name, MajorRegistry)
}
func NewMajorMeter(name string) metrics.Meter {
  return metrics.NewRegisteredMeter(name, MajorRegistry)
}
func NewMinorGauge(name string) metrics.Gauge {
  return metrics.NewRegisteredGauge(name, MinorRegistry)
}
func NewMinorCounter(name string) metrics.Counter {
  return metrics.NewRegisteredCounter(name, MinorRegistry)
}
func NewMinorTimer(name string) metrics.Timer {
  return metrics.NewRegisteredTimer(name, MinorRegistry)
}
func NewMinorMeter(name string) metrics.Meter {
  return metrics.NewRegisteredMeter(name, MinorRegistry)
}

type MetricsAPI struct {}

func (api *MetricsAPI) Metrics() []metrics.Registry {
  return []metrics.Registry{
    MinorRegistry,
    MajorRegistry,
  }
}
