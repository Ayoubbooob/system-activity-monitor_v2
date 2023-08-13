package disk

import (
	"fmt"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/process"
)

type DISKMetrics struct {
	// 	diskTotal       *prometheus.CounterVec
	// 	diskUsed        *prometheus.GaugeVec	//TODO - adding more interesting metrics
	// 	diskFree        *prometheus.GaugeVec
	devices         *prometheus.GaugeVec
	partitions      *prometheus.GaugeVec
	readRatePerApp  *prometheus.GaugeVec
	writeRatePerApp *prometheus.GaugeVec
}

func RegisterDiskMetrics() *DISKMetrics {
	diskMetrics := &DISKMetrics{

		devices: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "node_disk_device_devices",
			Help: " -- This metric gives informations about available disk devices",
		}, []string{"deviceName"}),

		partitions: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "node_disk_partitions",
			Help: " -- This metric gives informations about available disk partitions",
		}, []string{"deviceName", "mountpoint", "fstype"}),

		readRatePerApp: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "node_disk_read_bytes_rate_per_app",
			Help: " -- This metric gives { ** READ RATE ** }of bytes from disk per application",
		}, []string{"pid", "process_name", "read_rate"}),

		writeRatePerApp: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "node_disk_write_bytes_rate_per_app",
			Help: " -- This metric gives { ** WRITE RATE ** } of bytes from disk per application",
		}, []string{"pid", "process_name", "write_rate"}),
	}
	prometheus.MustRegister(diskMetrics.devices)
	prometheus.MustRegister(diskMetrics.partitions)
	prometheus.MustRegister(diskMetrics.readRatePerApp)
	prometheus.MustRegister(diskMetrics.writeRatePerApp)

	return diskMetrics
}

func CollectDiskMetrics(diskMetrics *DISKMetrics, done <-chan struct{}) {

	for {

		devices, err := disk.IOCounters()

		if err != nil {
			fmt.Println("Error while trying to get disks devices:", err)
		} else {
			i := 0.0
			for device := range devices {
				i = i + 1
				deviceName := device
				diskMetrics.devices.WithLabelValues(deviceName).Set(i)
			}

		}

		partitions, err := disk.Partitions(true)

		i := 0.0
		if err != nil {
			fmt.Println("Error while trying to get disks partitions:", err)
		} else {
			for _, partition := range partitions {
				deviceName := partition.Device
				mountpoint := partition.Mountpoint
				fstype := partition.Fstype
				i = i + 1
				diskMetrics.partitions.WithLabelValues(deviceName, mountpoint, fstype).Set(i)
			}
		}

		//diskStats, err := disk.IOCounters()
		processes, _ := process.Processes()

		if err != nil {
			fmt.Println("Error While trying to get disk write & read rates:", err)
		} else {

			//deviceName := device
			for _, myProcess := range processes {
				pid := myProcess.Pid
				processName, _ := myProcess.Name()
				pIO, err := myProcess.IOCounters()

				if err == nil {
					readRate := float64(pIO.ReadBytes) / (1024 * 1024) // in MB
					writeRate := float64(pIO.WriteCount) / (1024 * 1024)
					diskMetrics.readRatePerApp.WithLabelValues(fmt.Sprintf("%d", pid), processName, strconv.FormatFloat(readRate, 'f', 1, 64)).Set(readRate)
					diskMetrics.writeRatePerApp.WithLabelValues(fmt.Sprintf("%d", pid), processName, strconv.FormatFloat(writeRate, 'f', 1, 64)).Set(writeRate)
				}

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
