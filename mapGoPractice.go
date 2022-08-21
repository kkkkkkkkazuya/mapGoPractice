package main

import "fmt"

var slice = []string{"Ruby", "Java", "Go", "PHP", "JavaScript"}

func main() {
	var var1 int = 10
	var var2 *int = &var1
	fmt.Println(var2)
	var var3 *int
	fmt.Println(var3)

	mapEx := make(map[string]string, 2)
	fmt.Println(mapEx)
	mapEx["firstName"] = "Kazuya"
	mapEx["lastName"] = "Kojima"
	fmt.Println(mapEx)

	mapEx2 := map[string]int{"key": 1, "value": 2}

	fmt.Println(mapEx2["key"])

	for index, values := range slice {
		fmt.Println(index, values)
	}
}
