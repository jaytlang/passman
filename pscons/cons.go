package pscons

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func prompt() {
	fmt.Print("passman > ")
}

func RunPsCons() {
	sc := bufio.NewScanner(os.Stdin)
	prompt()
	for sc.Scan() {
		s := sc.Text()
		if s == "" {
			printUsageMsg()
			prompt()
			continue
		}

		err := runCmd(strings.Split(s, " "))
		if err != nil {
			printError(err)
		}
		prompt()
	}
}
