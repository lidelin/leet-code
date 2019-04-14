package utils

import (
	"fmt"
	"testing"
)

func TestTree(t *testing.T) {
	root := MakeTree([]int{10, 3, 21, 7, 55, 4, 12})

	fmt.Println(PreOrderTraversal(root))
	fmt.Println(InOrderTraversal(root))
	fmt.Println(PostOrderTraversal(root))
}
