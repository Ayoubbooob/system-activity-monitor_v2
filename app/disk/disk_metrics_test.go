package disk

import (
	"testing"

	"github.com/prometheus/client_golang/prometheus"
)

func TestRegisterMemoryMetrics(t *testing.T) {
	diskMetrics := RegisterDiskMetrics()

	// Create a new registry to register metrics
	reg := prometheus.NewRegistry()
	reg.MustRegister(diskMetrics.devices)
	reg.MustRegister(diskMetrics.partitions)
	reg.MustRegister(diskMetrics.readRatePerApp)
	reg.MustRegister(diskMetrics.writeRatePerApp)

	// Verify if the metrics are registered
	_, err := reg.Gather()
	if err != nil {
		t.Fatalf("failed to gather metrics: %v", err)
	}
}
