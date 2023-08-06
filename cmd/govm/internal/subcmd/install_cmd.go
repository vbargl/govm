package subcmd

import (
	"fmt"
	"strings"

	"barglvojtech.net/govm/pkg/embi/ops"
	"barglvojtech.net/govm/pkg/x/opsutil"
	"barglvojtech.net/govm/pkg/x/versionutil"

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

	var err error

	rawVersion := args.Take()

	version, err := versionutil.Normalize(rawVersion)
	errutil.LogFatalIfErr(err, nil)

	filename, err := versionutil.Format(versionutil.FilenameFormat, versionutil.Qualify(version))
	errutil.LogFatalIfErr(err, nil)

	err = opsutil.Process(
		ops.NewDownloadOp(
			ops.SetFilenameToDownloadOp(filename),
			ops.SetVersionToDownloadOp(rawVersion),
		),
		ops.NewExtractOp(
			ops.SetFilenameToExtractOp(filename),
			ops.SetVersionToExtractOp(rawVersion),
		),
	)

	errutil.LogPrintIfErr(err, nil)
}
