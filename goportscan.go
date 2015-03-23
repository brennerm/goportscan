package goportscan

import (
	"log"
	"net"
	"strconv"
	"time"
)

type PortScanner struct {
	host string
	timeout time.Duration
}

const UNKNOWN = "unknown"
const DEFAULT_TIMEOUT = 500 * time.Millisecond

// returns new PortScanner instance
// if host is empty "localhost" is assumed
func NewPortScanner(host string, given_timeout ...time.Duration) *PortScanner {
	timeout := DEFAULT_TIMEOUT
	if len(given_timeout) > 0 {
		timeout = given_timeout[0]
	}
	
	if host == ""{
		host = "localhost"
	}
	return &PortScanner{host, timeout}
}

// checks for known ports and resolves services if port is open
func (ps PortScanner) ScanKnownPorts() map[int]string {
	results := make(map[int]string)

	for port, service := range KNOWN_TCP_PORTS {
		if isOpen(ps.host, port, ps.timeout) {
			results[port] = service
		}
	}
	return results
}

// checks ports within range
// resolves service if open port is known
func (ps PortScanner) ScanPortRange(start_port int, end_port int) map[int]string {
	if start_port > end_port {
		log.Panic("start_port has to be smaller than end_port")
	}
	results := make(map[int]string)

	for start_port <= end_port {
		if isOpen(ps.host, start_port, ps.timeout) {
			if val, ok := KNOWN_TCP_PORTS[start_port]; ok {
				results[start_port] = val
			} else {
				results[start_port] = UNKNOWN
			}
		}
		start_port++
	}
	return results
}

// checks given ports
// resolves service if open port is known
func (ps PortScanner) ScanPorts(ports []int) map[int]string {
	results := make(map[int]string)
	
	if len(ports) == 0 {
		return results
	}

	for _, port := range ports {
		if isOpen(ps.host, port, ps.timeout) {
			if val, ok := KNOWN_TCP_PORTS[port]; ok {
				results[port] = val
			} else {
				results[port] = UNKNOWN
			}
		}
	}
	return results
}

// checks if given host is reachable and port is open
func isOpen(host string, port int, timeout time.Duration) bool {
	conn, err := net.DialTimeout("tcp", assembleEndpoint(host, port), timeout)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

// concats host ip and port
func assembleEndpoint(host string, port int) string {
	return host + ":" + strconv.Itoa(port)
}
