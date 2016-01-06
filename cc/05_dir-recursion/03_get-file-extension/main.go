package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	filepath.Walk("../../", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		// ext is empty if no file extension
		ext := filepath.Ext(path)
		fmt.Println(ext)
		return nil
	})
}
