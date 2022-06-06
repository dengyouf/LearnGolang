package main

import (
	"fmt"
	"net"
	"os/exec"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
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
		defer conn.Close()

		buf := make([]byte, 4096)
		n, err := conn.Read(buf)
		checkErr(err)
		fmt.Println("收到数据", string(buf[:n]))

		cmd := exec.Command(string(buf[:n]))
		out, err := cmd.CombinedOutput()
		checkErr(err)
		conn.Write([]byte(out))
	}

}
