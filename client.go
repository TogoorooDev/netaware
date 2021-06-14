package main

import (
	"bufio"
	"log"
	"net"
	"time"
)

func networkInit() {

	/*hosts, err := Hosts("10.0.0.0/24")

	if err != nil {
		log.Fatal("Peer search error", err.Error())
	}*/

	//Build Peerlist

	//peerlistc := make(chan string)
	//peerlistc_end := make(chan bool)

	time.Sleep(8 * time.Second)

	log.Println("Starting peer discovery")

	/*for _, host := range hosts {
		//log.Println("Starting peer discovery")
		go peerConn(host)
	}*/

	go peerConn("10.0.0.179")

}

func clientMain(conn net.Conn) {

	//Send a LANSync Ping to the supposed peer
	conn.Write([]byte{'\x00', '\n'})
	ping, err := bufio.NewReader(conn).ReadBytes('\n')

	if err != nil {
		log.Println("Read error reading. Assuming device is not a peer", err.Error)
	}

	if ping[0] != '\x00' {
		log.Println("Ping not returned. Not a peer")
	}

	log.Println("Ping returned. Peer assumed")

	for {
		buffer, err := bufio.NewReader(conn).ReadBytes('\n')

		if err != nil {
			log.Println("Peer", conn.RemoteAddr().String(), "Disconnected")
			return
		}

		log.Println(buffer)

	}
}

func peerConn(ip string) {
	log.Println("Testing ", ip)
	conn, err := net.Dial(connType, ip+":"+connPort)
	if err != nil {
		log.Println(ip, err.Error())
		return
	} else {
		//log.Println(ip, " is a peer")

		clientMain(conn)
	}
}
