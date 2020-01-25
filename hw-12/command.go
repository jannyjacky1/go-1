package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("start")

	a, ok := os.LookupEnv("A_VAR")
	if !ok {
		fmt.Println("a not ok")
		os.Exit(1)
	}

	b, ok := os.LookupEnv("B_VAR")
	if !ok {
		fmt.Println("b not ok")
		os.Exit(1)
	}

	sign, ok := os.LookupEnv("SIGN_ENV")
	if !ok {
		fmt.Println("sign not ok")
		os.Exit(1)
	}

	fmt.Println(a + sign + b)
}
