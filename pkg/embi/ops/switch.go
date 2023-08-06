package ops

import (
	"fmt"
	"go/build"
	"log"
	"os"
	"path/filepath"

	"barglvojtech.net/govm/pkg/embi/env"
	"barglvojtech.net/govm/pkg/embi/types"
	"barglvojtech.net/govm/pkg/internal/executil"
	"barglvojtech.net/govm/pkg/internal/fsutil"
	"barglvojtech.net/govm/pkg/internal/optionutil"

	"barglvojtech.net/x/pkg/errutil"
)

type SwitchOpt func(*SwitchOp)

func NewSwitchOp(options ...SwitchOpt) *SwitchOp {
	var b optionutil.Builder[SwitchOp, SwitchOpt]
	b.AddGiven(options...)
	b.AddDefaults(
		SetEnvVarsToSwitchOp(env.Default()),
	)

	var op = &SwitchOp{}
	b.Build(op)
	return op
}

type SwitchOp struct {

	// config variables

	vars    env.Variables
	version types.Version

	// process variables

	err error

	versionGoExecutable string
}

func (op *SwitchOp) Process() error {
	if op.err != nil {
		return op.err
	}

	op.verifyExistanceOfVersion()
	op.relinkGoExecutable()
	return op.err
}

func (op *SwitchOp) verifyExistanceOfVersion() {
	if op.err != nil {
		return
	}

	goExec := "go"
	if ext := executil.Ext(build.Default.GOOS); ext != "" {
		goExec += "."
		goExec += ext
	}

	versionGoExecutable := filepath.Join(op.vars.VersionsDir(), op.version.String(), "bin", goExec)
	if !fsutil.Exists(versionGoExecutable) {
		op.err = fmt.Errorf("version %s not found in runtime", op.version)
		return
	}

	log.Printf("version %s found", op.version)
	op.versionGoExecutable = versionGoExecutable
}

func (op *SwitchOp) relinkGoExecutable() {
	if op.err != nil {
		return
	}

	goExec := op.vars.GoExecutable()

	if fsutil.Exists(goExec) {
		log.Printf("go executable already exists, overwriting it with %s", op.versionGoExecutable)
		err := os.Remove(goExec)
		errutil.AssignIfErr(&op.err, err, errutil.PrefixWithFormatted("could not remove %s", goExec))
	} else {
		log.Printf("creating symlink to %s", op.versionGoExecutable)
	}

	err := os.Symlink(op.versionGoExecutable, goExec)
	errutil.AssignIfErr(&op.err, err, errutil.PrefixWithFormatted("could not symlink version %s into %s", op.versionGoExecutable, goExec))
}
