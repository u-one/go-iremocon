package iremocon

import "fmt"
import "net"
import "testing"

func TestIremocon(t *testing.T) {
	conn, err := net.Dial("tcp", "10.0.1.200:51013")
	if err != nil {
		fmt.Printf("error")
	}

	iremocon := NewIremocon(conn)
	result := iremocon.Vr()
	if result != "1.0.0\r\n" {
		t.Fatalf("vr error: %v", result)
	} else {
		t.Log(result)
	}
}
