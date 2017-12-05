package main

import "fmt"

func main() {
	// 建立新的channel.
	messages := make(chan string)

	// 發送一個值到 'channel <-'
	go func() { messages <- "ping" }()

	// '<-channel' 從channel接收值
	msg := <-messages
	fmt.Println(msg)
}
