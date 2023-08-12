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

	/*********************** -  CPU Metrics - **************************/

	//Register CPU Metrics

	cpuMetrics := cpu.RegisterCpuMetrics()

	// Collecting CPU Metrics, in the background, we use here goRoutine to handle this

	go cpu.CollectCpuMetrics(cpuMetrics, done)

	//fmt.Println("** ** Starting Http server to expose cpu metrics ...")

	//go cpu.ExposeCpuMetrics() // Expose cpu Metrics via http - You can access the results on this endpoints : localhost:9091/metrics/cpu

	/*********************** -  Memory Metrics - **************************/

	//Register Memory Metrics

	memoryMetrics := memory.RegisterMemoryMetrics()

	// Collecting Memory Metrics, in the background, we use here goRoutine to handle this

	go memory.CollectMemoryMetrics(memoryMetrics, done)

	/*************************************************
	*												 *
	*					 - Disk Metrics -            *

	**************************************************/

	//Register Disk Metrics
	diskMetrics := disk.RegisterDiskMetrics()

	// Collecting Disk Metrics, int the background , we use here goRoutine
	go disk.CollectDiskMetrics(diskMetrics, done)

	fmt.Println("** ** Starting Http server to expose Metrics to localhost:9091/metrics")

	go ExposeMetrics()
	<-done

}

func ExposeMetrics() {
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":9091", nil)
}
