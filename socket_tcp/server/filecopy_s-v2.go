package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println("报错信息：", err)
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
		defer conn.Close()

		buf := make([]byte, 4096)
		n, err := conn.Read(buf)
		checkErr(err)

		fileName := string(buf[:n])
		conn.Write([]byte("ok"))

		go recvFile(conn, fileName)
	}
}

func recvFile(conn net.Conn, fileName string) {
	pathDir := "D:\\testdir\\"

	file, err := os.Create(pathDir + fileName)
	checkErr(err)

	defer file.Close()

	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("读取完了")
				return
			}
			fmt.Println("读取Err")

		}

		_, err = file.Write(buf[:n])

	}

}
