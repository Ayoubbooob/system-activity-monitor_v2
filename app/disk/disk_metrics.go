package disk

// import (
// 	"fmt"
// 	"strconv"
// 	"time"

// 	"github.com/prometheus/client_golang/prometheus"
// 	"github.com/shirou/gopsutil/disk"
// 	"github.com/shirou/gopsutil/process"
// )

// type DISKMetrics struct {
// 	// 	diskTotal       *prometheus.CounterVec
// 	// 	diskUsed        *prometheus.GaugeVec
// 	// 	diskFree        *prometheus.GaugeVec
// 	devices         *prometheus.GaugeVec
// 	partitions      *prometheus.GaugeVec
// 	readRatePerApp  *prometheus.GaugeVec
// 	writeRatePerApp *prometheus.GaugeVec
// }

// func RegisterDiskMetrics() *DISKMetrics {
// 	diskMetrics := &DISKMetrics{

// 		// diskTotal: prometheus.NewCounterVec(prometheus.CounterOpts{
// 		// 	Name: "node_disk_total",
// 		// 	Help: " -- This metric gives the total disk space.",
// 		// }, []string{"deviceName"}),

// 		// diskUsed: prometheus.NewGaugeVec(prometheus.GaugeOpts{
// 		// 	Name: "node_disk_usage_gegabytes",
// 		// 	Help: " -- This metric represents the total disk usage",
// 		// }, []string{"deviceName"}),

// 		// diskFree: prometheus.NewGaugeVec(prometheus.GaugeOpts{
// 		// 	Name: "node_disk_free_gegabytes",
// 		// 	Help: " -- This metric represents the amount of free disk available",
// 		// }, []string{"deviceName"}),
// 		devices: prometheus.NewGaugeVec(prometheus.GaugeOpts{
// 			Name: "node_disk_device_devices",
// 			Help: " -- This metric gives informations about available disk devices",
// 		}, []string{"deviceName"}),

// 		partitions: prometheus.NewGaugeVec(prometheus.GaugeOpts{
// 			Name: "node_disk_partitions",
// 			Help: " -- This metric gives informations about available disk partitions",
// 		}, []string{"deviceName", "mountpoint", "fstype"}),

// 		readRatePerApp: prometheus.NewGaugeVec(prometheus.GaugeOpts{
// 			Name: "node_disk_read_bytes_rate_per_app",
// 			Help: " -- This metric gives { ** READ RATE ** }of bytes from disk per application",
// 		}, []string{"device", "pid", "process_name", "read_rate"}),

// 		writeRatePerApp: prometheus.NewGaugeVec(prometheus.GaugeOpts{
// 			Name: "node_disk_write_bytes_rate_per_app",
// 			Help: " -- This metric gives { ** WRITE RATE ** } of bytes from disk per application",
// 		}, []string{"device", "pid", "process_name", "write_rate"}),
// 	}

// 	// prometheus.MustRegister(diskMetrics.diskTotal)
// 	// prometheus.MustRegister(diskMetrics.diskUsed)
// 	// prometheus.MustRegister(diskMetrics.diskFree)
// 	prometheus.MustRegister(diskMetrics.devices)
// 	prometheus.MustRegister(diskMetrics.partitions)
// 	prometheus.MustRegister(diskMetrics.readRatePerApp)
// 	prometheus.MustRegister(diskMetrics.writeRatePerApp)

// 	return diskMetrics
// }

// func CollectDiskMetrics(diskMetrics *DISKMetrics, done <-chan struct{}) {

// 	for {

// 		// diskStorageInfo, err := disk.Partitions(false)
// 		// if err != nil {
// 		// 	fmt.Println("Error while trying to get disks storage info:", err)
// 		// } else {
// 		// 	for _, device := range diskStorageInfo {
// 		// 		deviceName := device.Device
// 		// 		usage, err := disk.Usage(device.Mountpoint)
// 		// 		if err != nil {
// 		// 			fmt.Println("Error while trying to get disk usage:", err)
// 		// 		} else {

// 		// 			totalGB := float64(usage.Total) / (1024 * 1024 * 1024)
// 		// 			usedGB := float64(usage.Used) / (1024 * 1024 * 1024)
// 		// 			freeGB := float64(usage.Free) / (1024 * 1024 * 1024)

// 		// 			diskMetrics.diskTotal.WithLabelValues(deviceName).Add(totalGB)
// 		// 			diskMetrics.diskUsed.WithLabelValues(deviceName).Set(usedGB)
// 		// 			diskMetrics.diskFree.WithLabelValues(deviceName).Set(freeGB)

// 		// 		}
// 		// 	}
// 		// }

// 		devices, err := disk.IOCounters()

// 		if err != nil {
// 			fmt.Println("Error while trying to get disks devices:", err)
// 		} else {
// 			i := 0.0
// 			for device := range devices {
// 				i = i + 1
// 				deviceName := device
// 				diskMetrics.devices.WithLabelValues(deviceName).Set(i) //this will give {....} device size -- I need to set the real size of device
// 			}

// 		}

// 		partitions, err := disk.Partitions(true)

// 		i := 0.0
// 		if err != nil {
// 			fmt.Println("Error while trying to get disks partitions:", err)
// 		} else {
// 			for _, partition := range partitions {
// 				deviceName := partition.Device
// 				mountpoint := partition.Mountpoint
// 				fstype := partition.Fstype
// 				// size, _ := disk.Usage(mountpoint)
// 				// sizeInGB := float64(size.Total) / (1024 * 1024 * 1024)
// 				i = i + 1
// 				diskMetrics.partitions.WithLabelValues(deviceName, mountpoint, fstype).Set(i)
// 			}
// 		}

// 		diskStats, err := disk.IOCounters()
// 		processes, _ := process.Processes()

// 		if err != nil {
// 			fmt.Println("Error While trying to get disk write & read rates:", err)
// 		} else {
// 			for device := range diskStats {
// 				deviceName := device
// 				for _, process := range processes {
// 					pid := process.Pid
// 					processName, _ := process.Name()
// 					pIO, err := process.IOCounters()
// 					if err != nil {
// 						fmt.Println("Error while trying to get IO counters for the process:", processName, pid, err)
// 					} else {
// 						readRate := float64(pIO.ReadBytes) / (1024 * 1024) // in MB
// 						writeRate := float64(pIO.WriteCount) / (1024 * 1024)
// 						diskMetrics.readRatePerApp.WithLabelValues(deviceName, fmt.Sprintf("%d", pid), processName, strconv.FormatFloat(readRate, 'f', 1, 64)).Set(readRate)
// 						diskMetrics.writeRatePerApp.WithLabelValues(deviceName, fmt.Sprintf("%d", pid), processName, strconv.FormatFloat(writeRate, 'f', 1, 64)).Set(writeRate)
// 					}
// 				}
// 			}
// 		}

// 		time.Sleep(time.Second * 5)

// 		select {
// 		case <-done:
// 			return // Exit the goroutine when done signal is received
// 		default:
// 		}

// 	}

// }
