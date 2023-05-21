package main

import "fmt"

func rangeFunc() {
	nums := []int{2, 3, 4}
	sum := 0

	for _, num := range nums {
		sum += num
	}

	fmt.Println("sum:", sum)

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for key, val := range kvs {
		fmt.Println("key->val:", key, val)
	}

	for i, c := range "go" {
		fmt.Printf("%d->%d->%s\n", i, c, string(c))
	}

}
