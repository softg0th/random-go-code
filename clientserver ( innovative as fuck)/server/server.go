package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Server struct {
	host string
	port string
}

type Clients struct {
	allClients map[string]net.Conn
	usernames  map[string]string
}

func NewClients() *Clients {
	return &Clients{
		allClients: make(map[string]net.Conn),
		usernames:  make(map[string]string),
	}
}

func createServer(host, port string) *Server {
	return &Server{
		host: host,
		port: port,
	}
}

func broadcastMessageSend(users *Clients, message []byte) {
	for _, userConn := range users.allClients {
		data, err := userConn.Write(message)
		if err != nil {
			fmt.Printf("Error sending message: %s\n", err)
			return
		}
		if data != len(message) {
			fmt.Printf("Failed to send entire message\n")
			return
		}
	}
}

func disconnectClient(users *Clients, connAddr string) {
	delete(users.allClients, connAddr)
	delete(users.usernames, connAddr)
}

func handleRequest(conn net.Conn, users *Clients) {
	time.Sleep(3)
	connectionUsername := users.usernames[conn.RemoteAddr().String()]
	fmt.Printf("User connected: %s", connectionUsername)
	greetingMessage := []byte("New user connected:" + connectionUsername)
	go broadcastMessageSend(users, greetingMessage)
	for {
		input := make([]byte, (1024 * 4))
		bytes, err := conn.Read(input)
		if err != nil {
			conn.Close()
			return
		}
		message := input[0:bytes]
		incommingMessage := string(input[:bytes])
		if incommingMessage != "exit" {
			go broadcastMessageSend(users, message)
		} else {
			toRemove := ""
			for addr, con := range users.allClients {
				if con == conn {
					toRemove = addr
					break
				}
			}
			disconnectClient(users, toRemove)
		}
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Got exception for: %s", err)
	}
	clients := NewClients()
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	server := createServer(host, port)
	server.port = ":" + server.port
	connector, errorr := net.Listen("tcp", server.port)
	if errorr != nil {
		fmt.Printf("Got exception for: %s", err)
		defer connector.Close()
		return
	}

	for {
		conn, err := connector.Accept()
		if err != nil {
			fmt.Printf("Got exception for: %s", err)
		}
		clients.allClients[conn.RemoteAddr().String()] = conn
		currentTime := time.Now()
		timeBytes := []byte(currentTime.String())
		hash := md5.Sum(timeBytes)
		hashString := hex.EncodeToString(hash[:])
		clients.usernames[conn.RemoteAddr().String()] = hashString
		go handleRequest(conn, clients)
	}
}
