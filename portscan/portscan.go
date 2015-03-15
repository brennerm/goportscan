package portscan

import(
	"net"
	"fmt"
	"time"
)

type PortScanner struct {
	host string
}

var TIMEOUT time.Duration = 1 * time.Second // 5 seconds

func NewPortScanner(host string) *PortScanner{
	return &PortScanner{host}
}

func (ps PortScanner) IsOpen(port int) bool{
	tcpAddr, err := net.ResolveTCPAddr("tcp4", assembleEndpoint(ps.host, port))
	if err != nil {
		return false
	}
	
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return false
	}
	
	conn.Close()
	return true
}

func (ps PortScanner) Scan(portmask []int) map[int]string{
	results := make(map[int]string)
	
	if len(portmask) <= 0 || portmask == nil {
		for port, desc := range KNOWN_TCP_PORTS{
			if ps.IsOpen(port){
				results[port] = desc
			}
		}
	}
	
	return results
}

func assembleEndpoint(host string, port int) string{
	return fmt.Sprintf("%s:%d", host, port)
}