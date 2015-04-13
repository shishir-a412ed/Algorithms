//Build a simple ECHO client server application using UNIX datagram sockets.
//Server side implementation.

package main

import (
	"fmt"
	"log"
	"net"
)

func main() {

	conn, err := net.ListenUnixgram("unixgram", &net.UnixAddr{"/tmp/unixdomain", "unixgram"})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Server listening for incoming messages")
	var buf [1024]byte
	for {
		n, err := conn.Read(buf[:])
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("%s\n", string(buf[:n]))
	}

}
