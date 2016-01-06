package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	filepath.Walk("sample-files-after/", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		fmt.Println(path)
		f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			panic(err)
		}

		defer f.Close()

		text := ` /* THIS WAS ADDED TO THE FILE
		and this is the second line
		AND THIS IS THE THIRD LINE
		*/`

		if _, err = f.WriteString(text); err != nil {
			panic(err)
		}
		return nil
	})
}
