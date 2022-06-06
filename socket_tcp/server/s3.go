package main

import (
	"fmt"
	"io"
	"net"
	"strings"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println("报错：", err)
		return
	}
}
func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	checkErr(err)
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		checkErr(err)
		fmt.Printf("客户端【%s】加入", conn.RemoteAddr().String())
		go HandleConn(conn)
	}
}

func HandleConn(conn net.Conn) {
	defer fmt.Printf("客户端【%s】退出", conn.RemoteAddr().String())
	defer conn.Close()
	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			if err == io.EOF {
				break
			}
		}

		fmt.Printf("服务端收到数据:%s\n", string(buf[:n]))

		str := fmt.Sprintf("服务端返回：" + strings.ToUpper(string(buf[:n])))
		conn.Write([]byte(str))
	}
}
