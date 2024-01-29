# Protyper

## Description

This program allows you to press some random keys to produce predefined output in terminal. Protyper works on both Linux and Windows systems by switching terminal to so called **raw** input mode.
Auto mode just creates output without typing, good for background

## Installation
```go install github.com/fdemchenko/protyper``` 


## Usage

```
Usage: protyper [-auto] [-color value] [-interval int] [-speed int] [FILE]
  -auto
        auto typing, not using keyboard
  -color value
        output ANSI color (default white)
  -interval int
        auto typing interval in milliseconds, has no effect without auto mode (default 50)
  -speed int
        characters amount to output in one signal (default 1)
```
### Available colors
- white
- green
- cyan
- yellow
- blue
- red
- magenta

## Examples
