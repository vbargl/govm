package ops

import (
	"io"

	"barglvojtech.net/govm/pkg/embi/env"
	"barglvojtech.net/govm/pkg/internal/versionutil"
	"barglvojtech.net/x/pkg/errutil"
)

func SetEnvVarsToDownloadOp(value env.Variables) DownloadOpt {
	return func(op *DownloadOp) {
		op.vars = value
	}
}

func SetVersionToDownloadOp(version string) DownloadOpt {
	return func(op *DownloadOp) {
		var err error

		op.version, err = versionutil.Normalize(version)
		errutil.AssignIfErr(&op.err, err, nil)
	}
}

func SetOutputWriterToDownloadOp(value io.Writer) DownloadOpt {
	return func(op *DownloadOp) {
		op.outputWriter = value
	}
}
