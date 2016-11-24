package kit

import "testing"

import "fmt"

func TestFirstLower(t *testing.T) {
	data1 := "HelloWorld"
	fmt.Printf("%v\n", FirstLower(data1))

	data2 := "世界, hello"
	fmt.Printf("%v\n", FirstLower(data2))
}
