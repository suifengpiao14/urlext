package urlext

import "net/url"

func IsAbsHttpUri(uri string) (yes bool) {
	if len(uri) > 4 && uri[:4] == "http" {
		return true
	}
	return false
}
func AbsUri(uri string, parentAbs string) (absUri string, fragment string, err error) {
	u, err := ParseAddress(uri)
	if err != nil {
		return "", "", err
	}
	if u.IsAbs() {
		absUri = u.String()
		fragment = u.Fragment
		return absUri, fragment, nil
	}
	if parentAbs == "" {
		return AbsFile(uri, parentAbs)
	}
	abs, err := ParseAddress(parentAbs)
	if err != nil {
		return "", "", err
	}
	if len(abs.Scheme) <= 1 { // http，ssh 等协议，一般大于2个字符，window路径一般为1个、linux为空
		return AbsFile(uri, parentAbs)
	}

	nu := abs.ResolveReference(u)
	fragment = nu.Fragment
	nu.RawQuery = ""
	nu.Fragment = ""
	absUri = nu.String()
	return absUri, fragment, nil
}

func ParseAddress(address string) (u *url.URL, err error) {
	u, err = url.Parse(address)
	if err == nil {
		return u, nil
	}
	u, err = detectSSH(address)
	if err != nil {
		return u, err
	}
	return u, nil
}
