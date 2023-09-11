package urlext

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAbsUri(t *testing.T) {
	t.Run("file", func(t *testing.T) {
		uri := "./example/doc/adAdd.md"
		parentUri := ""
		basUri, _, err := AbsUri(uri, parentUri)
		require.NoError(t, err)
		fmt.Println(basUri)
	})
	t.Run("file_with_parent", func(t *testing.T) {
		uri := "../common.md"
		parentUri := "/Users/admin/Documents/go/dml/example/doc/adAdd.md"
		basUri, _, err := AbsUri(uri, parentUri)
		require.NoError(t, err)
		fmt.Println(basUri)
	})
	t.Run("github", func(t *testing.T) {
		uri := "git@github.com:suifengpiao14/apidml.git/example/doc/adList.md"
		parentUri := ""
		basUri, _, err := AbsUri(uri, parentUri)
		require.NoError(t, err)
		fmt.Println(basUri)
	})
	t.Run("github_with_parent", func(t *testing.T) {
		uri := "./adAdd.md"
		parentUri := "git@github.com:suifengpiao14/apidml.git/example/doc/adList.md"
		basUri, _, err := AbsUri(uri, parentUri)
		require.NoError(t, err)
		fmt.Println(basUri)
	})
	t.Run("ssh_with_parent", func(t *testing.T) {
		uri := "./adAdd.md"
		parentUri := "ssh://git@github.com/suifengpiao14/apidml.git/example/doc/adList.md"
		basUri, _, err := AbsUri(uri, parentUri)
		require.NoError(t, err)
		fmt.Println(basUri)
	})
}
