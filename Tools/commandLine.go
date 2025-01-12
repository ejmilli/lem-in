package Tools

import (
	"fmt"
	"os"
	"strings"
)

var lines []string

func CheckCommandLineArg() []string {
	if len(os.Args) != 2 {
		fmt.Println("Usage: rooms <roomfile>")
		os.Exit(1)
	}
	arg := os.Args[1]

	if arg[len(arg)-4:] != ".txt" {
		fmt.Println("Room file must end with .txt")
		os.Exit(1)
	}

	file, err := os.ReadFile(arg)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	lines = strings.Split(string(file), "\n")
	if len(lines) < 6 {
		fmt.Println("Error: Room file must contain at least 6 lines")
		os.Exit(1)
	}

	return lines
}
