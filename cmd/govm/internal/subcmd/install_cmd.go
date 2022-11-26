package subcmd

import (
	"fmt"
	"strings"

	"barglvojtech.net/govm/pkg/embi/ops"
	"barglvojtech.net/x/pkg/errutil"
	"barglvojtech.net/x/pkg/flagutil"
)

var installUsage = strings.TrimSpace(`
Usage:
	govm install {version}

Parameters:
	version		version of go to install

Examples:
	# install version go.1.20
	govm install go1.20

	# install version go1.20rc2
	govm install 1.20rc2
`)

func Install(args *flagutil.Args) {
	if args.Len() != 1 {
		fmt.Println(installUsage)
		return
	}

	version := args.Take()

	op := ops.NewDownloadOp(
		ops.SetVersionToDownloadOp(version),
	)

	err := op.Process()
	errutil.LogPrintIfErr(err, nil)
}
