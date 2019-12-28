package time

import (
	"strconv"
	"time"
)

// Unix is time for unix
func Unix() string {
	naiveTime := time.Now().Unix()
	naiveTimeString := strconv.FormatInt(naiveTime, 10)
	return naiveTimeString
}
