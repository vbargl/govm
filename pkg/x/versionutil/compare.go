package versionutil

import (
	"strings"

	"barglvojtech.net/govm/pkg/embi/types"
)

const (
	versionPrefix = "go"
)

func Compare(a, b types.Version) int {
	rawA := strings.TrimPrefix(a.String(), versionPrefix)
	rawB := strings.TrimPrefix(b.String(), versionPrefix)

	partsA, errA := Parse(rawA)
	partsB, errB := Parse(rawB)

	if errA != errB {
		if errA != nil {
			return 1
		} else {
			return -1
		}
	}

	lMajorA, lMajorB := len(partsA.Major), len(partsB.Major)
	majorDiff := strings.Compare(partsA.Major, partsB.Major)
	switch {
	case lMajorA != lMajorB:
		return lMajorB - lMajorA
	case majorDiff != 0:
		return majorDiff
	}

	lMinorA, lMinorB := len(partsA.Minor), len(partsB.Minor)
	minorDiff := strings.Compare(partsA.Minor, partsB.Minor)
	switch {
	case lMinorA != lMinorB:
		return lMinorB - lMinorA
	case minorDiff != 0:
		return -minorDiff
	}

	lBuildA, lBuildB := len(partsA.Patch), len(partsB.Patch)
	buildDiff := strings.Compare(partsA.Patch, partsB.Patch)
	switch {
	case lBuildA != lBuildB:
		return lBuildB - lBuildA
	case buildDiff != 0:
		return -buildDiff
	}

	lMarkA, lMarkB := len(partsA.PreRelease), len(partsB.PreRelease)
	switch {
	// return the one without mark first
	case lMarkA == 0 && lMarkB != 0:
		return -1
	case lMarkA != 0 && lMarkB == 0:
		return 1

	// return rc before beta
	case strings.HasPrefix(partsA.PreRelease, "rc") && strings.HasPrefix(partsB.PreRelease, "beta"):
		return -1
	case strings.HasPrefix(partsA.PreRelease, "beta") && strings.HasPrefix(partsB.PreRelease, "rc"):
		return 1

	// return the one with bigger number
	default:
		return -strings.Compare(partsA.PreRelease, partsB.PreRelease)
	}
}
