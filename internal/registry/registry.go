package registry

import (
	"encoding/json"
	"io"
)

type Registry struct {
	t map[string]string
}

type entry struct {
	Url  string `json:"url"`
	Path string `json:"path"`
}

func New() *Registry {
	return &Registry{
		t: make(map[string]string),
	}
}

func (r *Registry) Set(url, path string) {
	r.t[url] = path
}

func (r *Registry) JsonEncoder(w io.Writer) error {
	var entries []entry

	for k, v := range r.t {
		entries = append(entries, entry{Url: k, Path: v})
	}

	return json.NewEncoder(w).Encode(entries)
}
