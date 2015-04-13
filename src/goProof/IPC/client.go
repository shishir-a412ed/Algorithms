//Build a simple ECHO client server application using UNIX datagram sockets.
//Client side code.

package main

import (
	"log"
	"net"
)

func main() {

	conn, err := net.DialUnix("unixgram", nil, &net.UnixAddr{"/tmp/unixdomain", "unixgram"})
	if err != nil {
		log.Fatalln(err)
	}

	defer conn.Close()

	conn.Write([]byte("Hello World"))
	conn.Write([]byte("Unix socket connected successfully"))
	conn.Write([]byte("Bye Bye"))

}
