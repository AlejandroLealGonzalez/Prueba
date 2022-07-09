package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	go servidor()
	var input string
	fmt.Scanln(&input)
}

func servidor() {
	s, err := net.Listen("tcp", ":3000")
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Println("Escuchando por el puerto localhost:3000")
	for {
		c, err := s.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleClient(c)
	}
}

func handleClient(c net.Conn) {
	b := make([]byte, 100)
	bs, err := c.Read(b)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Mensaje: ", b[:bs])
		fmt.Println("Bytes: ", bs)
	}

}
