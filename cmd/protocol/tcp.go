package protocol

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

type TCP struct {
	Port int `yaml:"port"`
}

func (t *TCP) InitFrom(channel chan string) {
	log.Printf("TCP: InitFrom: port=%d", t.Port)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", t.Port))
	if err != nil {
		log.Fatalf("Error creating listener: %v\n", err)
	}
	defer listener.Close()

	conn, err := listener.Accept()
	if err != nil {
		log.Fatalf("Error accepting connection: %v\n", err)
	}
	defer conn.Close()

	log.Printf("TCP: connection accepted: %s", conn.RemoteAddr().String())

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		message := scanner.Text()
		log.Printf("TCP: from: \"%s\"", message)
		channel <- message
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error reading from connection: %v\n", err)
	}
}

func (t *TCP) InitTo(channel chan string) {
	log.Printf("TCP: InitTo: port=%d", t.Port)

	for {
		t.initConnection(channel)
		time.Sleep(1 * time.Second)
	}
}

func (t *TCP) initConnection(channel chan string) {
	var conn net.Conn
	var err error

	conn, err = net.Dial("tcp", fmt.Sprintf(":%d", t.Port))
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		return
	}

	defer conn.Close()

	log.Printf("TCP: connection accepted: %s", conn.RemoteAddr().String())

	for {
		message := <-channel
		_, err := conn.Write([]byte(message + "\n"))
		if err != nil {
			fmt.Println("Error sending message:", err.Error())
			return
		}

		log.Printf("TCP: to: \"%s\"", message)
	}
}
