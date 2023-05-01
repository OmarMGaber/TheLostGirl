package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Where did you find the girl?")

	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	for { // if I want the user to send multiple times
		var userInput string
		fmt.Print("Enter the location name or ('exit') to log out from the server: ")
		fmt.Scanln(&userInput)
		_, er := conn.Write([]byte(userInput))

		if er != nil {
			fmt.Println(er)
			return
		}
		if userInput == "exit" || userInput == "Exit" {
			conn.Close()
			return
		}
	}
}
