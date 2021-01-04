package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	path := filepath.Join("a", "b")
	fmt.Println(path)
}
