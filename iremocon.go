package iremocon

import "fmt"
import "net"
import "bufio"
import "strings"

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
	res := iremocon.send("vr")
	return res
}

func (iremocon *Iremocon) send(command string, params ...string) string {
	param := strings.Join(params, ";")
	fmt.Fprintf(iremocon.Conn, "*%v%v\r\n", command, param)
	res, _ := bufio.NewReader(iremocon.Conn).ReadString('\n')
	return res
}
