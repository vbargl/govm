package versionutil

import (
	"fmt"
	"strings"
)

func interpolate(s string, variables map[string]string) string {
	for variable, value := range variables {
		s = strings.ReplaceAll(s, fmt.Sprintf("${%s}", variable), value)
	}
	return s
}
