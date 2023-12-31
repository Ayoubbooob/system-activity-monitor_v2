# SysMetrix Pulse Monitor

SysMetrix Pulse Monitor is a system activity monitor that collects various statistics and displays them in real-time. This project focuses on monitoring CPU, memory, and disk usage, allowing users to gain insights into their system's performance. Whether you're a system administrator, developer, or curious user, SysMetrix Pulse Monitor provides a user-friendly interface to keep track of essential metrics.

Video Preview: [Watch the Video](https://www.youtube.com/watch?v=7jQIYY9qpWs&t=140s)

## Features

- **Real-time Monitoring**: SysMetrix Pulse Monitor collects and displays system metrics in real-time with a refresh rate of 1-5 seconds.

- **CPU Metrics**: Monitor overall CPU usage, free CPU capacity, and usage per application or process.

- **Memory Metrics**: Track total memory usage, available memory, and memory usage per application or process.

- **Disk Metrics**: Keep an eye on disk devices and partitions, along with read and write rates.

## Getting Started

Follow these steps to build and run SysMetrix Pulse Monitor on your system:

1. Install Go 1.20: [Go Installation Guide](https://go.dev/doc/install)

2. Install Prometheus: [Prometheus Installation Guide](https://www.cherryservers.com/blog/install-prometheus-ubuntu)
   - Start Prometheus:
     ```
     sudo systemctl enable prometheus
     sudo systemctl start prometheus
     ```
    **Note:** To ensure proper configuration for your monitoring setup, you'll also need to replace your `prometheus.yml` configuration file located in `/etc/prometheus/` with the one provided in this repository at `system-activity-monitor_v2/etc/prometheus/`.


3. Install Grafana: [Grafana Installation Guide](https://grafana.com/grafana/download)
   - Start Grafana Server:
     ```
     sudo systemctl daemon-reload
     sudo systemctl start grafana-server
     ```
    **Note:** You need to import the SYSMETRIX PULSE MONITOR dashboard to view graphs. Import the `SYSMETRIX-PULSE-MONITOR.json` file located in the `system-activity-monitor_v2/etc/grafana` folder to your Grafana server. For detailed guidance on importing dashboards in Grafana, refer to the [Grafana Dashboard Import Guide](https://grafana.com/docs/grafana/latest/dashboards/manage-dashboards/).


4. Clone the repository:
    ```
     git clone git@github.com:Ayoubbooob/system-activity-monitor_v2.git
     ```

5. Navigate to the project directory:
    ```
    cd system-activity-monitor_v2
    ```

6. Run the monitor:
    ```
    go run main.go
    ```
    ![App running](screenshots/app_running.png)

    - Access Prometheus metrics: [http://localhost:9091/metrics](http://localhost:9091/metrics)
    - Visualize metrics in Grafana: [http://localhost:3000](http://localhost:3000)
    
  

## Screenshots

### CPU Metrics
- Real-time CPU Usage Overview
![Real-time CPU Usage Overview](screenshots/cpu_usage_overview.png)

- CPU Usage Over Time + App & Process-specific CPU Utilization
![CPU Usage Over Time](screenshots/cpu_usage_over_time.png)


### Memory Metrics
- Memory Motion: Utilization and Availability Over Time
![Memory Motion: Utilization and Availability Over Time](screenshots/memory_motion_over_time.png)

- App & Process-specific Memory Utilization
![App & Process-specific Memory Utilization](screenshots/memory_utilization_per_app.png)

### Disk Metrics
- Disk Device Distribution & Partition Overview
![Disk Device Distribution](screenshots/disk_device_distribution.png)

- Disk Read Rate per App
![Disk Read Rate per App](screenshots/disk_read_rate_per_app.png)

- Disk Write Rate per App
![Disk Write Rate per App](screenshots/disk_write_rate_per_app.png)

### Prometheus UI
- Display of collected metrics in Prometheus web interface
![Metrics exposed in prometheus](screenshots/prometheus_ui.png)

### Jenkins Pipeline
- Successful execution of the Jenkins pipeline
![Screenshot of Your Jenkins Pipeline](screenshots/jenkins_pipeline.png)

### Project GIF
- A GIF showcasing the application in action
![Project GIF](screenshots/project_gif.gif)

## Verification Checklist

- [x] `README.md` with build instructions
- [x] Builds on Linux / Linux VM
- [x] Used one of the allowed programming languages (Golang)
- [x] Monitor CPU:
  - [x] Current total usage
  - [x] Current usage per application
  - [x] Current free
- [x] Monitor Memory:
  - [x] Current total usage
  - [x] Current total free
  - [x] Current usage per application
- [x] Monitor Disk:
  - [x] Available devices
  - [x] Available partitions
  - [x] Read rate per application
  - [x] Write rate per application
- [x] Real-time stats collection
- [x] UI:
  - [ ]  Terminal-based UI
  - [x] Prometheus + Grafana

## Contribution

Contributions to SysMetrix Pulse Monitor are welcome! Whether you're an experienced developer or just starting, your input can make a big difference. Feel free to fork the repository, make your changes, and submit a pull request. Let's collaborate and enhance this project together!

## License

SysMetrix Pulse Monitor is released under the [MIT License](LICENSE).

📊 Happy Monitoring! 🚀

---

This project was created as part of the [Acceptance test for the Missing semester ](https://gist.github.com/chermehdi/ae3f10775526d44a9821d8e889df798c).


