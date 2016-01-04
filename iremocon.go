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

func (iremocon *Iremocon) Au() string {
	res := iremocon.send("au")
	return res
}

func (iremocon *Iremocon) Is(ch string) string {
	res := iremocon.send("is", ch)
	return res
}

func (iremocon *Iremocon) Ic(ch string) string {
	res := iremocon.send("ic", ch)
	return res
}

func (iremocon *Iremocon) Cc() string {
	res := iremocon.send("cc")
	return res
}

func (iremocon *Iremocon) Vr() string {
	res := iremocon.send("vr")
	return res
}

func (iremocon *Iremocon) send(command string, params ...string) string {
	param := ""
	for _, p := range params {
		param += ";" + p
	}
	fmt.Fprintf(iremocon.Conn, "*%v%v\r\n", command, param)
	res, _ := bufio.NewReader(iremocon.Conn).ReadString('\n')
	return res
}
