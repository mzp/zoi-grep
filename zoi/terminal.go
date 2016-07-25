package zoi

import (
	"encoding/base64"
	"os"
	"strings"
)

var term = os.Getenv("TERM")

func osc() string {
	if strings.Contains(term, "screen") {
		return "\033Ptmux;\033\033]"
	} else {
		return "\033]"
	}
}

func st() string {
	if strings.Contains(term, "screen") {
		return "\a\033\\"
	} else {
		return "\a"
	}
}

func encode(image []byte) string {
	return base64.StdEncoding.EncodeToString(image)
}

func InlineImage(image []byte) string {
	s := ""

	s += osc()
	s += "1337;File="
	s += ";inline=1:"
	s += encode(image)
	s += st()

	return s
}
