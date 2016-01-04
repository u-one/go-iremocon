package iremocon

import "net"
import "testing"

import "io"

func runEchoServer(listener net.Listener) {
	for {
		conn, _ := listener.Accept()
		go func(c net.Conn) {
			io.Copy(c, c)
			c.Close()
		}(conn)
	}
}

func TestIremocon(t *testing.T) {
	laddr := ":50000"
	ln, _ := net.Listen("tcp", laddr)
	defer ln.Close()
	go runEchoServer(ln)

	conn, _ := net.Dial("tcp", laddr)
	iremocon := NewIremocon(conn)
	echo := ""
	echo = iremocon.Au()
	if echo != "*au\r\n" {
		t.Errorf("au error: %v", echo)
	}

	echo = iremocon.Is("1")
	if echo != "*is;1\r\n" {
		t.Errorf("is error: %v", echo)
	}

	echo = iremocon.Ic("1")
	if echo != "*ic;1\r\n" {
		t.Errorf("ic error: %v", echo)
	}

	echo = iremocon.Cc()
	if echo != "*cc\r\n" {
		t.Errorf("cc error: %v", echo)
	}

	echo = iremocon.Vr()
	if echo != "*vr\r\n" {
		t.Errorf("vr error: %v", echo)
	}
}
