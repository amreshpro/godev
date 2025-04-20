package tcpserver

import (
	"bufio"
	"log"
	"net"
)

func Main() {
	listen, err := net.Listen("tcp", ":5678")
	if err != nil {
		log.Fatal("failed to start server",err)
	}
	defer listen.Close()
	// infinte loop
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal("error: ", err)
			return
		}
		go handleConnection(conn)
		defer conn.Close()
	}

}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("error: ", err)
		return
	}
	println(line)

}
