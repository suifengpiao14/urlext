package urlext

import (
	"os"
	"path/filepath"
	"strings"
)

func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()

}

func AbsFile(filename string, parentAbs string) (absFile string, fragment string, err error) {
	filename, fragment = parseFilename(filename)
	if filepath.IsAbs(filename) {
		return filename, filename, nil
	}
	parentAbs, _ = parseFilename(parentAbs)
	parentAbs, err = filepath.Abs(parentAbs)
	if err != nil {
		return "", "", err
	}
	if !IsDir(parentAbs) { // 不是路径，取上级路径
		parentAbs = filepath.Dir(parentAbs)
	}
	fullname := filepath.Join(parentAbs, filename)
	absFile, err = filepath.Abs(fullname)
	if err != nil {
		return "", "", err
	}
	return absFile, fragment, nil
}

func parseFilename(uri string) (newUri string, fragment string) {
	fragmentIndex := strings.Index(uri, "#")
	if fragmentIndex > -1 {
		fragment = uri[fragmentIndex+1:]
		uri = uri[:fragmentIndex]
		queryIndex := strings.Index(fragment, "?")
		if queryIndex > -1 {
			fragment = fragment[:queryIndex]
		}
	}
	queryIndex := strings.Index(uri, "?")
	if queryIndex > -1 {
		uri = uri[:queryIndex]
	}
	return uri, fragment
}
