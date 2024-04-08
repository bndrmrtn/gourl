package query

import (
	"github.com/bndrmrtn/gourl/internal/parts"
	"strings"
)

type Builder struct {
	Parts   *parts.UrlParts
	queries map[string][]string
}

func New(parts *parts.UrlParts) *Builder {
	b := &Builder{
		Parts:   parts,
		queries: make(map[string][]string),
	}

	p := strings.Split(strings.TrimPrefix(b.Parts.Search, "?"), "&")

	for _, sp := range p {
		part := strings.SplitN(sp, "=", 2)
		if len(part) != 2 {
			b.queries[part[0]] = make([]string, 0)
		} else {
			key, val := part[0], part[1]
			if strings.HasPrefix(val, "[") && strings.HasSuffix(val, "]") {
				val = strings.TrimPrefix(val, "[")
				val = strings.TrimSuffix(val, "]")
				b.queries[key] = strings.Split(val, ",")
			} else {
				b.queries[key] = []string{val}
			}
		}
	}

	return b
}

func (b *Builder) Get(key string) string {
	val, ok := b.queries[key]
	if !ok {
		return ""
	}
	return val[0]
}

func (b *Builder) GetList(key string) []string {
	val, ok := b.queries[key]
	if !ok {
		return nil
	}
	return val
}

func (b *Builder) Set(key string, val string) {
	b.queries[key] = []string{val}
	b.update()
}

func (b *Builder) SetList(key string, val []string) {
	b.queries[key] = val
	b.update()
}

func (b *Builder) Delete(key string) {
	delete(b.queries, key)
	b.update()
}

func (b *Builder) All() map[string][]string {
	return b.queries
}

func (b *Builder) update() {
	b.Parts.Search = ""
	for k, v := range b.queries {
		b.Parts.Search += k

		if len(v) == 1 {
			b.Parts.Search += "=" + v[0]
		} else if len(v) > 1 {
			b.Parts.Search += "=["
			for _, j := range v {
				b.Parts.Search += j + ","
			}
			b.Parts.Search = strings.TrimSuffix(b.Parts.Search, ",") + "]"
		} else {
			b.Parts.Search += v[0]
		}

		b.Parts.Search += "&"
	}
	b.Parts.Search = "?" + strings.TrimSuffix(b.Parts.Search, "&")
}
