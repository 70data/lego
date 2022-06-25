package util

import (
	"fmt"
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
