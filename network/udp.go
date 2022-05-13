package network

import (
	"log"
	"net"
)

func startUdp() error {
	var addr, _ = net.ResolveUDPAddr("udp", "127.0.0.1:9999")
	var listen, err = net.ListenUDP("udp", addr)
	if err != nil {
		return err
	}
	defer listen.Close()

	log.Println("udp listen:", addr.String())

	for {
		// 读消息
		var data = make([]byte, 1024)
		len, tcpConn, err := listen.ReadFromUDP(data)
		if err != nil {
			log.Println(err)
			continue
		}
		log.Printf("接收消息,长度:%d, 数据:%s", len, string(data))

		//发送消息
		var writeData = tcpConn.AddrPort().Addr().String()
		len, err = listen.WriteToUDP([]byte(writeData), tcpConn)
		if err != nil {
			log.Println(err)
			continue
		}
		log.Printf("发送消息,长度:%d, 数据:%s", len, writeData)
	}
}
