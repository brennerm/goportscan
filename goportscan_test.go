package goportscan

import(
	"testing"
)

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

func TestAssembleEndpoint(t * testing.T){
	expected := "localhost:80"
	actual := assembleEndpoint("localhost", 80)
	
	if actual != expected {
		t.Errorf("Got: %q; Expected: %q", actual, expected)
	}
}
