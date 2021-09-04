package process

// removeDuplicatesAndEmpty is delete duplicate data from slice.
func removeDuplicatesAndEmpty(basicArray []string) (ret []string) {
	basicArrayLen := len(basicArray)
	for i := 0; i < basicArrayLen; i++ {
		if (i > 0 && basicArray[i-1] == basicArray[i]) || len(basicArray[i]) == 0 {
			continue
		}
		ret = append(ret, basicArray[i])
	}
	return
}
