package ops

import (
	"fmt"
	"path/filepath"

	"barglvojtech.net/govm/pkg/embi/env"
	"barglvojtech.net/govm/pkg/embi/types"
	"barglvojtech.net/govm/pkg/internal/archiveutil"
	"barglvojtech.net/govm/pkg/internal/fsutil"
	"barglvojtech.net/govm/pkg/internal/optionutil"
	"barglvojtech.net/govm/pkg/internal/versionutil"
	"barglvojtech.net/x/pkg/errutil"
)

type ExtractOpt func(*ExtractOp)

func NewExtractOp(options ...ExtractOpt) *ExtractOp {
	var b optionutil.Builder[ExtractOp, ExtractOpt]
	b.AddGiven(options...)
	b.AddGiven(
		SetEnvVarsToExtractOp(env.Default()),
	)

	var op = &ExtractOp{}
	b.Build(op)
	return op
}

type ExtractOp struct {

	// config variables

	vars    env.Variables
	version types.Version

	// process varialbes

	err error

	cacheFile, versionDir string
}

func (op *ExtractOp) Process() error {
	if op.err != nil {
		return op.err
	}

	op.verifyFS()
	op.extractFile()

	return op.err
}

func (op *ExtractOp) verifyFS() {
	if op.err != nil {
		return
	}

	var err error

	qualified := versionutil.Qualify(types.QualifiedVersion{Version: op.version})
	filename, err := versionutil.Filename("", qualified)
	if errutil.AssignIfErr(&op.err, err, errutil.PrefixWithFormatted("could not get filename for %s", op.version)) {
		return
	}

	op.cacheFile = filepath.Join(op.vars.CacheDir(), filename)
	op.versionDir = filepath.Join(op.vars.VersionsDir(), op.version.String())

	if !fsutil.Exists(op.cacheFile) {
		op.err = fmt.Errorf("file %s not found in cache", op.cacheFile)
		return
	}
}

func (op *ExtractOp) extractFile() {
	if op.err != nil {
		return
	}

	err := archiveutil.Extract(op.cacheFile, op.versionDir)
	errutil.AssignIfErr(&op.err, err, errutil.PrefixWithFormatted("could not extract file %s", op.cacheFile))
}
