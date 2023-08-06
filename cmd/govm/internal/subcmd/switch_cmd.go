package subcmd

import (
	"fmt"
	"strings"

	"barglvojtech.net/govm/pkg/embi/ops"

	"barglvojtech.net/x/pkg/errutil"
	"barglvojtech.net/x/pkg/flagutil"
)

var switchUsage = strings.TrimSpace(`
Usage:
	govm switch {version}

Parameters:
	version		version of go to download

Examples:
	# switch to version go.1.20
	govm switch go1.20

	# switch to version go1.20rc2
	govm switch 1.20rc2
`)

func Switch(args *flagutil.Args) {
	if args.Len() != 1 {
		fmt.Println(switchUsage)
		return
	}

	version := args.Take()

	op := ops.NewSwitchOp(
		ops.SetVersionToSwitchOp(version),
	)

	err := op.Process()
	errutil.LogPrintIfErr(err, nil)
}
