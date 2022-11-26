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

	lMajorA, lMajorB := len(partsA.major), len(partsB.major)
	majorDiff := strings.Compare(partsA.major, partsB.major)
	switch {
	case lMajorA != lMajorB:
		return lMajorB - lMajorA
	case majorDiff != 0:
		return majorDiff
	}

	lMinorA, lMinorB := len(partsA.minor), len(partsB.minor)
	minorDiff := strings.Compare(partsA.minor, partsB.minor)
	switch {
	case lMinorA != lMinorB:
		return lMinorB - lMinorA
	case minorDiff != 0:
		return -minorDiff
	}

	lBuildA, lBuildB := len(partsA.build), len(partsB.build)
	buildDiff := strings.Compare(partsA.build, partsB.build)
	switch {
	case lBuildA != lBuildB:
		return lBuildB - lBuildA
	case buildDiff != 0:
		return -buildDiff
	}

	lMarkA, lMarkB := len(partsA.mark), len(partsB.mark)
	switch {
	// return the one without mark first
	case lMarkA == 0 && lMarkB != 0:
		return -1
	case lMarkA != 0 && lMarkB == 0:
		return 1

	// return rc before beta
	case strings.HasPrefix(partsA.mark, "rc") && strings.HasPrefix(partsB.mark, "beta"):
		return -1
	case strings.HasPrefix(partsA.mark, "beta") && strings.HasPrefix(partsB.mark, "rc"):
		return 1

	// return the one with bigger number
	default:
		return -strings.Compare(partsA.mark, partsB.mark)
	}
}
