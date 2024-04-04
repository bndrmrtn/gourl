package gourl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Simple(t *testing.T) {
	url, err := New("http://localhost/api/users?id=123")

	if err != nil {
		t.Fatal("Cannot create URL", err)
	}

	assert.Equal(t, "http://localhost/api/users?id=123", url.Href(), "Must be valid")

	url.Hash("footer")
	assert.Equal(t, "http://localhost/api/users?id=123#footer", url.Href(), "Hash must be added")

	url.Host("example.com:3000")
	assert.Equal(t, "http://example.com:3000/api/users?id=123#footer", url.Href(), "Host must be changed")
	assert.Equal(t, "example.com", url.Hostname(), "Hostname must be changed")
	assert.Equal(t, "3000", url.Port(), "Port must be changed")

	url.Protocol("https")
	assert.Equal(t, "https://example.com:3000/api/users?id=123#footer", url.Href(), "Protocol must be changed to secure")

	url.Hostname("api.example.com")
	assert.Equal(t, "https://api.example.com:3000/api/users?id=123#footer", url.Href(), "Hostname must be changed")

	url.Port("")
	assert.Equal(t, "", url.Port(), "Port must be removed")

	url.Search(map[string]string{"username": "john"})
	assert.Equal(t, "https://api.example.com/api/users?username=john#footer", url.Href(), "Search must be changed")

	url.Protocol("sftp")
	assert.Equal(t, "sftp://api.example.com/api/users?username=john#footer", url.Href(), "Protocol must be changed")
}

func Test_UsernameAndPassword(t *testing.T) {
	url, err := New("http://admin:pass@example.com")

	if err != nil {
		t.Fatal("Cannot create URL", err)
	}

	assert.Equal(t, "http://admin:pass@example.com", url.Href())

	assert.Equal(t, "admin", url.Username())
	assert.Equal(t, "pass", url.Password())

	url.Password("secret")
	assert.Equal(t, "http://admin:secret@example.com", url.Href())

	url.Username("john")
	assert.Equal(t, "john", url.Username())
}
