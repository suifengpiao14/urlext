package urlext

import (
	"strings"

	"github.com/go-git/go-git/v5"
)

func IsAbsGitUri(uri string) (yes bool) {
	yes = strings.Contains(uri, git.GitDirName)
	return yes
}
