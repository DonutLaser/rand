package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(filePath string) (result []string, success bool) {
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0755)
	if err != nil {
		fmt.Println(err)
		return []string{}, false
	}

	var scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	return result, true
}
