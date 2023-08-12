package main

import (
	"fmt"
	"os"
)

func main() {
	fileName := "data.txt"
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	for i := 1; i <= 30; i++ {
		line := fmt.Sprintf("%d\n", i)
		_, err := file.WriteString(line)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}

	fmt.Println("File created:", fileName)
}
