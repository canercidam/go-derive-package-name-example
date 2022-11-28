package main

import (
	"fmt"
	"foo/baz"
	"strings"
)

func main() {
	packagePath := baz.PackagePath()

	// outputs "foo/baz.PackagePath"
	fmt.Println("original:", packagePath)

	parts := strings.Split(packagePath, ".")
	if len(parts) == 2 {
		packagePath = parts[0]
	}

	// outputs "foo/baz"
	fmt.Println("fixed:", packagePath)
}
