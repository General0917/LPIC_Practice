package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 1; i <= 7; i++ {
		fileName := fmt.Sprintf("%d日目.md", i)
		fileContent := fmt.Sprintf("# Day %d\n\nWrite your content here for day %d.", i, i)

		// ファイルが既に存在するか確認
		if _, err := os.Stat(fileName); os.IsNotExist(err) {
			// ファイルが存在しない場合のみ作成
			err := os.WriteFile(fileName, []byte(fileContent), 0644)
			if err != nil {
				fmt.Println("Error creating file:", err)
			} else {
				fmt.Println("File created:", fileName)
			}
		} else {
			fmt.Println("File already exists, skipping:", fileName)
		}
	}
}
