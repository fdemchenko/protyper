package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fatih/color"
)

const (
	EndOfTransmissionCode = 4
)

type Config struct {
	outputSpeed int
	outputColor ColorArg
	autoTyping  bool
	interval    int
}

func main() {
	var cfg Config
	flag.IntVar(&cfg.outputSpeed, "speed", 1, "characters amount to output in one signals")
	flag.Var(&cfg.outputColor, "color", "output ANSI color (default white)")
	flag.BoolVar(&cfg.autoTyping, "auto", false, "auto typing, not using keyboard")
	flag.IntVar(&cfg.interval, "interval", 50, "auto typing interval in milliseconds, has no effect without auto mode")
	flag.Parse()

	if flag.Arg(0) == "" {
		log.Fatalln("input file not provided")
	}

	file, err := os.Open(flag.Arg(0))
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer file.Close()
	fileReader := bufio.NewReader(file)

	err = CBreakMode()
	if err != nil {
		log.Fatalln(err.Error())
	}

	var typer Typer
	if cfg.autoTyping {
		typer = NewAutoTyper(time.Millisecond * time.Duration(cfg.interval))
	} else {
		typer = NewKeyboardTyper(os.Stdin)
	}

	SetUpSignalsHandling()

	buffer := make([]rune, cfg.outputSpeed)
	colorizer := color.New(color.Attribute(cfg.outputColor))

	for range typer.Start() {
		n, err := ReadRunes(fileReader, buffer)
		if err != nil {
			typer.Stop()
		} else {
			colorizer.Print(string(buffer[:n]))
		}
	}
	CannonicalMode()
}

func SetUpSignalsHandling() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT)

	go func() {
		<-signals
		CannonicalMode()
		os.Exit(0)
	}()
}
