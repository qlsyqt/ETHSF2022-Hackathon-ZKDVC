package utils

import (
	"net/url"
	"path"
)

func Concat(host string, components ...string) (string, error) {
	u, err := url.Parse(host)
	if err != nil {
		return "", err
	}
	for _, part := range components {
		u.Path = path.Join(u.Path, part)
	}

	s := u.String()
	return s, nil
}
