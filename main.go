package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ticker := time.NewTicker(5 * time.Minute) // Check system status every 5 minutes
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				checkSystemStatus()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	// Wait for termination signal (e.g., CTRL+C) to gracefully stop the program
	waitForTerminationSignal()
	close(quit)
}

func checkSystemStatus() {
	fmt.Println("===== System Status Report =====")

	checkUptime()
	checkLoadAverage()
	checkDiskUsage("/")
	checkMemoryUsage()
	checkNetworkStatus()
	// Add more checks here as needed

	fmt.Println("================================")
}

func checkUptime() {
	cmd := exec.Command("uptime")
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Failed to check uptime: %v\n", err)
		return
	}

	uptime := string(output)
	fmt.Printf("System Uptime:\n%s\n", uptime)
}

func checkLoadAverage() {
	cmd := exec.Command("cat", "/proc/loadavg")
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Failed to check load average: %v\n", err)
		return
	}

	loadAverage := string(output)
	fmt.Printf("Load Average:\n%s\n", loadAverage)
}

func checkDiskUsage(path string) {
	cmd := exec.Command("df", "-h", path)
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Failed to check disk usage for %s: %v\n", path, err)
		return
	}

	diskUsage := string(output)
	fmt.Printf("Disk Usage for %s:\n%s\n", path, diskUsage)
}

func checkMemoryUsage() {
	cmd := exec.Command("free", "-h")
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Failed to check memory usage: %v\n", err)
		return
	}

	memoryUsage := string(output)
	fmt.Printf("Memory Usage:\n%s\n", memoryUsage)
}

func checkNetworkStatus() {
	cmd := exec.Command("ifconfig")
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Failed to check network status: %v\n", err)
		return
	}

	networkStatus := string(output)
	fmt.Printf("Network Status:\n%s\n", networkStatus)
}

func waitForTerminationSignal() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
}
