package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	// "os"
)

func Main() {
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		fmt.Println("Error dialing", err.Error())
		return
	}
	defer conn.Close()

	// create a bufio writer to send messages to the server
	writer := bufio.NewWriter(conn)

	// read input from the user
	// reader := bufio.NewReader(os.Stdin)

	for i := 0; i < 100; i++ {
		fmt.Print("Enter text:", i)
		// text, _ := reader.ReadString('\n')

		// send the message to the server
		_, err := writer.WriteString(strconv.Itoa(i))
		if err != nil {
			fmt.Println("Error sending message", err.Error())
			return
		}

		// flush the writer to ensure the message is sent
		err = writer.Flush()
		if err != nil {
			fmt.Println("Error flushing writer", err.Error())
			return
		}
	}
	func() {
		int i = 0
		for {
			i++
			fmt.Print("Enter text:", i)
			// text, _ := reader.ReadString('\n')
		}
	}
}
