package cli

import "fmt"

type PositionalCmdArgs string

const (
	Create PositionalCmdArgs = "create"
)

func CLI(args []string) {
	cmd := PositionalCmdArgs(args[0])
	arg := args[1]

	switch cmd {
	case Create:
		CreateTask(arg)
	default:
		panic(fmt.Sprintf("unknown command argument '%s'", cmd))
	}
}
