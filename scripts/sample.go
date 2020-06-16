package main

import (
	"fmt"
	"time"
)

func main() {

	fn1 := func(args ...*time.Duration) string {
		rs := fmt.Sprintf("%v", args[0])

		return rs
	}

	fmt.Printf("%#v\n", fn1)

}
