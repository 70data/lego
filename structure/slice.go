package structure

import (
	"regexp"
	"strings"
)

func SliceToString(slice []string) string {
	var naiveList string
	for _, v := range slice {
		if naiveList == "" {
			naiveList = v
		} else {
			naiveList = strings.Join([]string{naiveList, v}, ",")
		}
	}
	return naiveList
}

func InterfaceSliceToString(slice []interface{}) string {
	var naiveList string
	for _, v := range slice {
		if naiveList == "" {
			naiveList = v.(string)
		} else {
			naiveList = strings.Join([]string{naiveList, v.(string)}, ",")
		}
	}
	return naiveList
}

func ContainsStringFromSlice(slice []string, s string) bool {
	for _, item := range slice {
		if item == s {
			return true
		}
	}
	return false
}

func JudgeStrInStringSlice(s string, slice []string) bool {
	for _, i := range slice {
		if i == s {
			return true
		}
	}
	return false
}

func RemoveStringFromSlice(slice []string, s string) []string {
	var result []string
	for _, item := range slice {
		if item == s {
			continue
		}
		result = append(result, item)
	}
	return result
}

func RemoveDuplicatesAndEmptyFromSlice(slice []string) []string {
	// < 1024 use slice
	if len(slice) < 1024 {
		return removeRepByLoop(slice)
	}
	// < 1024 use map
	return removeRepByMap(slice)
}

// RemoveDuplicatesAndEmptyFromSlice2 is delete duplicate data from slice
func RemoveDuplicatesAndEmptyFromSlice2(slice []string) (ret []string) {
	basicArrayLen := len(slice)
	for i := 0; i < basicArrayLen; i++ {
		if (i > 0 && slice[i-1] == slice[i]) || len(slice[i]) == 0 {
			continue
		}
		ret = append(ret, slice[i])
	}
	return
}

func removeRepByLoop(slice []string) []string {
	// mark data
	var result []string
	for i := range slice {
		// find \s
		reg := regexp.MustCompile(`[^s]+`)
		r := reg.FindAllString(slice[i], -1)
		if len(r) != 0 {
			flag := true
			for j := range result {
				if slice[i] == result[j] {
					// mark duplicate data and mark false
					flag = false
					break
				}
			}
			// false not to data
			if flag {
				result = append(result, slice[i])
			}
		}
	}
	return result
}

func removeRepByMap(slice []string) []string {
	var result []string
	// mark data
	tempMap := map[string]byte{}
	for _, e := range slice {
		// find \s
		reg := regexp.MustCompile(`[^s]+`)
		r := reg.FindAllString(e, -1)
		if len(r) != 0 {
			l := len(tempMap)
			tempMap[e] = 0
			// len metamorphic is not duplicate
			if len(tempMap) != l {
				result = append(result, e)
			}
		}
	}
	return result
}
