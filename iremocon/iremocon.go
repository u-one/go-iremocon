package iremocon

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
)

func Au(conn net.Conn) (string, error) {
	res, err := Send(conn, "au")
	return res, err
}

func Is(conn net.Conn, ch int) (string, error) {
	res, err := Send(conn, "is", strconv.Itoa(ch))
	return res, err
}

func Ic(conn net.Conn, ch int) (string, error) {
	res, err := Send(conn, "ic", strconv.Itoa(ch))
	return res, err
}

func Cc(conn net.Conn) (string, error) {
	res, err := Send(conn, "cc")
	return res, err
}

func Tm(conn net.Conn, ch int, time int, interval int) (string, error) {
	res, err := Send(conn, "tm", strconv.Itoa(ch), strconv.Itoa(time), strconv.Itoa(interval))
	return res, err
}

func Tl(conn net.Conn) (string, error) {
	res, err := Send(conn, "tl")
	return res, err
}

func Td(conn net.Conn, timerId int) (string, error) {
	res, err := Send(conn, "td", strconv.Itoa(timerId))
	return res, err
}

func Ts(conn net.Conn, time int) (string, error) {
	res, err := Send(conn, "ts", strconv.Itoa(time))
	return res, err
}

func Tg(conn net.Conn) (string, error) {
	res, err := Send(conn, "tg")
	return res, err
}

func Vr(conn net.Conn) (string, error) {
	res, err := Send(conn, "vr")
	return res, err
}

func Li(conn net.Conn) (string, error) {
	res, err := Send(conn, "li")
	return res, err
}

func Hu(conn net.Conn) (string, error) {
	res, err := Send(conn, "hu")
	return res, err
}

func Te(conn net.Conn) (string, error) {
	res, err := Send(conn, "te")
	return res, err
}

func Se(conn net.Conn) (string, error) {
	res, err := Send(conn, "se")
	return res, err
}

func Send(conn net.Conn, command string, params ...string) (string, error) {
	param := ""
	for _, p := range params {
		param += ";" + p
	}
	fmt.Fprintf(conn, "*%v%v\r\n", command, param)
	res, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		return res, err
	}
	if strings.HasPrefix(res, ("*" + command + ";err;")) {
		return res, errors.New(res)
	}
	return res, err
}
