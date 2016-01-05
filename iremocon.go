package iremocon

import "fmt"
import "net"
import "bufio"
import "strconv"

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

func (iremocon *Iremocon) Is(ch int) string {
	res := iremocon.send("is", strconv.Itoa(ch))
	return res
}

func (iremocon *Iremocon) Ic(ch int) string {
	res := iremocon.send("ic", strconv.Itoa(ch))
	return res
}

func (iremocon *Iremocon) Cc() string {
	res := iremocon.send("cc")
	return res
}

func (iremocon *Iremocon) Tm(ch int, time int, interval int) string {
	res := iremocon.send("tm", strconv.Itoa(ch), strconv.Itoa(time), strconv.Itoa(interval))
	return res
}

func (iremocon *Iremocon) Tl() string {
	res := iremocon.send("tl")
	return res
}

func (iremocon *Iremocon) Td(timerId int) string {
	res := iremocon.send("td", strconv.Itoa(timerId))
	return res
}

func (iremocon *Iremocon) Ts(time int) string {
	res := iremocon.send("ts", strconv.Itoa(time))
	return res
}

func (iremocon *Iremocon) Tg() string {
	res := iremocon.send("tg")
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
