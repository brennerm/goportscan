package goportscan

import (
	"log"
	"net"
	"strconv"
)

type PortScanner struct {
	host string
}

const UNKNOWN = "unknown"

// returns new PortScanner instance
// if host is empty "localhost" is assumed
func NewPortScanner(host string) *PortScanner {
	if host == ""{
		host = "localhost"
	}
	return &PortScanner{host}
}

// checks for known ports and resolves services if port is open
func (ps PortScanner) ScanKnownPorts() map[int]string {
	results := make(map[int]string)

	for port, service := range KNOWN_TCP_PORTS {
		if isOpen(ps.host, port) {
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
		if isOpen(ps.host, start_port) {
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
		if isOpen(ps.host, port) {
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
func isOpen(host string, port int) bool {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", assembleEndpoint(host, port))
	if err != nil {
		log.Print(err)
		return false
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
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
