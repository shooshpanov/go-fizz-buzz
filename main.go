package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		switch {
		case i%15 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}

//package main

//Version with Goroutine

// func main() {
// 	for out := range FizzBuzz(100) {
// 		fmt.Println(out)
// 	}
// }

// func FizzBuzz(amount int) <-chan string {

// 	out := make(chan string, amount)

// 	go func() {
// 		for i := 1; i <= amount; i++ {
// 			result := ""
// 			if i%3 == 0 {
// 				result += "Fizz"
// 			}
// 			if i%5 == 0 {
// 				result += "Buzz"
// 			}
// 			if i%15 == 0 {
// 				result += "FizzBuzz"
// 			}
// 			if result == "" {
// 				result = fmt.Sprintf("%v", i)
// 			}
// 			out <- result
// 		}
// 		close(out)
// 	}()

// 	return out
// }
