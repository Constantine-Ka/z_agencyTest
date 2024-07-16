package adapters

import (
	"strconv"
	"strings"
)

func StrToIntSLice(src string) []int {
	var out []int
	for _, s := range strings.Split(src, ",") {
		num, _ := strconv.Atoi(strings.TrimSpace(s))
		out = append(out, num)
	}
	return out
}
