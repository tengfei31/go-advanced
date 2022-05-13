package network

import (
	"log"
	"net"
)

func startTcp() error {
	var addr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
	var listen, err = net.ListenTCP("tcp", addr)
	if err != nil {
		return err
	}
	log.Println("tcp listen:", addr.String())

	connMap := make(map[string]*net.TCPConn)
	for {
		tcpConn, _ := listen.AcceptTCP()

		connMap[tcpConn.RemoteAddr().String()] = tcpConn
		log.Print("客户端地址：", tcpConn.RemoteAddr().String())

		go handleConn(tcpConn)
	}
}

func handleConn(conn *net.TCPConn) {
	defer conn.Close()

	//接收消息
	var recvData = make([]byte, 1024)
	len, err := conn.Read(recvData)
	if err != nil {
		log.Print(err)
		return
	}
	log.Printf("接收数据，长度:%d, 数据:%s", len, string(recvData))

	//给客户端发送消息
	var data = conn.RemoteAddr().String()
	len, err = conn.Write([]byte(data))
	if err != nil {
		log.Print(err)
		return
	}
	log.Printf("发送数据，长度:%d, 数据:%s", len, data)

}
