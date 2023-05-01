package main

import (
	"fmt"
	"net"
)

const maxNumberOfRepeatedLocation = 3

type location struct {
	name  string
	count int
}

func main() {
	locations := make(map[string]*location)

	found := make(chan bool)

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()

	fmt.Println("Server started")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleConnection(conn, locations, found)
		select {
		case <-found:
			listener.Close()
			fmt.Println("Server closed")
			return
		default:
		}
	}
}

func handleConnection(conn net.Conn, locations map[string]*location, found chan bool) {
	defer conn.Close()
	fmt.Printf("User: %s has entered the server \n", conn.RemoteAddr().String())

	for {
		buffer := make([]byte, 1024)
		_, err := conn.Read(buffer)
		if err != nil {
			fmt.Println(err)
			return
		}

		userInput := string(buffer)

		if (userInput == "exit") || (userInput == "Exit") {
			fmt.Printf("User: %s has exit the server \n", conn.RemoteAddr().String())
		} else {
			fmt.Printf("User: %s entered: %s\n", conn.RemoteAddr().String(), userInput)
			if loc, ok := locations[userInput]; ok {
				loc.count++
				if loc.count >= maxNumberOfRepeatedLocation {
					fmt.Printf("Location \"%s\" has been repeated %d times, closing server\n", userInput, maxNumberOfRepeatedLocation)
					found <- true
					conn.Close()
					return
				}
			} else {
				locations[userInput] = &location{name: userInput, count: 1}
			}
		}
	}
}
