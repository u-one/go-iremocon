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

	echo = iremocon.Is(1)
	if echo != "*is;1\r\n" {
		t.Errorf("is error: %v", echo)
	}

	echo = iremocon.Ic(1)
	if echo != "*ic;1\r\n" {
		t.Errorf("ic error: %v", echo)
	}

	echo = iremocon.Cc()
	if echo != "*cc\r\n" {
		t.Errorf("cc error: %v", echo)
	}

	echo = iremocon.Tm(1, 1577804400, 10)
	if echo != "*tm;1;1577804400;10\r\n" {
		t.Errorf("tm error: %v", echo)
	}

	echo = iremocon.Tl()
	if echo != "*tl\r\n" {
		t.Errorf("tl error: %v", echo)
	}

	echo = iremocon.Td(1)
	if echo != "*td;1\r\n" {
		t.Errorf("td error: %v", echo)
	}

	echo = iremocon.Ts(1577804400)
	if echo != "*ts;1577804400\r\n" {
		t.Errorf("ts error: %v", echo)
	}

	echo = iremocon.Tg()
	if echo != "*tg\r\n" {
		t.Errorf("tg error: %v", echo)
	}

	echo = iremocon.Vr()
	if echo != "*vr\r\n" {
		t.Errorf("vr error: %v", echo)
	}
}
