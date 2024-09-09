// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")
	//fmt.Println(fibonacci(10))
	for i := 0; i <= 10; i++ {
		fmt.Println(fibonacci(i))
	}

}

func fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {

		return fibonacci(n-1) + fibonacci(n-2)
	}
}
