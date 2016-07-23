package zoi

import (
	"io/ioutil"
	"os"
	"path"
)

var root = os.Getenv("ZOI_ROOT")

func Resolve(name string) string {
	return path.Join(root, name)
}

func DefaultPanels() ([]Entry, error) {
	path := Resolve("data/panels.json")
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	return Parse(file)
}

func ReadImage(path string) ([]byte, error) {
	return ioutil.ReadFile(Resolve(path))
}
