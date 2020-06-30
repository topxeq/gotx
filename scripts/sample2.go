package main

import "fmt"

func main() {

	m := make(map[string]string, 0)

	fmt.Println(m)

	m["a"] = fmt.Sprintf("%v", 0.1)

	fmt.Println(m)

	m["b"] = string(fmt.Sprintf("%v", 0.1))

	fmt.Println(m)

}
