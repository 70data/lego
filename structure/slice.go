package structure

import (
	"regexp"
)

func ContainsStringFromSlice(slice []string, s string) bool {
	for _, item := range slice {
		if item == s {
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

func CompareIntSlice(currentList, expectedList []int) int {
	for _, i := range currentList {
		for index, j := range expectedList {
			if i == j {
				expectedList = append(expectedList[:index], expectedList[index+1:]...)
				break
			}
		}
	}
	return expectedList[0]
}

func removeRepByLoop(slc []string) []string {
	// mark data
	var result []string
	for i := range slc {
		// find \s
		reg := regexp.MustCompile(`[^s]+`)
		r := reg.FindAllString(slc[i], -1)
		if len(r) != 0 {
			flag := true
			for j := range result {
				if slc[i] == result[j] {
					// mark duplicate data and mark false
					flag = false
					break
				}
			}
			// false not to data
			if flag {
				result = append(result, slc[i])
			}
		}
	}
	return result
}

func removeRepByMap(slc []string) []string {
	var result []string
	// mark data
	tempMap := map[string]byte{}
	for _, e := range slc {
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

func RemoveDuplicatesAndEmpty(slc []string) []string {
	// < 1024 use slice
	if len(slc) < 1024 {
		return removeRepByLoop(slc)
	}
	// < 1024 use map
	return removeRepByMap(slc)
}

// RemoveDuplicatesAndEmpty2 is delete duplicate data from slice.
func RemoveDuplicatesAndEmpty2(basicArray []string) (ret []string) {
	basicArrayLen := len(basicArray)
	for i := 0; i < basicArrayLen; i++ {
		if (i > 0 && basicArray[i-1] == basicArray[i]) || len(basicArray[i]) == 0 {
			continue
		}
		ret = append(ret, basicArray[i])
	}
	return
}
