package main

import (
	"fmt"
	"lemin/Tools"
)

func main() {

	lines := Tools.CheckCommandLineArg()
	_, _, err := Tools.ReadInput(lines)
	if err != nil {
		fmt.Println(err)
		return
	}
}
