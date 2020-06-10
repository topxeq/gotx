package main

func div(a, b int) (result int) {
	defer func() {

		// fmt.Printf("r = %#v\n", r)

		if r := recover(); r != nil {
			result = 0
		}
	}()

	return a / b
}

func main() {
	println(div(30, 2))

}
