package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

type Job struct {
	workerID int
	message  string
}

func worker(workerID int, jobChan <-chan Job) {
	for job := range jobChan {
		fmt.Printf("Worker %d processing job: %s\n", workerID, job.message)
		if strings.Contains(job.message, "PRIORIDADE") {
			time.Sleep(time.Second * 1) // Processa mais rápido
		} else {
			time.Sleep(time.Second * 2) // Processa mais lentamente
		}
	}
}

func handleConnection(conn net.Conn, jobChan chan<- Job) {
	defer conn.Close()
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			break
		}
		message := string(buffer[:n])
		jobChan <- Job{message: message}
	}
}

func main() {
	jobChan := make(chan Job)
	for i := 0; i < 3; i++ {
		go worker(i, jobChan)
	}

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		go handleConnection(conn, jobChan)
	}
}
