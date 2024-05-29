package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
)

type Server struct {
	url string
}

func sendMessage(conn net.Conn, exitMessage string) {
	fmt.Println(exitMessage)
	if exitMessage == "exit" {
		fmt.Println("exitstart")
		if _, err := conn.Write([]byte(exitMessage)); err != nil {
			fmt.Println("Error while sending!")
			return
		}
	}
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("Error while sending!")
		return
	}
	if data, err := conn.Write([]byte(input)); data == 0 || err != nil {
		fmt.Println("Error while sending!")
		return
	}
}

func receiveMessage(conn net.Conn) {
	data := make([]byte, 1024)
	messageLen, err := conn.Read(data)
	if err != nil {
		fmt.Println("Error while reading!")
		return
	}
	incomingMessage := string(data[:messageLen])
	fmt.Println(incomingMessage)
}

func main() {
	var serverUrl string

	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Got exception for: %s", err)
	}

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	serverUrl = host + ":" + port
	server := &Server{
		url: serverUrl,
	}
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	conn, err := net.Dial("tcp", server.url)
	if err != nil {
		fmt.Printf("Got exception for: %s", err)
		defer conn.Close()
	}
	for {
		go sendMessage(conn, "")
		select {
		case <-quit:
			fmt.Println("Exiting...")
			sendMessage(conn, "exit")
			return
		default:
			receiveMessage(conn)
		}
	}
}
