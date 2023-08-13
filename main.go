package main

import (
	"fmt"
	"net/http"

	"github.com/Ayoubbooob/system-activity-monitor_v2/app/cpu"
	"github.com/Ayoubbooob/system-activity-monitor_v2/app/disk"
	"github.com/Ayoubbooob/system-activity-monitor_v2/app/memory"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	done := make(chan struct{})

	fmt.Println(`
████████╗██╗  ██╗██╗███████╗    ██████╗ ██╗████████╗███████╗
╚══██╔══╝██║  ██║██║██╔════╝    ██╔══██╗██║╚══██╔══╝██╔════╝
    ██║   ███████║██║███████╗    ██████╔╝██║   ██║   ███████╗
    ██║   ██╔══██║██║╚════██║    ██╔══██╗██║   ██║   ╚════██║
    ██║   ██║  ██║██║███████║    ██████╔╝██║   ██║   ███████║
    ╚═╝   ╚═╝  ╚═╝╚═╝╚══════╝    ╚═════╝ ╚═╝   ╚═╝   ╚══════╝

     ██████╗ █████╗ ███╗   ███╗███████╗    ██████╗ ██████╗ ██████╗
    ██╔════╝██╔══██╗████╗ ████║██╔════╝    ██╔══██╗██╔══██╗██╔══██╗
    ██║     ███████║██╔████╔██║█████╗      ██████╔╝██████╔╝██║  ██║
    ██║     ██╔══██║██║╚██╔╝██║██╔══╝      ██╔═══╝ ██╔══██╗██║  ██║
    ╚██████╗██║  ██║██║ ╚═╝ ██║███████╗    ██║     ██║  ██║██████╔╝
     ╚═════╝╚═╝  ╚═╝╚═╝     ╚═╝╚══════╝    ╚═╝     ╚═╝  ╚═╝╚═════╝

** ** Initiating System Monitoring and Analytics ** **

            -> To access metrics in Prometheus, visit: (http://localhost:9091/metrics)
            -> For visualizing metrics using Grafana, explore: (http://localhost:3000)

            * Note: Ensure Prometheus and Grafana servers are configured to unlock the full potential.
`)

	cpuMetrics := cpu.RegisterCpuMetrics()
	go cpu.CollectCpuMetrics(cpuMetrics, done)

	memoryMetrics := memory.RegisterMemoryMetrics()
	go memory.CollectMemoryMetrics(memoryMetrics, done)

	diskMetrics := disk.RegisterDiskMetrics()
	go disk.CollectDiskMetrics(diskMetrics, done)

	go ExposeMetrics()

	<-done
}

// Exposing metrics to prometheus Data source in the url localhost:9091/metrics'
func ExposeMetrics() {
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":9091", nil)
}
