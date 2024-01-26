package main

import (
	"math"
	"syscall"
	"unsafe"
)

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
