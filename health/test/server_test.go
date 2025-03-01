package health

import (
	"testing"

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
