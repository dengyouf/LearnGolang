package main

import (
	"fmt"
	"net"
	"os"
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

	conn.Write([]byte(os.Args[1]))

	// buf := make([]byte, 1024)
	// n, err := conn.Read(buf)
	// checkErr(err)
	// fmt.Printf("客户端接收到数据：%s", string(buf[:n]))
	read(conn)
}

func read(conn net.Conn) {
	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	checkErr(err)
	// recstr, err := simplifiedchinese.GBK.NewEncoder().Bytes([]byte(buf[:n]))
	// checkErr(err)
	fmt.Printf("客户端接收到数据：%s", string(buf[:n]))
}
