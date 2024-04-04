package gourl

import (
	"errors"
	"github.com/bndrmrtn/gourl/internal/parts"
	"regexp"
	"strings"
)

var regex = regexp.MustCompile(`^(https?|s?ftp|wss?)://[^\s/$.?#].\S*$`)

type GoURL struct {
	u *parts.UrlParts
}

func New(url string, base ...string) (*GoURL, error) {
	if is(base) {
		url = strings.Join([]string{
			strings.TrimSuffix(base[0], "/"),
			strings.TrimPrefix(url, "/"),
		}, "/")
	}

	if err := Check(url); err != nil {
		return nil, err
	}

	return &GoURL{
		u: parts.Parse(url),
	}, nil
}

func (g *GoURL) Hash(set ...string) string {
	if is(set) {
		g.u.Hash = strings.TrimPrefix(set[0], "#")
	}
	return g.u.Hash
}

func (g *GoURL) Host(set ...string) string {
	if is(set) {
		g.u.Host = set[0]
		if strings.Contains(set[0], ":") {
			sp := strings.SplitN(set[0], ":", 2)
			g.u.Hostname = sp[0]
			g.u.Port = sp[1]
		}
	}
	return g.u.Host
}

func (g *GoURL) Hostname(set ...string) string {
	if is(set) {
		g.u.Hostname = set[0]
		g.u.Host = set[0]
		if g.u.Port != "" {
			g.u.Host += ":" + g.u.Port
		}
	}
	return g.u.Hostname
}

func (g *GoURL) Password(set ...string) string {
	if is(set) {
		g.u.Password = set[0]
	}
	return g.u.Password
}

func (g *GoURL) PathName(set ...string) string {
	if is(set) {
		g.u.Pathname = strings.TrimPrefix(set[0], "/")
	}
	return g.u.Pathname
}

func (g *GoURL) Port(set ...string) string {
	if is(set) {
		g.u.Port = strings.TrimPrefix(set[0], ":")
		if g.u.Port != "" {
			g.u.Host = g.u.Hostname + ":" + g.u.Port
		} else {
			g.u.Host = g.u.Hostname
		}
	}
	return g.u.Port
}

func (g *GoURL) Protocol(set ...string) string {
	if is(set) {
		g.u.Protocol = set[0]
	}
	return g.u.Protocol
}

func (g *GoURL) Search(set ...map[string]string) string {
	if len(set) > 0 {
		g.u.Search = "?"
		for k, v := range set[0] {
			g.u.Search += k + "=" + v + "&"
		}
		g.u.Search = strings.TrimSuffix(g.u.Search, "&")
	}
	return g.u.Search
}

func (g *GoURL) Username(set ...string) string {
	if is(set) {
		g.u.Username = set[0]
	}
	return g.u.Username
}

func (g *GoURL) Href() string {
	return g.u.String()
}

func (g *GoURL) SafeHref() (string, error) {
	if err := g.Check(); err != nil {
		return "", err
	}
	return g.Href(), nil
}

func (g *GoURL) Check() error {
	return Check(g.Href())
}

func Check(url string) error {
	if !regex.MatchString(url) {
		return errors.New("regex not match, invalid URL")
	}
	return nil
}
