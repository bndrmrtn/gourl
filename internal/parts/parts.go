package parts

type UrlParts struct {
	Hash     string
	Host     string
	Hostname string
	Password string
	Pathname string
	Port     string
	Protocol string
	Search   string
	Username string
}

func Parse(s string) *UrlParts {
	hostAndCredentials := getHostAndCredentials(s)
	host := getHost(hostAndCredentials)
	credentials := getCredentials(hostAndCredentials)

	return &UrlParts{
		Hash:     getHash(s),
		Host:     host,
		Hostname: getHostName(host),
		Password: credentials[1],
		Pathname: getPathname(s),
		Port:     getPort(host),
		Protocol: getProtocol(s),
		Search:   getSearch(s),
		Username: credentials[0],
	}
}

func (u *UrlParts) String() string {
	var url string

	if u.Protocol != "" {
		url += u.Protocol + ":"
	}

	url += "//"

	if u.Username != "" || u.Password != "" {
		url += u.Username + ":" + u.Password + "@"
	}

	url += u.Host

	if u.Pathname != "" {
		url += "/" + u.Pathname
	}

	url += u.Search

	if u.Hash != "" {
		url += "#" + u.Hash
	}

	return url
}
