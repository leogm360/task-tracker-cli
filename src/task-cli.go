package main

import (
	"os"
	"task-tracker-cli/src/cli"
)

func main() {
	cli.CLI(os.Args[1:])
}
