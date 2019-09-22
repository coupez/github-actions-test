package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("Hello github actions, I'm running on %s!!!\n", runtime.GOOS)
}
