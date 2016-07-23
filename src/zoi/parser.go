package zoi

import (
	"encoding/json"
	"io"
)

type data map[string]map[string]Entry

func toArray(xs map[string]Entry) []Entry {
	ys := []Entry{}

	for _, v := range xs {
		ys = append(ys, v)
	}

	return ys
}

func Parse(r io.Reader) ([]Entry, error) {
	dec := json.NewDecoder(r)
	var d data
	err := dec.Decode(&d)

	if err != nil {
		return nil, err
	} else {
		entries := toArray(d["_default"])
		Sort(entries)
		return entries, nil
	}
}
