package stack

import "fmt"

type myStruct struct {
	x int
}

func printSlice(v []any) {
	fmt.Println(v)
}

func main() {
	list := []myStruct{{1}, {2}}
	intSlice := make([]any, len(list))

	for i := range list {
		intSlice[i] = list[i]
	}

	printSlice(intSlice)
}
