package main

import (
	"fmt"
	"os"

	"barglvojtech.net/govm/cmd/govm/internal/subcmd"
	"barglvojtech.net/x/pkg/flagutil"
)

func main() {
	var (
		args = flagutil.Args(os.Args[1:])
		cmd  subcommand
	)

	switch subCmd := args.Take(); subCmd {
	case "install":
		cmd = subcmd.Install
	case "switch":
		cmd = subcmd.Switch
	case "list":
		cmd = subcmd.List
	default:
		cmd = subcmd.Help
	}

	switch {
	case cmd != nil:
		cmd(&args)
	default:
		fmt.Println("missing subcommand, please contact developers")
	}
}

type subcommand func(*flagutil.Args)
