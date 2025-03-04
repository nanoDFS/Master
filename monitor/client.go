package monitor

import (
	hc "github.com/nanoDFS/Master/monitor/health"
)

type Monitor struct {
	HC *hc.HealthClient
}

func NewMonitor(addr string) (*Monitor, error) {
	healthClient, _ := hc.NewHealthClient(addr)
	return &Monitor{
		HC: healthClient,
	}, nil
}

func (t *Monitor) Start() {
	t.HC.Start()
}

func (t *Monitor) Stop() {
	t.HC.Stop()
}
