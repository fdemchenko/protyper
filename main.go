package main

import (
	"bufio"
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
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "%s\n", "filename is not provided")
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
	defer file.Close()
	fileReader := bufio.NewReader(file)

	err = CBreakMode()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
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

	reader := bufio.NewReader(os.Stdin)
	for {
		code, err := reader.ReadByte()
		if err != nil {
			if err != io.EOF {
				fmt.Fprintf(os.Stderr, "%s\n", err.Error())
			}
			break
		}
		if code == EndOfTransmissionCode {
			break
		}

		char, _, err := fileReader.ReadRune()
		if err != nil {
			if err != io.EOF {
				fmt.Fprintf(os.Stderr, "%s\n", err.Error())
			}
			break
		}
		fmt.Printf("%c", char)
	}
	CannonicalMode()
}
