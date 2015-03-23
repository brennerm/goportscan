package goportscan

import(
	"io/ioutil"
	"log"
	"net"
	"reflect"
	"strconv"
	"testing"
	"time"
)

const TEST_HOST = "localhost"
const TEST_TIMEOUT = time.Second

func TestMain(m *testing.M){
	listener_tcp_known := setupLocalListener("tcp4", KNOWN_TEST_PORT)
	defer listener_tcp_known.Close()
	
	listener_tcp_unknown := setupLocalListener("tcp4", UNKNOWN_TEST_PORT)
	defer listener_tcp_unknown.Close()
	
	// disable log output
	log.SetOutput(ioutil.Discard)
	
	m.Run()
}

func setupLocalListener(network string, port int) net.Listener {
	listener, err := net.Listen(network, assembleEndpoint(TEST_HOST, port))
	if err != nil {
		panic(err)
	}
	return listener
}

func TestNewPortScanner(t *testing.T){
	expected := PortScanner{TEST_HOST, DEFAULT_TIMEOUT}
	actual := *NewPortScanner(TEST_HOST)
	
	if actual != expected {
		t.Errorf("Got: %q; Expected: %q", actual, expected)
	}
	
	actual = *NewPortScanner("")
	
	if actual != expected {
		t.Errorf("Got: %q; Expected: %q", actual, expected)
	}
	
	expected = PortScanner{TEST_HOST, TEST_TIMEOUT}
	actual = *NewPortScanner(TEST_HOST, TEST_TIMEOUT)
	
	if actual != expected {
		t.Errorf("Got: %q; Expected: %q", actual, expected)
	}
}

func TestScanPorts(t *testing.T){
	ps := NewPortScanner(TEST_HOST)
	
	expected := make(map[int]string)
	actual := ps.ScanPorts([]int{})
	
	if !reflect.DeepEqual(actual, expected){
		t.Errorf("Got: %q; Expected: %q", actual, expected)
	}
	
	expected = map[int]string {
  	    KNOWN_TEST_PORT: "test port",
	}
	actual = ps.ScanPorts([]int{KNOWN_TEST_PORT})
	
	if !reflect.DeepEqual(actual, expected){
		t.Errorf("Got: %q; Expected: %q", actual, expected)
	}
	
	expected = map[int]string {
  	    UNKNOWN_TEST_PORT: "unknown",
	}
	actual = ps.ScanPorts([]int{UNKNOWN_TEST_PORT})
	
	if !reflect.DeepEqual(actual, expected){
		t.Errorf("Got: %q; Expected: %q", actual, expected)
	}
	
}

func TestScanKnownPorts(t *testing.T){
	ps := NewPortScanner(TEST_HOST)
	
	expected := map[int]string {
  	    KNOWN_TEST_PORT: "test port",
	}
	actual := ps.ScanKnownPorts()
	
	// only compare our test port, cause there could be other open ports on localhost
	if actual[KNOWN_TEST_PORT] != expected[KNOWN_TEST_PORT]{
		t.Errorf("Got: %q; Expected: %q", actual, expected)
	}
}

func TestScanPortRange(t *testing.T){
	ps := NewPortScanner(TEST_HOST)
	
	expected := map[int]string {
  	    KNOWN_TEST_PORT: "test port",
		UNKNOWN_TEST_PORT: "unknown",
	}
	actual := ps.ScanPortRange(KNOWN_TEST_PORT, UNKNOWN_TEST_PORT)
	
	if !reflect.DeepEqual(actual, expected){
		t.Errorf("Got: %q; Expected: %q", actual, expected)
	}
	
	defer func() {
		if err := recover(); err != "start_port has to be smaller than end_port" {
		    t.Errorf("Got: %q; Expected: %q", err, "start_port has to be smaller than end_port")
		}
		
	}()
	
	// trigger panic
	ps.ScanPortRange(1, 0)
}

func TestIsOpen(t *testing.T){
	expected := true
	actual := isOpen(TEST_HOST, KNOWN_TEST_PORT, TEST_TIMEOUT)
	
	if actual != expected{
		t.Errorf("Got: %q; Expected: %q", actual, expected)
	}
	
	expected = false
	// port 0 is wildcard for a random free port
	actual = isOpen(TEST_HOST, 0, TEST_TIMEOUT)
	
	if actual != expected{
		t.Errorf("Got: %q; Expected: %q", actual, expected)
	}
}

func TestAssembleEndpoint(t * testing.T){
	expected := TEST_HOST + ":" + strconv.Itoa(KNOWN_TEST_PORT)
	actual := assembleEndpoint(TEST_HOST, KNOWN_TEST_PORT)
	
	if actual != expected {
		t.Errorf("Got: %q; Expected: %q", actual, expected)
	}
}
