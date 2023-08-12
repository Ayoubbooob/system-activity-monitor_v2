package memory

// import (
// 	"testing"

// 	"github.com/prometheus/client_golang/prometheus"
// )

// func TestRegisterMemoryMetrics(t *testing.T) {
// 	memoryMetrics := RegisterMemoryMetrics()

// 	// Create a new registry to register metrics
// 	reg := prometheus.NewRegistry()
// 	reg.MustRegister(memoryMetrics.memoryTotal)
// 	reg.MustRegister(memoryMetrics.memoryFree)
// 	reg.MustRegister(memoryMetrics.memoryUsagePerApp)

// 	// Verify if the metrics are registered
// 	_, err := reg.Gather()
// 	if err != nil {
// 		t.Fatalf("failed to gather metrics: %v", err)
// 	}
// }
