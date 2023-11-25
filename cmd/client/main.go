package main

import (
	"fmt"
	"math/rand"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}
	defer conn.Close()

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 1000; i++ {
		var msg string
		if rand.Intn(2) == 0 {
			msg = "mensagem normal\n"
		} else {
			msg = "PRIORIDADE\n"
		}
		conn.Write([]byte(msg))
		time.Sleep(time.Millisecond * 10)
	}
}
