package zoi

import "strings"

func Filter(query []string, xs []Entry) []Entry {
	results := []Entry{}

	for _, v := range xs {
		matched := true

		for _, q := range query {
			if !strings.Contains(v.Script, q) {
				matched = false
			}
		}

		if matched {
			results = append(results, v)
		}
	}

	return results
}
