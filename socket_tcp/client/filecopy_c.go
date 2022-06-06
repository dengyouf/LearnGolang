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
	path := "c:\\Users\\Administrator\\Desktop\\dengyouf\\LearnGolang\\gomodProject\\socket_tcp\\client\\1.txt"

	fileInfo, err := os.Stat(path)
	checkErr(err)

	conn, err := net.Dial("tcp", "127.0.0.1:8001")
	checkErr(err)
	defer conn.Close()

	_, err = conn.Write([]byte(fileInfo.Name()))
	checkErr(err)

	buf := make([]byte, 2048)
	n, err := conn.Read(buf)
	checkErr(err)

	if string(buf[:n]) == "ok" {
		SendFile(conn, path)
	}
}

func SendFile(conn net.Conn, filepath string) {
	file, err := os.Open(filepath)
	checkErr(err)
	defer file.Close()

	buf := make([]byte, 2048)
	for {
		n, err := file.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("发送文件完成")
			} else {
				fmt.Println(err)
			}
			return
		}
		checkErr(err)
		_, err = conn.Write(buf[:n])
		checkErr(err)
	}
}
