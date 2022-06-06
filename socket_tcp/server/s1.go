package main

import (
	"fmt"
	"net"
	"strings"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}
func main() {
	addr := "127.0.0.1:8000"
	listener, err := net.Listen("tcp", addr)
	checkErr(err)
	defer listener.Close()

	conn, err := listener.Accept()
	checkErr(err)
	defer conn.Close()

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	checkErr(err)
	fmt.Printf("服务端接收到数据：%s, by 【%s】\n", string(buf[:n]), conn.RemoteAddr().String())

	str := strings.ToUpper(string(buf[:n]))
	conn.Write([]byte(str))

}
