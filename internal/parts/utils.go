package parts

import (
	"strings"
)

func getHash(s string) string {
	hash := strings.SplitN(s, "#", 2)
	if len(hash) != 2 {
		hash = append(hash, "")
	}
	return hash[1]
}

func getHostAndCredentials(s string) string {
	host := strings.SplitN(s, "//", 2)
	if strings.Contains(host[1], "/") {
		host = strings.SplitN(host[1], "/", 2)
		return host[0]
	}
	return host[1]
}

func getHost(s string) string {
	if strings.Contains(s, "@") {
		h := strings.SplitN(s, "@", 2)
		return h[1]
	}
	return s
}

func getCredentials(s string) [2]string {
	if strings.Contains(s, "@") {
		h := strings.SplitN(s, "@", 2)
		if strings.Contains(h[0], ":") {
			h = strings.SplitN(h[0], ":", 2)
			return [2]string{h[0], h[1]}
		}
	}
	return [2]string{"", ""}
}

func getHostName(s string) string {
	return strings.SplitN(s, ":", 2)[0]
}

func getPort(s string) string {
	if strings.Contains(s, ":") {
		h := strings.SplitN(s, ":", 2)
		return trim(h[1], []string{"/", "?", "#"})
	}
	return ""
}

func getProtocol(s string) string {
	if strings.Contains(s, "://") {
		h := strings.SplitN(s, "://", 2)
		return h[0]
	}
	return ""
}

func getPathname(s string) string {
	h := make([]string, 2)
	if strings.Contains(s, "://") {
		h = strings.SplitN(s, "://", 2)
	}

	if strings.Contains(h[1], "/") {
		h = strings.SplitN(h[1], "/", 2)
		p := trim(h[1], []string{"?", "#"})
		return p
	}

	return ""
}

func getSearch(s string) string {
	h := make([]string, 2)
	if strings.Contains(s, "://") {
		h = strings.SplitN(s, "://", 2)
	}

	if strings.Contains(h[1], "?") {
		h = strings.SplitN(h[1], "?", 2)
		return "?" + trim(h[1], []string{"#"})
	}

	return ""
}

func trim(s string, d []string) string {
	for _, v := range d {
		if strings.Contains(s, v) {
			cut := strings.SplitN(s, v, 2)
			s = cut[0]
		}
	}
	return s
}
