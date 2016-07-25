package zoi

import "sort"

type entries []Entry

func (x entries) Len() int {
	return len(x)
}

func (x entries) Less(i, j int) bool {
	return x[i].Path < x[j].Path
}

func (x entries) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

func Sort(es []Entry) {
	sort.Sort(entries(es))
}
