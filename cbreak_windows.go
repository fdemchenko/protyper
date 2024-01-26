package main

import (
	"golang.org/x/sys/windows"
)

func getConsoleParams() (windows.Handle, uint32, error) {
	handle, err := windows.GetStdHandle(windows.STD_INPUT_HANDLE)
	if err != nil {
		return 0, 0, err
	}
	var consoleMode uint32
	err = windows.GetConsoleMode(handle, &consoleMode)
	if err != nil {
		return 0, 0, nil
	}
	return handle, consoleMode, nil
}

func CBreakMode() error {
	handle, consoleMode, err := getConsoleParams()
	if err != nil {
		return err
	}
	consoleMode &^= windows.ENABLE_LINE_INPUT
	consoleMode &^= windows.ENABLE_ECHO_INPUT
	return windows.SetConsoleMode(handle, consoleMode)
}

func CannonicalMode() error {
	handle, consoleMode, err := getConsoleParams()
	if err != nil {
		return err
	}
	consoleMode |= windows.ENABLE_LINE_INPUT
	consoleMode |= windows.ENABLE_ECHO_INPUT
	return windows.SetConsoleMode(handle, consoleMode)
}
