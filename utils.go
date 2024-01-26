package main

import (
	"bufio"
	"io"
)

func ReadRunes(reader *bufio.Reader, chars []rune) (int, error) {
	for i := 0; i < len(chars); i++ {
		char, _, err := reader.ReadRune()
		if err != nil {
			if err == io.EOF && i > 0 {
				return i, nil
			}
			return i, err
		}
		chars[i] = char
	}
	return len(chars), nil
}
