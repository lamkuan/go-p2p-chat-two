package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

func main() {
	port := flag.Int("p", 0, "port")
	username := flag.String("u", "", "username")
	remoteIP := flag.String("i", "", "remote ip")
	remotePort := flag.Int("r", 0, "remote port")

	flag.Parse()

	localAddr := net.UDPAddr{Port: *port}
	remoteAddr := net.UDPAddr{IP: net.ParseIP(*remoteIP), Port: *remotePort}

	conn, err := net.DialUDP("udp", &localAddr, &remoteAddr)
	errorHandler(err)
	conn.Write([]byte("I am" + *username))
	buf := make([]byte, 256)
	n, _, err := conn.ReadFromUDP(buf)
	errorHandler(err)

	toAddr := parseAddr(string(buf[:n]))
	fmt.Println("User Address is: ", toAddr)

	conn.Close()

	p2p(&localAddr, &toAddr)

}

func parseAddr(sourceAddr string) net.UDPAddr {
	addr := strings.Split(sourceAddr, ":")
	ip := net.ParseIP(addr[0])
	port, err := strconv.Atoi(addr[1])
	errorHandler(err)
	return net.UDPAddr{IP: ip, Port: port}
}

func p2p(srcAddr *net.UDPAddr, dstAddr *net.UDPAddr) {
	conns2d, err := net.DialUDP("udp", srcAddr, dstAddr)
	errorHandler(err)
	// defer conns2d.Close()
	conns2d.Write([]byte("Hello\n"))

	go func() {
		buf := make([]byte, 256)
		for {
			n, _, _ := conns2d.ReadFromUDP(buf)
			if n > 0 {
				fmt.Printf("%s: %s", dstAddr, string(buf))
			}
		}
	}()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("p2p>")
		data, err := reader.ReadString('\n')
		errorHandler(err)
		conns2d.Write([]byte(data))
	}
}

func errorHandler(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
