package memory

import (
	"fmt"
	"math"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/process"
)

type MEMORYMetrics struct {
	memoryTotal       prometheus.Gauge
	memoryFree        prometheus.Gauge
	memoryUsagePerApp *prometheus.GaugeVec
}

func RegisterMemoryMetrics() *MEMORYMetrics {
	memoryMetrics := &MEMORYMetrics{
		memoryTotal: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "node_memory_usage_gegabytes",
			Help: " -- This metric represents the total memory usage",
		}),

		memoryFree: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "node_memory_free_gegabytes",
			Help: " -- This metric represents the amount of free memory available",
		}),

		memoryUsagePerApp: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "process_memory_usage_gegabytes",
			Help: " -- This metric memory usage per app/process",
		}, []string{"pid", "process_name"}),
	}

	prometheus.MustRegister(memoryMetrics.memoryTotal)
	prometheus.MustRegister(memoryMetrics.memoryFree)
	prometheus.MustRegister(memoryMetrics.memoryUsagePerApp)

	return memoryMetrics
}

func CollectMemoryMetrics(metrics *MEMORYMetrics, done <-chan struct{}) {
	for {
		memoryInf, err := mem.VirtualMemory()

		if err != nil {
			fmt.Println("Error while trying to get memory usage:", err)
		} else {
			metrics.memoryTotal.Set(formatMemoryToGega((memoryInf.Total - memoryInf.Available)))
			metrics.memoryFree.Set(formatMemoryToGega(memoryInf.Available))
		}

		processes, err := process.Processes()

		if err != nil {
			fmt.Println("Error while trying to get memory usage per process")
		} else {
			for _, p := range processes {
				pid := p.Pid
				processName, _ := p.Name()
				memoryUsage, _ := p.MemoryInfo()
				metrics.memoryUsagePerApp.WithLabelValues(fmt.Sprintf("%d", pid), processName).Set(formatMemoryToMega(memoryUsage.RSS))
			}
		}

		time.Sleep(time.Second * 5)

		select {
		case <-done:
			return // Exit the goroutine when done signal is received
		default:
		}
	}

}

func formatMemoryToGega(memoryValue uint64) float64 {
	return math.Round(float64(memoryValue)/1024*1024*1024*100) / 100
}

func formatMemoryToMega(memoryValue uint64) float64 {
	return math.Round(float64(memoryValue)/1024*1024*100) / 100
}
func ExposeMemoryMetrics() {
	http.Handle("/metrics/memory", promhttp.Handler())
	http.ListenAndServe(":9092", nil)
}
