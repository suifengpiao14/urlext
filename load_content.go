package urlext

import (
	"net/http"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"github.com/suifengpiao14/gitauto"
)

const (
	SOURCE_SCHEME_HTTP = "http"
	SOURCE_SCHEME_FILE = "file"
	SOURCE_SCHEME_GIT  = "git"
)

func LoadConent(absUri string) (b []byte, err error) {
	if IsAbsHttpUri(absUri) {
		b, err = loadFromURL(absUri)
	} else if IsAbsGitUri(absUri) {
		b, err = gitauto.ReadFile(absUri)
	} else {
		b, err = loadFromFile(absUri)
	}
	if err != nil {
		err = errors.WithMessagef(err, "uri:%s", absUri)
		return nil, err
	}
	return b, nil
}

func loadFromURL(url string) (b []byte, err error) {
	client := resty.New()
	resp, err := client.R().
		EnableTrace().
		Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != http.StatusOK {
		err := errors.Errorf("http code %d,body:%s", resp.StatusCode(), resp.Body())
		return nil, err
	}
	return resp.Body(), nil
}

func loadFromFile(file string) (b []byte, err error) {
	source, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	return source, nil
}
