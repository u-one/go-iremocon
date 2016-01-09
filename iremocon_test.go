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

	var echo string
	var err error
	echo, err = Au(conn)
	if err != nil {
		t.Errorf("got error: %v", err)
	}
	if echo != "*au\r\n" {
		t.Errorf("au error: %v", echo)
	}

	echo, err = Is(conn, 1)
	if err != nil {
		t.Errorf("got error: %v", err)
	}
	if echo != "*is;1\r\n" {
		t.Errorf("is error: %v", echo)
	}

	echo, err = Ic(conn, 1)
	if err != nil {
		t.Errorf("got error: %v", err)
	}
	if echo != "*ic;1\r\n" {
		t.Errorf("ic error: %v", echo)
	}

	echo, err = Cc(conn)
	if err != nil {
		t.Errorf("got error: %v", err)
	}
	if echo != "*cc\r\n" {
		t.Errorf("cc error: %v", echo)
	}

	echo, err = Tm(conn, 1, 1577804400, 10)
	if err != nil {
		t.Errorf("got error: %v", err)
	}
	if echo != "*tm;1;1577804400;10\r\n" {
		t.Errorf("tm error: %v", echo)
	}

	echo, err = Tl(conn)
	if err != nil {
		t.Errorf("got error: %v", err)
	}
	if echo != "*tl\r\n" {
		t.Errorf("tl error: %v", echo)
	}

	echo, err = Td(conn, 1)
	if err != nil {
		t.Errorf("got error: %v", err)
	}
	if echo != "*td;1\r\n" {
		t.Errorf("td error: %v", echo)
	}

	echo, err = Ts(conn, 1577804400)
	if err != nil {
		t.Errorf("got error: %v", err)
	}
	if echo != "*ts;1577804400\r\n" {
		t.Errorf("ts error: %v", echo)
	}

	echo, err = Tg(conn)
	if err != nil {
		t.Errorf("got error: %v", err)
	}
	if echo != "*tg\r\n" {
		t.Errorf("tg error: %v", echo)
	}

	echo, err = Vr(conn)
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

	var res string
	var err error
	res, err = Send(conn, "au")
	if err == nil {
		t.Errorf("no error: %v", res)
	}
}
