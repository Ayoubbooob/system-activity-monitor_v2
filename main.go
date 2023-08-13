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

	/*************************************************
	*												 	*
	*					 - CPU Metrics -            	*

	**************************************************/

	cpuMetrics := cpu.RegisterCpuMetrics()     //Register CPU Metrics
	go cpu.CollectCpuMetrics(cpuMetrics, done) // Collecting CPU Metrics, in the background, we use here goRoutine to handle this

	/*************************************************
	*												 	*
	*					 - Memory Metrics  -            *

	**************************************************/

	memoryMetrics := memory.RegisterMemoryMetrics()     //Register Memory Metrics
	go memory.CollectMemoryMetrics(memoryMetrics, done) // Collecting Memory Metrics, in the background, we use here goRoutine to handle this

	/*************************************************
	*												 	*
	*					 - Disk Metrics -            	*

	**************************************************/

	diskMetrics := disk.RegisterDiskMetrics()     //Register Disk Metrics
	go disk.CollectDiskMetrics(diskMetrics, done) // Collecting Disk Metrics, int the background , we use here goRoutine

	fmt.Println("** ** Starting Http server to expose Metrics to localhost:9091/metrics")
	go ExposeMetrics()
	<-done

}

// Exposing metrics to prometheus Data source in the url localhost:9091/metrics'
func ExposeMetrics() {
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":9091", nil)
}
