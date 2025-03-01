package health

import (
	"testing"
	"time"

	"github.com/nanoDFS/Master/health"
	"github.com/nanoDFS/Master/utils"
)

func TestNewHealthMonitor(t *testing.T) {
	port := utils.RandLocalAddr()
	_, err := health.NewHealthMonitor(port)
	if err != nil {
		t.Errorf("failed to create monitor , %v", err)
	}
}

func TestHealthMonitorStop(t *testing.T) {
	port := utils.RandLocalAddr()
	monitor, err := health.NewHealthMonitor(port)
	if err != nil {
		t.Errorf("failed to create monitor , %v", err)
	}
	monitor.Listen()
	monitor.Stop()
	time.Sleep(time.Second * 5)
}
