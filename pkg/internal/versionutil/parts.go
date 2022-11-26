package versionutil

import (
	"strings"

	"barglvojtech.net/govm/pkg/embi/types"
)

type VersionParts struct {
	major string
	minor string
	build string
	mark  string
}

func (p *VersionParts) Version() types.Version {
	var b strings.Builder
	b.WriteString("go")
	b.WriteString(p.major)
	b.WriteRune('.')
	b.WriteString(p.minor)
	if p.build != "" {
		b.WriteRune('.')
		b.WriteString(p.build)
	}
	if p.mark != "" {
		b.WriteString(p.mark)
	}
	return types.Version(b.String())
}
