package main

import "fmt"

func main() {

	m := make(map[string]string, 0)

	println(m)

	m["a"] = fmt.Sprintf("%v", 0.1)

	println(m)

	m["b"] = string(fmt.Sprintf("%v", 0.1))

	println(m)

}
