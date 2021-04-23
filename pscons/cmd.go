package pscons

import (
	"errors"
	"fmt"
	"os"
)

func checkLen(args []string, l int) error {
	if len(args) == l {
		return nil
	}
	return errors.New("incorrect arguments")
}

func runCmd(args []string) error {
	switch args[0] {
	case "get":
		if err := checkLen(args, 2); err != nil {
			return err
		}
		pass, err := get(args[1])
		if err != nil {
			return err
		}
		fmt.Printf("Your password: %s\n", pass)
		return nil
	case "gen":
		if err := checkLen(args, 2); err != nil {
			return err
		}
		pass, err := gen(args[1], false)
		if err != nil {
			return err
		}
		fmt.Printf("New password: %s\n", pass)
		return nil
	case "find":
		if err := checkLen(args, 2); err != nil {
			return err
		}
		opt, err := find(args[1])
		if err != nil {
			return err
		}
		fmt.Printf("Passwords matching: %v\n", opt)
		return nil
	case "del":
		if err := checkLen(args, 2); err != nil {
			return err
		}
		if err := del(args[1]); err != nil {
			return err
		}
		fmt.Printf("Deleted password for %s\n", args[1])
		return nil
	case "help":
		printUsageMsg()
		return nil
	case "quit":
		fmt.Println("Bye!")
		os.Exit(0)
	}
	return errors.New("unknown command")
}
