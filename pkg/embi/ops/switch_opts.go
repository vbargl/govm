package ops

import (
	"barglvojtech.net/govm/pkg/embi/env"
	"barglvojtech.net/govm/pkg/internal/versionutil"
	"barglvojtech.net/x/pkg/errutil"
)

func SetEnvVarsToSwitchOp(value env.Variables) SwitchOpt {
	return func(op *SwitchOp) {
		op.vars = value
	}
}

func SetVersionToSwitchOp(version string) SwitchOpt {
	return func(op *SwitchOp) {
		version, err := versionutil.Normalize(version)
		if errutil.AssignIfErr(&op.err, err, nil) {
			return
		}

		op.version = version
	}
}
