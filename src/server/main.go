package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"server/rrd"
)

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", ":7777") //获取一个tcpAddr
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr) //监听一个端口
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	result, err := ioutil.ReadAll(conn)
	checkError(err)
	go rrd.Handler(string(result))
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
