package common

import (
	"net"
)

func ServerStart(port string, handle func(msg string)) (err error) {

	listenr, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return
	}
	for {
		conn, err := listenr.Accept()
		if err != nil {
			return err
		}
		res := execute(conn)
		go func() {
			handle(res)
		}()
	}
}

func execute(conn net.Conn) string {
	var builder StringBuilder
	var byt [1024]byte
	for {
		n, err := conn.Read(byt[0:])
		if err != nil {
			conn.Close()
			break
		}
		resp := byt[0:n]
		if resp[0] == '\n' {
			conn.Close()
			break
		}
		builder.Concat(string(resp))
	}
	return builder.ToString()
}
