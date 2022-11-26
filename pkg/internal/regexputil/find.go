package regexputil

import "regexp"

func FindString(pattern, s string) (string, error) {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return "", err
	}
	return re.FindString(s), nil
}
