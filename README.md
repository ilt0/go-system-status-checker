# System Status Checker

This program provides a simple system status checker that periodically checks the system status and reports various metrics such as uptime, load average, disk usage, memory usage, and network status. It runs as a continuous process and provides real-time information about the system.

## Prerequisites

Before running this code, make sure you have Go installed on your system. You can download and install Go from the official Go website (https://golang.org).

## Usage

1. Clone the repository or create a new Go project.
2. Create a new Go file (e.g., `main.go`) and copy the code into it.
3. Open a terminal or command prompt and navigate to the project directory.
4. Run the code using the following command:

```
    go run main.go
```

The system status checker will start running and display the system status report every 5 minutes.

## How It Works

The code uses Go's built-in `time` package to schedule the system status checks every 5 minutes. It sets up a ticker that triggers the `checkSystemStatus` function at regular intervals. The `checkSystemStatus` function calls various helper functions to collect and display different system metrics.

The program also waits for a termination signal (e.g., CTRL+C) to gracefully stop the program. When the termination signal is received, it closes the `quit` channel, which stops the ticker and terminates the program.

## Functions

### `main`

The `main` function is the entry point of the program. It creates a ticker to schedule system status checks, starts a goroutine to perform the checks, and waits for a termination signal.

### `checkSystemStatus`

The `checkSystemStatus` function is responsible for collecting and displaying the system status. It calls several helper functions to retrieve different metrics and prints the results to the console.

### `checkUptime`

The `checkUptime` function uses the `uptime` command to retrieve the system uptime and prints it to the console.

### `checkLoadAverage`

The `checkLoadAverage` function reads the load average from the `/proc/loadavg` file and prints it to the console.

### `checkDiskUsage`

The `checkDiskUsage` function uses the `df` command to check the disk usage for a specified path and prints the result to the console.

### `checkMemoryUsage`

The `checkMemoryUsage` function uses the `free` command to check the memory usage and prints the result to the console.

### `checkNetworkStatus`

The `checkNetworkStatus` function uses the `ifconfig` command to retrieve the network status and prints it to the console.

### `waitForTerminationSignal`

The `waitForTerminationSignal` function waits for a termination signal (e.g., CTRL+C) and terminates the program gracefully when the signal is received.

## Sample Output:

```yaml
    ===== System Status Report =====
    System Uptime:
     21:37:29 up 4 days, 10:23,  3 users,  load average: 0.42, 0.54, 0.60
    
    Load Average:
    0.42 0.54 0.60 1/516 1234
    
    Disk Usage for /:
    Filesystem      Size  Used Avail Use% Mounted on
    /dev/sda1       100G   50G   50G  50% /
    
    Memory Usage:
                  total        used        free      shared  buff/cache   available
    Mem:           7.7G        3.2G        2.1G        512M        2.4G        4.5G
    Swap:          2.0G        100M        1.9G
    
    Network Status:
    eth0: flags=4163<UP,BROADCAST,RUNNING,MULTICAST>  mtu 1500
            inet 192.168.0.10  netmask 255.255.255.0  broadcast 192.168.0.255
            inet6 fe80::c123:4567:89ab:cdef  prefixlen 64  scopeid 0x20<link>
            ether 01:23:45:67:89:ab  txqueuelen 1000  (Ethernet)
            RX packets 51838213  bytes 5046749248 (4.7 GiB)
            RX errors 0  dropped 0  overruns 0  frame 0
            TX packets 12345678  bytes 1234567890 (1.1 GiB)
            TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0
    
    ================================

```

## Customization

You can customize the system status checks by modifying the `checkSystemStatus` function. You can add or remove checks as needed by calling the appropriate helper functions within the `checkSystemStatus` function.

For example, if you want to add a check for CPU temperature, you can create a new helper function (`checkCPUTemperature`), implement the logic to retrieve the temperature, and call it within the `checkSystemStatus` function.

## Disclaimer

This code is provided as-is without any warranty. Use it at your own risk. The author is not responsible for any damage or loss caused by the use of this code.



