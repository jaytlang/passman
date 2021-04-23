package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/jaytlang/passman/pscons"
)

func parseArgs(args []string) (string, error) {
	if len(args) < 1 {
		return "", errors.New("too few arguments")
	} else if len(args) > 1 {
		return "", errors.New("too any arguments")
	}

	s, err := os.Stat(args[0])
	if err != nil {
		return "", err
	} else if !s.IsDir() {
		return "", errors.New("provided path is not a directory")
	}
	return args[0], nil
}

func main() {
	d, err := parseArgs(os.Args[1:])
	if err != nil {
		fmt.Println("Usage: passman <path/to/sd>")
		fmt.Println("Error: ", err)
		return
	}

	if err := os.Chdir(d); err != nil {
		fmt.Println("Error:", err)
	}

	pscons.RunPsCons()
}
