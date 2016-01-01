package iremocon

import "fmt"
import "net"
import "bufio"

type Iremocon struct {
	Conn net.Conn
}

func NewIremocon(conn net.Conn) *Iremocon {
	i := &Iremocon{
		Conn: conn,
	}
	return i
}

func (iremocon *Iremocon) Vr() string {
	fmt.Fprintf(iremocon.Conn, "*%v\r\n", "vr")
	res, _ := bufio.NewReader(iremocon.Conn).ReadString('\n')
	return res
}
