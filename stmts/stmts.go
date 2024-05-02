package stmts

import "fmt"

func ShortDec() {
	eggs := 42
	fmt.Println(eggs)
	eggs = 50
	fmt.Println(eggs)
}

func Bigger(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Xor(x, y bool) bool {
	if x == y {
		return false
	}
	return true
}

func Greet(person string) string {
	switch person {
	case "Alice":
		return "Hey, Alice."
	case "Bob":
		return "What's up, Bob?"
	default:
		return "Hello, stranger."
	}
}

func Total(slice []int) int {
	var sum int
	for _, e := range slice {
		sum += e
	}
	return sum
}

func Evens() []int {
	var result []int
	for i := 0; i < 101; i++ {
		if i%2 == 0 {
			result = append(result, i)
		}
	}
	return result
}

func NonNegative(s []int) []int {
	var result []int
	for _, e := range s {
		if e < 0 {
			continue
		}
		fmt.Println(e)
		result = append(result, e)
	}
	return result
}
