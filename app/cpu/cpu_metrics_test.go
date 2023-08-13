package cpu

import (
	"testing"

	"github.com/prometheus/client_golang/prometheus"
)

func TestRegisterCpuMetrics(t *testing.T) {
	cpuMetrics := RegisterCpuMetrics()

	// Create a new registry to register metrics
	reg := prometheus.NewRegistry()
	reg.MustRegister(cpuMetrics.CpuUsage)
	reg.MustRegister(cpuMetrics.CpuFree)
	reg.MustRegister(cpuMetrics.CpuUsagePerApp)

	// Verify if the metrics are registered
	_, err := reg.Gather()
	if err != nil {
		t.Fatalf("failed to gather metrics: %v", err)
	}
}
