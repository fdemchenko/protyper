package main

import (
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"
)

const (
	EndOfTransmissionCode = 4
)

func main() {
	err := CBreakMode()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT)

	go func() {
		<-signals

		fmt.Println("Signal recieved SIGINT")
		CannonicalMode()
		os.Exit(0)
	}()

	buffer := make([]byte, 1)
	for {
		_, err := os.Stdin.Read(buffer)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err.Error())
			}
			break
		}
		if buffer[0] == EndOfTransmissionCode {
			break
		}
		fmt.Printf("%d\n", buffer[0])
	}
	CannonicalMode()
}
