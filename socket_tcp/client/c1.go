package main

import (
	"fmt"
	"net"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	addr := "127.0.0.1:8000"

	conn, err := net.Dial("tcp", addr)
	checkErr(err)
	defer conn.Close()

	conn.Write([]byte("Hello, Server~"))

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	checkErr(err)
	fmt.Printf("客户端接收到数据：%s, by 【%s】\n", string(buf[:n]), conn.RemoteAddr().String())
}
