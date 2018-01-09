package syncSocket

import (
	"fmt"
	"net"
	"io"
)

//import "net"

type SyncSocket struct {
	Host string
	Port string
	Conn net.Conn
	MacAddress string
}


func (socket *SyncSocket) Start() {
	url := socket.Host+":"+socket.Port

	conn,err:= net.Dial("tcp",url)
	if err != nil {
		fmt.Println("建立连接失败",err)
		return
	}
	defer conn.Close()
	socket.Conn = conn

	writeStopChan := make(chan interface{})
	readStopChan := make(chan interface{})
	go socket.handleWrite(writeStopChan)
	go socket.handleRead(readStopChan)
	fmt.Println(<-writeStopChan)
	fmt.Println(<-readStopChan)
}

func (socket *SyncSocket)handleWrite(writeStopChan chan interface{}) {
	_,err:= socket.Conn.Write([]byte("iOS-"+socket.MacAddress))
	if err != nil {
		writeStopChan <- err
	}
}

func (socket *SyncSocket)handleRead(readStopChan chan interface{}) {

	buf := make([]byte, 0, 4096) // big buffer
	tmp := make([]byte, 256)     // using small tmo buffer for demonstrating
	for {
		n, err := socket.Conn.Read(tmp)
		if err != nil {
			if err != io.EOF {
				fmt.Println("read error:", err)
				readStopChan <- err
			}
			break
		}
		buf = append(buf, tmp[:n]...)
	}
	fmt.Println("收到socket 数据")
	fmt.Println(string(buf))

}