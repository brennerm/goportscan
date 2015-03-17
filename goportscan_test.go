package goportscan

import(
	"net"
	"reflect"
	"testing"
)

func setupLocalListener(t *testing.T, port int) net.Listener {
	listener, err := net.Listen("tcp4", assembleEndpoint("localhost", port))
	if err != nil {
		t.Fatal(err)
	}
	return listener
}

func TestNewPortScanner(t *testing.T){
	expected := PortScanner{"localhost"}
	actual := *NewPortScanner("localhost")
	
	if actual != expected {
		t.Errorf("Got: %q; Expected: %q", actual, expected)
	}
	
	actual = *NewPortScanner("")
	
	if actual != expected {
		t.Errorf("Got: %q; Expected: %q", actual, expected)
	}
}

func TestScanPorts(t *testing.T){
	ps := NewPortScanner("localhost")
	
	expected := make(map[int]string)
	actual := ps.ScanPorts([]int{})
	
	if !reflect.DeepEqual(actual, expected){
		t.Errorf("Got: %q; Expected: %q", actual, expected)
	}
	
	listener_known := setupLocalListener(t, KNOWN_TEST_PORT)
	defer listener_known.Close()
	
	expected = map[int]string {
  	    KNOWN_TEST_PORT: "test port",
	}
	actual = ps.ScanPorts([]int{KNOWN_TEST_PORT})
	
	if !reflect.DeepEqual(actual, expected){
		t.Errorf("Got: %q; Expected: %q", actual, expected)
	}
	
	listener_unknown := setupLocalListener(t, UNKNOWN_TEST_PORT)
	defer listener_unknown.Close()
	
	expected = map[int]string {
  	    UNKNOWN_TEST_PORT: "unknown",
	}
	actual = ps.ScanPorts([]int{UNKNOWN_TEST_PORT})
	
	if !reflect.DeepEqual(actual, expected){
		t.Errorf("Got: %q; Expected: %q", actual, expected)
	}
	
}

func TestIsOpen(t *testing.T){
	listener := setupLocalListener(t, KNOWN_TEST_PORT)
	defer listener.Close()
	
	expected := true
	actual := isOpen("localhost", KNOWN_TEST_PORT)
	
	if actual != expected{
		t.Errorf("Got: %q; Expected: %q", actual, expected)
	}
	
	expected = false
	// port 0 is wildcard for a random free port
	actual = isOpen("localhost", 0)
	
	if actual != expected{
		t.Errorf("Got: %q; Expected: %q", actual, expected)
	}
}

func TestAssembleEndpoint(t * testing.T){
	expected := "localhost:80"
	actual := assembleEndpoint("localhost", 80)
	
	if actual != expected {
		t.Errorf("Got: %q; Expected: %q", actual, expected)
	}
}
