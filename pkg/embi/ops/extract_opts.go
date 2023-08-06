package ops

import (
	"barglvojtech.net/govm/pkg/embi/env"
	"barglvojtech.net/govm/pkg/internal/versionutil"
	"barglvojtech.net/x/pkg/errutil"
)

func SetEnvVarsToExtractOp(value env.Variables) ExtractOpt {
	return func(op *ExtractOp) {
		op.vars = value
	}
}

func SetVersionToExtractOp(version string) ExtractOpt {
	return func(op *ExtractOp) {
		var err error

		op.version, err = versionutil.Normalize(version)
		errutil.AssignIfErr(&op.err, err, nil)
	}
}
