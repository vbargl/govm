package versionutil

import (
	"reflect"
	"testing"

	"barglvojtech.net/govm/pkg/embi/types"
	"golang.org/x/exp/slices"
)

func TestCompare(t *testing.T) {
	param := []types.Version{
		"go1.3.1",
		"go1.3rc2",
		"go1.3rc3",
		"go1.3.2",
		"go1.3beta1",
		"go1.3",
		"go1.3rc1",
	}

	expected := []types.Version{
		"go1.3.2",
		"go1.3.1",
		"go1.3",
		"go1.3rc3",
		"go1.3rc2",
		"go1.3rc1",
		"go1.3beta1",
	}

	slices.SortFunc(param, Compare)
	if !reflect.DeepEqual(param, expected) {
		t.Errorf("expected %v, got %v", expected, param)
	}
}
