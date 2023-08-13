package disk

import (
	"testing"

	"github.com/prometheus/client_golang/prometheus"
)

func TestRegisterMemoryMetrics(t *testing.T) {
	diskMetrics := RegisterDiskMetrics()

	reg := prometheus.NewRegistry() // Create a new registry to register metrics
	reg.MustRegister(diskMetrics.devices)
	reg.MustRegister(diskMetrics.partitions)
	reg.MustRegister(diskMetrics.readRatePerApp)
	reg.MustRegister(diskMetrics.writeRatePerApp)

	_, err := reg.Gather() // Verify if the metrics are registered
	if err != nil {
		t.Fatalf("failed to gather metrics: %v", err)
	}
}
