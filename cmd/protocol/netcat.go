package protocol

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type Netcat struct {
	Port int `yaml:"port"`
}

func (t *Netcat) InitFrom(channel chan string) {
	log.Printf("Netcat: InitFrom: port=%d", t.Port)

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

	log.Printf("Netcat: connection accepted: %s", conn.RemoteAddr().String())

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		message := scanner.Text()
		log.Printf("Netcat: from: \"%s\"", message)
		channel <- message
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error reading from connection: %v\n", err)
	}
}

func (t *Netcat) InitTo(channel chan string) {
	log.Printf("Netcat: InitTo: port=%d", t.Port)

	conn, err := net.Dial("tcp", fmt.Sprintf(":%d", t.Port))
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		return
	}
	defer conn.Close()

	log.Printf("Netcat: connection accepted: %s", conn.RemoteAddr().String())

	for {
		message := <-channel
		_, err = conn.Write([]byte(message + "\n"))
		if err != nil {
			fmt.Println("Error sending message:", err.Error())
			return
		}

		log.Printf("Netcat: to: \"%s\"", message)
	}
}
