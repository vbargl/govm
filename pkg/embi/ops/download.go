package ops

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"barglvojtech.net/govm/pkg/embi/env"
	"barglvojtech.net/govm/pkg/embi/types"
	"barglvojtech.net/govm/pkg/internal/archiveutil"
	"barglvojtech.net/govm/pkg/internal/optionutil"
	"barglvojtech.net/govm/pkg/internal/versionutil"
	"barglvojtech.net/x/pkg/errutil"
	"github.com/schollz/progressbar/v3"
)

type DownloadOpt func(*DownloadOp)

func NewDownloadOp(options ...DownloadOpt) *DownloadOp {
	var b optionutil.Builder[DownloadOp, DownloadOpt]
	b.AddGiven(options...)
	b.AddDefaults(
		SetEnvVarsToDownloadOp(env.Default()),
		SetOutputWriterToDownloadOp(os.Stdout),
	)

	var op = &DownloadOp{}
	b.Build(op)
	return op
}

type DownloadOp struct {

	// config variables

	vars         env.Variables
	version      types.Version
	outputWriter io.Writer

	// process variables

	err error // err can be setup by options

	cacheFile, versionDir string

	url  string
	resp *http.Response
}

func (op *DownloadOp) Process() error {
	if op.err != nil {
		return op.err
	}

	op.constructUrl()
	op.openRequest()
	op.createDirs()
	op.downloadFile()
	op.extractFile()
	return op.err
}

func (op *DownloadOp) constructUrl() {
	if op.err != nil {
		return
	}

	var err error

	op.url, err = versionutil.Url(op.vars.VersionUrl, types.QualifiedVersion{Version: op.version})
	errutil.AssignIfErr(&op.err, err, errutil.PrefixWith("could not create url"))
}

func (op *DownloadOp) openRequest() {
	if op.err != nil {
		return
	}

	resp, err := http.Get(op.url)
	if errutil.AssignIfErr(&op.err, err, errutil.PrefixWith("could not open request")) {
		return
	}

	switch resp.StatusCode {
	case 200:
		op.resp = resp
	case 404:
		op.err = fmt.Errorf("could not found version %s", op.version.String())
	default:
		op.err = fmt.Errorf("unexpected status returned: %s", resp.Status)
	}
}

func (op *DownloadOp) createDirs() {
	if op.err != nil {
		return
	}

	var err error

	cacheDir := op.vars.CacheDir()
	op.cacheFile = filepath.Join(cacheDir, path.Base(op.url))
	op.versionDir = filepath.Join(op.vars.VersionsDir(), op.version.String())

	err = os.MkdirAll(cacheDir, 0755)
	if errutil.AssignIfErr(&op.err, err, errutil.PrefixWithFormatted("could not create directory %s", cacheDir)) {
		return
	}

	err = os.MkdirAll(op.versionDir, 0755)
	if errutil.AssignIfErr(&op.err, err, errutil.PrefixWithFormatted("could not create directory %s", op.versionDir)) {
		return
	}
}

func (op *DownloadOp) downloadFile() {
	if op.err != nil {
		return
	}

	var (
		err  error
		file *os.File
	)

	file, err = os.OpenFile(op.cacheFile, os.O_CREATE|os.O_RDWR, 0644)
	if errutil.AssignIfErr(&op.err, err, errutil.PrefixWithFormatted("could not create file %s", op.cacheFile)) {
		return
	}
	defer file.Close()

	pb := progressbar.Default(op.resp.ContentLength, op.version.String())
	_, err = io.Copy(io.MultiWriter(file, pb), op.resp.Body)
	if errutil.AssignIfErr(&op.err, err, errutil.PrefixWith("error occured during download")) {
		return
	}
}

func (op *DownloadOp) extractFile() {
	if op.err != nil {
		return
	}

	err := archiveutil.Extract(op.cacheFile, op.versionDir)
	errutil.AssignIfErr(&op.err, err, errutil.PrefixWithFormatted("could not extract file %s", op.cacheFile))
}
