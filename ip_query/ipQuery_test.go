package ipQuery

import (
	"fmt"
	"testing"
)

func Test_ipQuery(t *testing.T) {
	LoadIPdata("IP.txt")
	position := ProvinceIsp("114.114.114.114")
	fmt.Printf("%s\n%s\n%s\n%s\n", position[0], position[1], position[2], position[3])
}
func Test_MakeStr(t *testing.T) {
	position := MakeStr("114.114.114.114")
	fmt.Println(position)
}
