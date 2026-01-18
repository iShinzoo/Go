package main

import (
	"fmt"
	"maps"
)

func main() {

	// maps -> hash, object, dictionary
	// maps are used to store key value pairs

	nums := make(map[string]string)

	nums["one"] = "key"
	nums["two"] = "backend"
	nums["three"] = "area"

	fmt.Println(nums["two"], nums["three"])
	fmt.Println(nums)

	// if key not present, returns zero value

	fmt.Println(len(nums))

	// deleting key
	delete(nums, "three")
	fmt.Println(nums)

	// to delete all keys
	clear(nums)
	fmt.Println(nums)

	num := map[string]int{
		"krish":   25,
		"alice":   30,
		"bob":     28,
		"charlie": 22,
	}
	fmt.Println(num)
	k, ok := num["krish"]
	fmt.Println(k)
	if ok {
		fmt.Println("all ok")
	} else {
		fmt.Println("not ok")
	}

	// maps fucntions
	map1 := map[int]string{
		1: "one",
		2: "two",
	}

	map2 := map[int]string{
		3: "three",
		4: "four",
	}

	fmt.Println(maps.Equal(map1, map2))
}
