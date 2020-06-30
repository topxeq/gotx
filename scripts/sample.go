package main

import (
	"fmt"
	"strconv"
)

func strToInt(strA string, defaultA int) int {
	n, err := strconv.ParseInt(strA, 10, 0)
	if err != nil {
		return defaultA
	}

	return int(n)
}

func StrToIntWithDefaultValue(strA string, defaultA int) int {
	nT, errT := strconv.ParseInt(strA, 10, 0)
	if errT != nil {
		return defaultA
	}

	return int(nT)
}

var r = 19

func main() {

	a := strToInt("30", -1)
	fmt.Printf("%#v\n", a)

	r := StrToIntWithDefaultValue("40", r)

	fmt.Println(r)

}
