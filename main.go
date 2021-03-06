package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func Hello() string {
	return "Hello release"
}

func main() {
	_, err := fmt.Println(Hello())
	if err != nil {
		println(errors.Cause(err))
	}
}
