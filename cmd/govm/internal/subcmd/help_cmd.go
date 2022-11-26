package subcmd

import (
	"fmt"
	"strings"

	"barglvojtech.net/x/pkg/flagutil"
)

var helpUsage = strings.TrimSpace(`
Govm is tool to quickly download and switch between version.
Version can be specified with 'go', 'v' or without prefix.

Versions are downloaded into ${libDir}/versions/{version}.
Quick switching is done via symlink of go executable into ${binDir}.

You can configure these path in ${configFile}.

--- 
There are 3 subcomands:
	install
	switch
	help
	list
`)

func Help(args *flagutil.Args) {
	fmt.Println(helpUsage)
}
