package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	filepath.Walk("../../../00_language-fundamentals/", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		fmt.Println(path)
		return nil
	})
}
