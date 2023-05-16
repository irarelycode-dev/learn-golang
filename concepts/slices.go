package main

import "fmt"

func slices() {
	b := make([]string, 3)
	fmt.Println("slice:", b)
	b[0] = "vignesh"
	b[1] = "pugazhendhi"

	fmt.Println("set:", b)
	fmt.Println("get:", b[2])

	b = append(b, "d")
	fmt.Println("append:", b)
	b = append(b, "tesla", "software engineer")
	fmt.Println("append two:", b)

	c := make([]string, len(b))
	copy(c, b)
	fmt.Println("c:", c)
	fmt.Println("check:", &b == &c)

	l := b[2:5]
	fmt.Println(l)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}

}
