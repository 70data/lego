package util

import (
	"fmt"
	goruntime "runtime"
	"strings"
	"time"
)

// TimeElapsed prints time it takes to execute a function
// Usage: defer TimeElapsed("function-name")()
func TimeElapsed(functionName string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", functionName, time.Since(start))
	}
}

func GetFuncName() string {
	p, _, _, _ := goruntime.Caller(1)
	tmp := strings.Split(goruntime.FuncForPC(p).Name(), "/")
	funcName := tmp[len(tmp)-1]
	return funcName
}
