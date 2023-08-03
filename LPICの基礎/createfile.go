package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 1; i <= 7; i++ {
		fileName := fmt.Sprintf("%d日目.md", i)
		fileContent := fmt.Sprintf("# Day %d\n\nWrite your content here for day %d.", i, i)

		err := os.WriteFile(fileName, []byte(fileContent), 0644)
		if err != nil {
			fmt.Println("Error creating file:", err)
		} else {
			fmt.Println("File created:", fileName)
		}
	}
}
