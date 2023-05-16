package main

import (
	"fmt"
)

func printMap(tmpMap map[string]string) {
	for k, v := range tmpMap {
		fmt.Printf("%s->%s\n", k, v)
	}
}

type hashMap map[string]string

func (tmpMap hashMap) consoleMap() {
	for k, v := range tmpMap {
		fmt.Printf("%s->%s\n", k, v)
	}
}

func maps() {
	fmt.Println("*******************maps********************")
	m := make(map[string]int)
	m["val1"] = 10
	m["val2"] = 11
	fmt.Println("map:", m)

	var myMap hashMap = make(map[string]string)
	var strings1 [5]string = [5]string{"amazon", "tesla", "microsoft", "apple", "netflix"}
	var strings2 [5]string = [5]string{"sofware engineer", "Senior software engineer", "member of technical staff", "sde2", "program manager"}

	for i := 0; i < len(strings1); i++ {
		myMap[strings1[i]] = strings2[i]
	}

	myMap.consoleMap()

	delete(myMap, "tesla")

	myMap.consoleMap()

}
