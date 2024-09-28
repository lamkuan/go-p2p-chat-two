package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	listener, err := net.ListenUDP("udp", &net.UDPAddr{Port: 9090})
	errorHanlder(err)
	defer listener.Close()
	fmt.Printf("Server Address:  %s]\n", listener.LocalAddr().String())
	peers := make([]*net.UDPAddr, 2)
	buf := make([]byte, 256)
	n, addr, _ := listener.ReadFromUDP(buf)
	peers[0] = addr
	fmt.Printf("Read from <%s>:%s\n", addr.String(), buf[:n])
	n, addr, _ = listener.ReadFromUDP(buf)
	fmt.Printf("Read from <%s>:%s\n", addr.String(), buf[:n])
	peers[1] = addr
	fmt.Println("Begin Nat")
	listener.WriteToUDP([]byte(peers[0].String()), peers[1])
	listener.WriteToUDP([]byte(peers[1].String()), peers[0])

	time.Sleep(10 * time.Second)
}

func errorHanlder(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
