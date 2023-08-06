package ops

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/exp/slices"

	"barglvojtech.net/govm/pkg/embi/types"
	"barglvojtech.net/govm/pkg/internal/optionutil"
	"barglvojtech.net/govm/pkg/x/versionutil"

	"barglvojtech.net/x/pkg/errutil"
)

const (
	// archiveKind "Archive" is kind govm is interested in.
	archiveKind = "Archive"
)

var (
	// listVersionsUrl is used for fetching and parsing all available versions of go.
	listVersionsUrl = url.URL{
		Scheme: "https",
		Host:   "go.dev",
		Path:   "dl",
	}
)

type ListOpt func(*ListOp)

func NewListOp(options ...ListOpt) *ListOp {
	var b optionutil.Builder[ListOp, ListOpt]
	b.AddGiven(options...)
	b.AddDefaults(
		setHttpClientToListOp(http.DefaultClient),
	)

	var op = &ListOp{}
	b.Build(op)
	return op
}

type ListOp struct {

	// configuration variables

	httpClient HttpClient

	// process variables

	err error

	res *http.Response
	doc *goquery.Document

	ver []types.VersionSet
}

func (op *ListOp) Process() error {
	if op.err != nil {
		return op.err
	}

	op.fetchHtml()
	op.parseHtml()
	op.findVersionNodes()
	return op.err
}

func (op *ListOp) Versions() []types.VersionSet {
	return op.ver
}

func (op *ListOp) fetchHtml() {
	if op.err != nil {
		return
	}

	var err error
	op.res, err = op.httpClient.Get(listVersionsUrl.String())
	errutil.AssignIfErr(&op.err, err, errutil.PrefixWith("could not fetch available versions"))
}

func (op *ListOp) parseHtml() {
	if op.err != nil {
		return
	}

	var err error
	defer op.res.Body.Close()
	op.doc, err = goquery.NewDocumentFromReader(op.res.Body)
	errutil.AssignIfErr(&op.err, err, errutil.PrefixWith("could not parse downloaded html"))
}

func (op *ListOp) findVersionNodes() {
	if op.err != nil {
		return
	}

	defer errutil.HandleIfRecovered(recover(), func(err error) {
		errutil.AssignIfErr(&op.err, err, errutil.PrefixWithFormatted("could not find any version on url '%v'", listVersionsUrl))
	})

	versions := make(map[string]*types.VersionSet)

	op.doc.Find("main article tr[class!=\"first\"]").Each(func(_ int, s *goquery.Selection) {
		dataElements := s.Children()

		kind := dataElements.Eq(kindColumn.Int()).Text()
		if kind != archiveKind {
			return
		}

		version := dataElements.Parent().Parent().Parent().Parent().Parent().Parent().AttrOr("id", "")

		// href contains absolute path to file, listVersionUrl is specified with /dl path.
		// Prefix '../' is there to ensure path referenced from root.
		link := listVersionsUrl.JoinPath("../" + dataElements.Eq(nameColumn.Int()).Children().First().AttrOr("href", ""))

		versionSet, ok := versions[version]
		if !ok {
			index := len(op.ver)
			op.ver = append(op.ver, types.VersionSet{
				Version: types.Version(version),
			})

			versionSet = &op.ver[index]
			versions[version] = &op.ver[index]
		}

		versionSet.Files = append(versionSet.Files, types.VersionFile{
			Version:  types.Version(version),
			Url:      link.String(),
			Goos:     strings.ToLower(dataElements.Eq(osColumn.Int()).Text()),
			Goarch:   strings.ToLower(dataElements.Eq(archColumn.Int()).Text()),
			Checksum: dataElements.Eq(checksumColumn.Int()).Text(),
		})
	})

	slices.SortFunc(op.ver, func(a, b types.VersionSet) int {
		return versionutil.Compare(a.Version, b.Version)
	})
}

type HttpClient interface {
	Get(url string) (*http.Response, error)
}

type versionTableColumn int

func (c versionTableColumn) Int() int {
	return int(c)
}

const (
	nameColumn     versionTableColumn = 0
	kindColumn     versionTableColumn = 1
	osColumn       versionTableColumn = 2
	archColumn     versionTableColumn = 3
	checksumColumn versionTableColumn = 4
)
