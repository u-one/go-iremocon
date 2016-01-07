package iremocon

import "net"
import "testing"
import "bufio"
import "io"

const laddr = ":50000"

func runEchoServer(listener net.Listener, t *testing.T) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			t.Fatal(err)
		}
		go func(c net.Conn) {
			io.Copy(c, c)
			c.Close()
		}(conn)
	}
}

func runErrorServer(listener net.Listener, t *testing.T) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			t.Fatal(err)
		}
		go func(c net.Conn) {
			r := bufio.NewReader(c)
			for {
				req, rerr := r.ReadString('\n')
				if rerr != nil {
					break
				}
				res := req[0:3] + ";err;000\r\n"
				c.Write([]byte(res))
			}
			c.Close()
		}(conn)
	}
}

func TestIremocon(t *testing.T) {
	ln, _ := net.Listen("tcp", laddr)
	defer ln.Close()
	go runEchoServer(ln, t)

	conn, _ := net.Dial("tcp", laddr)
	defer conn.Close()
	iremocon := NewIremocon(conn)

	var echo string
	var err error
	echo, err = iremocon.Au()
	if err != nil {
		t.Errorf("got error: %v", err)
	}
	if echo != "*au\r\n" {
		t.Errorf("au error: %v", echo)
	}

	echo, err = iremocon.Is(1)
	if err != nil {
		t.Errorf("got error: %v", err)
	}
	if echo != "*is;1\r\n" {
		t.Errorf("is error: %v", echo)
	}

	echo, err = iremocon.Ic(1)
	if err != nil {
		t.Errorf("got error: %v", err)
	}
	if echo != "*ic;1\r\n" {
		t.Errorf("ic error: %v", echo)
	}

	echo, err = iremocon.Cc()
	if err != nil {
		t.Errorf("got error: %v", err)
	}
	if echo != "*cc\r\n" {
		t.Errorf("cc error: %v", echo)
	}

	echo, err = iremocon.Tm(1, 1577804400, 10)
	if err != nil {
		t.Errorf("got error: %v", err)
	}
	if echo != "*tm;1;1577804400;10\r\n" {
		t.Errorf("tm error: %v", echo)
	}

	echo, err = iremocon.Tl()
	if err != nil {
		t.Errorf("got error: %v", err)
	}
	if echo != "*tl\r\n" {
		t.Errorf("tl error: %v", echo)
	}

	echo, err = iremocon.Td(1)
	if err != nil {
		t.Errorf("got error: %v", err)
	}
	if echo != "*td;1\r\n" {
		t.Errorf("td error: %v", echo)
	}

	echo, err = iremocon.Ts(1577804400)
	if err != nil {
		t.Errorf("got error: %v", err)
	}
	if echo != "*ts;1577804400\r\n" {
		t.Errorf("ts error: %v", echo)
	}

	echo, err = iremocon.Tg()
	if err != nil {
		t.Errorf("got error: %v", err)
	}
	if echo != "*tg\r\n" {
		t.Errorf("tg error: %v", echo)
	}

	echo, err = iremocon.Vr()
	if err != nil {
		t.Errorf("got error: %v", err)
	}
	if echo != "*vr\r\n" {
		t.Errorf("vr error: %v", echo)
	}
}

func TestIremoconError(t *testing.T) {
	ln, _ := net.Listen("tcp", laddr)
	defer ln.Close()
	go runErrorServer(ln, t)

	conn, _ := net.Dial("tcp", laddr)
	defer conn.Close()
	iremocon := NewIremocon(conn)

	var res string
	var err error
	res, err = iremocon.Au()
	if err == nil {
		t.Errorf("no error: %v", res)
	}

	res, err = iremocon.Is(1)
	if err == nil {
		t.Errorf("no error: %v", res)
	}

	res, err = iremocon.Ic(1)
	if err == nil {
		t.Errorf("no error: %v", res)
	}

	res, err = iremocon.Cc()
	if err == nil {
		t.Errorf("no error: %v", res)
	}

	res, err = iremocon.Tm(1, 1577804400, 10)
	if err == nil {
		t.Errorf("no error: %v", res)
	}

	res, err = iremocon.Tl()
	if err == nil {
		t.Errorf("no error: %v", res)
	}

	res, err = iremocon.Td(1)
	if err == nil {
		t.Errorf("no error: %v", res)
	}

	res, err = iremocon.Ts(1577804400)
	if err == nil {
		t.Errorf("no error: %v", res)
	}

	res, err = iremocon.Tg()
	if err == nil {
		t.Errorf("no error: %v", res)
	}

	res, err = iremocon.Vr()
	if err == nil {
		t.Errorf("no error: %v", res)
	}
}
