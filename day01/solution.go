package day01

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func part01(filename string) (int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	floor := 0

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanRunes)
	for s.Scan() {
		switch s.Text() {
		case "(":
			floor++
		case ")":
			floor--
		default:
			return 0, errors.New(fmt.Sprintf("invalid character found in input: %q", s.Text()))
		}
	}

	if err = s.Err(); err != nil {
		return 0, err
	}

	return floor, nil
}
