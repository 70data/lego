package utils

import (
	"strconv"
	"strings"
)

func StringArrayToString(naiveArray []string) string {
	var naiveList string
	for _, v := range naiveArray {
		if naiveList == "" {
			naiveList = v
		} else {
			naiveList = strings.Join([]string{naiveList, v}, ",")
		}
	}
	return naiveList
}

func U2S(us []uint8) string {
	var str []string
	var target string
	for _, b := range us {
		if (int(b)-int(b)%16)/16 >= 10 {
			switch (int(b) - int(b)%16) / 16 {
			case 10:
				str = append(str, "a")
			case 11:
				str = append(str, "b")
			case 12:
				str = append(str, "c")
			case 13:
				str = append(str, "d")
			case 14:
				str = append(str, "e")
			case 15:
				str = append(str, "f")
			}
		} else {
			str = append(str, strconv.Itoa((int(b)-int(b)%16)/16))
		}
		if int(b)%16 >= 10 {
			switch int(b) % 16 {
			case 10:
				str = append(str, "a")
			case 11:
				str = append(str, "b")
			case 12:
				str = append(str, "c")
			case 13:
				str = append(str, "d")
			case 14:
				str = append(str, "e")
			case 15:
				str = append(str, "f")
			}
		} else {
			str = append(str, strconv.Itoa(int(b)%16))
		}
	}
	for i, _ := range str {
		target += str[i]
	}
	return target
}
