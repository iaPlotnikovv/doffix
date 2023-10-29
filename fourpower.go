package main

import (
	"fmt"
)

func power_four(n int) bool {

	if n > 0 && n%4 == 0 || n == 1 {
		if n == 1 {
			return true
		}
		return power_four(n / 4)
	} else {
		return false
	}

}

func main() {
	fmt.Println(power_four(64))
}
