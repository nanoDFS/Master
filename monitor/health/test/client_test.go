package health

import (
	"testing"
	"time"

	"github.com/nanoDFS/Master/monitor/health"
	"github.com/nanoDFS/Master/utils"
)

func TestNewHealhClient(t *testing.T) {
	port := utils.RandLocalAddr()
	_, err := health.NewHealthClient(port)
	if err != nil {
		t.Errorf("failed to create monitor , %v", err)
	}
}

func TestHealthMonitorStop(t *testing.T) {
	port := utils.RandLocalAddr()
	monitor, err := health.NewHealthClient(port)
	if err != nil {
		t.Errorf("failed to create monitor , %v", err)
	}
	monitor.Start()
	monitor.Stop()
	time.Sleep(time.Second * 5)
}
