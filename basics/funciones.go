package main

import "fmt"

func add(n1 int, n2 int) int {
	fmt.Println("Add function")
	return n1 + n2

}

func returnTwoParams(name, pass string) (bool, string) {
	return true, name
}

func main() {
	fmt.Println("result", add(4, 5))
	sucessLogin, user := returnTwoParams("dfz", "123")

	fmt.Printf("successLogin %d, user: %d", sucessLogin, user)
}
