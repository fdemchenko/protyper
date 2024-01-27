package main

import (
	"bufio"
	"fmt"
	"io"
	"time"
)

type Typer interface {
	Start() (chan struct{}, error)
	Stop()
}

type KeyboardTyper struct {
	keyboardReader *bufio.Reader
	stopChan       chan struct{}
}

type AutoTyper struct {
	interval time.Duration
	stopChan chan struct{}
}

func NewAutoTyper(interval time.Duration) AutoTyper {
	return AutoTyper{interval: interval, stopChan: make(chan struct{}, 1)}
}

func NewKeyboardTyper(keyboard io.Reader) KeyboardTyper {
	return KeyboardTyper{keyboardReader: bufio.NewReader(keyboard), stopChan: make(chan struct{}, 1)}
}

func (at AutoTyper) Start() (chan struct{}, error) {
	outgoingSignals := make(chan struct{})

	go func() {
		for {
			select {
			case <-at.stopChan:
				close(outgoingSignals)
				return
			default:
				time.Sleep(at.interval)
				outgoingSignals <- struct{}{}
			}
		}
	}()

	return outgoingSignals, nil
}

func (kt KeyboardTyper) Start() (chan struct{}, error) {
	outgoingSignals := make(chan struct{})
	err := CBreakMode()
	if err != nil {
		return nil, fmt.Errorf("cbreak mode: %w", err)
	}

	go func() {
		for {
			select {
			case <-kt.stopChan:
				goto CLEAN_UP
			default:
				code, err := kt.keyboardReader.ReadByte()
				if err != nil || code == EndOfTransmissionCode {
					goto CLEAN_UP
				}
				outgoingSignals <- struct{}{}
			}
		}
	CLEAN_UP:
		close(outgoingSignals)
		CannonicalMode()
	}()

	return outgoingSignals, nil
}

func (kt KeyboardTyper) Stop() {
	kt.stopChan <- struct{}{}
}

func (at AutoTyper) Stop() {
	at.stopChan <- struct{}{}
}
