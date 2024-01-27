package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/fatih/color"
)

const (
	EndOfTransmissionCode = 4
)

type Config struct {
	outputSpeed int
	outputColor ColorArg
}

func main() {
	var cfg Config
	flag.IntVar(&cfg.outputSpeed, "speed", 1, "characters amount to output by button press")
	flag.Var(&cfg.outputColor, "color", "output ANSI color (default white)")
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

	SetUpSignalsHandling()

	reader := bufio.NewReader(os.Stdin)
	buffer := make([]rune, cfg.outputSpeed)
	colorizer := color.New(color.Attribute(cfg.outputColor))
	for {
		code, err := reader.ReadByte()
		if err != nil {
			if err != io.EOF {
				log.Println(err.Error())
			}
			break
		}
		if code == EndOfTransmissionCode {
			break
		}

		n, err := ReadRunes(fileReader, buffer)
		if err != nil {
			if err != io.EOF {
				log.Println(err.Error())
			}
			break
		}
		colorizer.Print(string(buffer[:n]))
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
