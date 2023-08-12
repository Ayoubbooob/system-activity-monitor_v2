package cpu

import (
	"fmt"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/process"
)

type CPUMetrics struct {
	CpuUsage       prometheus.Gauge
	CpuFree        prometheus.Gauge
	CpuUsagePerApp *prometheus.GaugeVec
}

func RegisterCpuMetrics() *CPUMetrics {
	cpuMetrics := &CPUMetrics{
		CpuUsage: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: "cpu_usage_percent",
				Help: " -- This gives overall cpu usage pourcentage.",
			}),

		CpuFree: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "cpu_free_percent",
			Help: " -- This gives overall free cpu.",
		}),

		CpuUsagePerApp: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "process_cpu_usage_percent",
				Help: " -- THus gives cpu usage per app/process",
			},
			[]string{"pid", "process_name"},
		),
	}

	prometheus.MustRegister(cpuMetrics.CpuUsage)
	prometheus.MustRegister(cpuMetrics.CpuFree)
	prometheus.MustRegister(cpuMetrics.CpuUsagePerApp)

	return cpuMetrics
}

func CollectCpuMetrics(metrics *CPUMetrics, done <-chan struct{}) {
	for {
		percent, err := cpu.Percent(time.Second, false)
		if err != nil {
			fmt.Println("Error appears while trying to get CPU overall usage: ", err)
		} else {
			metrics.CpuUsage.Set(percent[0])
			metrics.CpuFree.Set(100.0 - percent[0]) // This to calculate free cpu
		}

		processes, err := process.Processes()

		if err != nil {
			fmt.Println("Error appears while trying to get Cpu usage per app/processes")
		} else {
			for _, p := range processes {
				pid := p.Pid
				processName, _ := p.Name()
				cpuUsagePercent, _ := p.CPUPercent()
				metrics.CpuUsagePerApp.WithLabelValues(fmt.Sprintf("%d", pid), processName).Set(cpuUsagePercent)

			}
		}

		time.Sleep(time.Second * 5) // Collecting metrics in 5s refresh rate

		select {
		case <-done:
			return // Exit the goroutine when done signal is received
		default:
		}

	}
}
