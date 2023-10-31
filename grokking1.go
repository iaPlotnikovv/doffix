package main

import (
	"fmt"
)

func sum(n []int) int { // рекурсионная сумма массива

	if len(n) == 0 {
		return 0
	}
	return n[0] + sum(n[1:])
}

func count(n []int) (x int) { // подсчет элементов в массиве

	if len(n) == 0 {
		return 0
	}
	x++
	return x + count(n[1:])
}

func big(n []int) (item int) { // поиск наибольшего числа

	if len(n) == 0 {
		return 0
	}
	if n[0] > big(n[1:]) {

		item = n[0]

	} else {

		item = big(n[1:])
	}
	return item
}

func binary(n []int, item int) int { // бинарный поиск

	mid := (len(n) - 1) / 2

	if n[mid] > item {

		return binary(n[0:mid], item)

	} else if n[mid] < item {

		return binary(n[mid+1:], item)

	}

	fmt.Println(mid)
	return n[mid]
}
