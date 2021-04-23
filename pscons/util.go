package pscons

import "fmt"

func printError(e error) {
	fmt.Println("Error:", e)
	fmt.Println("Type 'help' for more help")
}

func printUsageMsg() {
	fmt.Println("Usage: ")
	fmt.Println("\tget <domain>")
	fmt.Println("\tgen <domain>")
	fmt.Println("\tfind <search string>")
	fmt.Println("\tquit")
}
