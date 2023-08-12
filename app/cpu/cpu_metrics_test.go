package cpu

// import (
// 	"testing"

// 	"github.com/prometheus/client_golang/prometheus"
// )

// func TestRegisterCpuMetrics(t *testing.T) {
// 	cpuMetrics := RegisterCpuMetrics()

// 	// Create a new registry to register metrics
// 	reg := prometheus.NewRegistry()
// 	reg.MustRegister(cpuMetrics.CpuUsage)
// 	reg.MustRegister(cpuMetrics.CpuFree)
// 	reg.MustRegister(cpuMetrics.CpuUsagePerApp)

// 	// Verify if the metrics are registered
// 	_, err := reg.Gather()
// 	if err != nil {
// 		t.Fatalf("failed to gather metrics: %v", err)
// 	}
// }

// // type mockCPUMetrics struct {
// // 	cpuUsage float64
// // 	cpuFree  float64
// // }

// // func (m *mockCPUMetrics) CpuUsage() prometheus.Gauge {
// // 	return prometheus.NewGauge(prometheus.GaugeOpts{})
// // }

// // func (m *mockCPUMetrics) CpuFree() prometheus.Gauge {
// // 	return prometheus.NewGauge(prometheus.GaugeOpts{})
// // }

// // func TestCollectCpuMetrics(t *testing.T) {
// // 	estMetrics := &CPUMetrics{
// // 		CpuUsage: NewTestGauge(),
// // 		CpuFree:  NewTestGauge(),
// // 	}

// // 	done := make(chan struct{})
// // 	defer close(done)

// // 	go cpu.CollectCpuMetrics(testMetrics, done)

// // 	// Let the function run for a short time
// // 	time.Sleep(time.Second)

// // 	// Add your assertions here to check if metrics are being updated correctly
// // 	assert.NotNil(t, testMetrics.CpuUsage)
// // 	assert.NotNil(t, testMetrics.CpuFree)

// // 	done <- struct{}{} // Stop the goroutine
// // }
