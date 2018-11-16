package treesort

import "testing"
import "fmt"

func TestX(t *testing.T) {
	var root *tree
	slice := []int{2, 3, 1, 4}

	for _, e := range slice {
		root = add(root, e)
	}

	fmt.Println(root)
	Sort(slice)
	fmt.Println(slice)
}
