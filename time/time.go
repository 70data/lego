package time

import (
	"strconv"
	"time"
)

func Unix() string {
	naiveTime := time.Now().Unix()
	naiveTimeString := strconv.FormatInt(naiveTime, 10)
	return naiveTimeString
}
