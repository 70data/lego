package convert

import (
	"strings"
)

func StringSliceToString(naiveSlice []string) string {
	var naiveList string
	for _, v := range naiveSlice {
		if naiveList == "" {
			naiveList = v
		} else {
			naiveList = strings.Join([]string{naiveList, v}, ",")
		}
	}
	return naiveList
}
