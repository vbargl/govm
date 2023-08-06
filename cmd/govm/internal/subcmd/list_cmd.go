package subcmd

import (
	"fmt"
	"strings"

	"golang.org/x/exp/slices"

	"barglvojtech.net/govm/pkg/embi/ops"

	"barglvojtech.net/x/pkg/errutil"
	"barglvojtech.net/x/pkg/flagutil"
)

var listUsage = strings.TrimSpace(`
Usage:
	govm list
`)

func List(args *flagutil.Args) {
	if args.Len() != 0 {
		fmt.Println(listUsage)
	}

	op := ops.NewListOp()

	err := op.Process()
	if errutil.LogPrintIfErr(err, nil) {
		return
	}

	ver := op.Versions()
	slices.Reverse(ver)

	fmt.Println("Available versions:")
	for _, v := range ver {
		fmt.Printf("\t%s\n", v.Version.String())
	}
}
