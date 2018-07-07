package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func main() {
	_, err := fmt.Println("Hello release")
	if err != nil {
		println(errors.Cause(err))
	}
}
