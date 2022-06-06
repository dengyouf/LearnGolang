package main

import (
	"fmt"
	"net"
	"strings"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println("Error：", err)
		return
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8001")
	checkErr(err)
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		checkErr(err)

		go HandleConn(conn)
	}

}

func HandleConn(conn net.Conn) {
	defer conn.Close()

	addr := conn.RemoteAddr().String()
	fmt.Println(addr, "connetct sucessful")

	buf := make([]byte, 1024)

	for {
		n, err := conn.Read(buf)
		checkErr(err)

		fmt.Printf("接收到客户端数据：%s\n", string(buf[:n]))
		if "exit" == string(buf[:n-2]) {
			fmt.Println(addr, "exit")
			return
		}

		str := strings.ToUpper(string(buf[:n]))
		conn.Write([]byte(str))
	}
}
