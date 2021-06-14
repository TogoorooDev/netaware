package main

import (
	//	"encoding/hex"
	"bufio"
	"log"
	"net"
	"os"
)

const (
	connHost = "localhost"
	connPort = "3295"
	connType = "tcp"
)

/*func buildPeerlist() []string {
	var peerlist []string
	for {
		potpeer := <-peerlistc
		if potpeer == "0" {
			break
		}
		peerlist = append(peerlist, <-peerlistc)
	}
	return peerlist
}*/

func main() {

	//start listening on port 3295
	listener, err := net.Listen(connType, ":"+connPort)
	if err != nil {
		log.Fatal("Could not listen to port ", connPort, err.Error())
		os.Exit(1)
	}

	//Connect to rest of network
	go networkInit()

	defer listener.Close()

	for {
		c, err := listener.Accept()
		if err != nil {
			log.Println("Connection Error ", err.Error())
			return
		}
		//log.Println("Peer Connected")

		//log.Println("Peer " + c.RemoteAddr().String() + " connected")
		go serverMain(c)
	}
}

func serverMain(conn net.Conn) {
	for {
		buffer, err := bufio.NewReader(conn).ReadBytes('\n')

		if err != nil {
			log.Println("Peer", conn.RemoteAddr().String(), "Disconnected")
			conn.Close()
			return
		}
		log.Println("Message from", conn.RemoteAddr().String(), ":", string(buffer[:len(buffer)-1]))
		//[]byte{'0', '0'}

		conn.Write([]byte{'\x00'})
	}
}
