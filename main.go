package main

import (
	"fmt"
	"math"
	"os"
	"os/signal"
	"syscall"
	"unsafe"
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
			fmt.Println(err.Error())
			break
		}
		if buffer[0] == 4 {
			break
		}
		fmt.Printf("%d\n", buffer[0])
	}
	CannonicalMode()
}

func CBreakMode() error {
	var ter syscall.Termios
	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, uintptr(syscall.Stdin), syscall.TCGETS, uintptr(unsafe.Pointer(&ter)))
	if errno != 0 {
		return errno
	}

	ter.Lflag &= math.MaxUint32 ^ (syscall.ICANON | syscall.ECHO)
	_, _, errno = syscall.Syscall(syscall.SYS_IOCTL, uintptr(syscall.Stdin), syscall.TCSETS, uintptr(unsafe.Pointer(&ter)))
	if errno != 0 {
		return errno
	}
	return nil
}

func CannonicalMode() error {
	var ter syscall.Termios
	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, uintptr(syscall.Stdin), syscall.TCGETS, uintptr(unsafe.Pointer(&ter)))
	if errno != 0 {
		return errno
	}

	ter.Lflag |= syscall.ICANON | syscall.ECHO
	_, _, errno = syscall.Syscall(syscall.SYS_IOCTL, uintptr(syscall.Stdin), syscall.TCSETS, uintptr(unsafe.Pointer(&ter)))
	if errno != 0 {
		return errno
	}
	return nil
}
