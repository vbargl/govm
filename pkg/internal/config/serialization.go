package config

import (
	"github.com/BurntSushi/toml"

	"barglvojtech.net/x/pkg/errutil"
)

// Parse takes absolute location of file and parse it into Config.
// If file does not exists or its content is malformed and
// cannot be be parse into config, err is returned.
func Parse(absoluteFilepath string) (cfg Config, err error) {
	_, err = toml.DecodeFile(absoluteFilepath, &cfg)
	errutil.AssignIfErr(&err, err, errutil.PrefixWithFormatted("could not parse config %s", absoluteFilepath))
	return
}
