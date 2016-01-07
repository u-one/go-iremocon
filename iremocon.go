package iremocon

import "fmt"
import "net"
import "bufio"
import "strconv"
import "strings"
import "errors"

type Iremocon struct {
	Conn net.Conn
}

func NewIremocon(conn net.Conn) *Iremocon {
	i := &Iremocon{
		Conn: conn,
	}
	return i
}

func (iremocon *Iremocon) Au() (string, error) {
	res, err := iremocon.send("au")
	return res, err
}

func (iremocon *Iremocon) Is(ch int) (string, error) {
	res, err := iremocon.send("is", strconv.Itoa(ch))
	return res, err
}

func (iremocon *Iremocon) Ic(ch int) (string, error) {
	res, err := iremocon.send("ic", strconv.Itoa(ch))
	return res, err
}

func (iremocon *Iremocon) Cc() (string, error) {
	res, err := iremocon.send("cc")
	return res, err
}

func (iremocon *Iremocon) Tm(ch int, time int, interval int) (string, error) {
	res, err := iremocon.send("tm", strconv.Itoa(ch), strconv.Itoa(time), strconv.Itoa(interval))
	return res, err
}

func (iremocon *Iremocon) Tl() (string, error) {
	res, err := iremocon.send("tl")
	return res, err
}

func (iremocon *Iremocon) Td(timerId int) (string, error) {
	res, err := iremocon.send("td", strconv.Itoa(timerId))
	return res, err
}

func (iremocon *Iremocon) Ts(time int) (string, error) {
	res, err := iremocon.send("ts", strconv.Itoa(time))
	return res, err
}

func (iremocon *Iremocon) Tg() (string, error) {
	res, err := iremocon.send("tg")
	return res, err
}

func (iremocon *Iremocon) Vr() (string, error) {
	res, err := iremocon.send("vr")
	return res, err
}

func (iremocon *Iremocon) send(command string, params ...string) (string, error) {
	param := ""
	for _, p := range params {
		param += ";" + p
	}
	fmt.Fprintf(iremocon.Conn, "*%v%v\r\n", command, param)
	res, err := bufio.NewReader(iremocon.Conn).ReadString('\n')
	if err != nil {
		return res, err
	}
	if strings.HasPrefix(res, ("*" + command + ";err;")) {
		return res, errors.New(res)
	}
	return res, err
}
