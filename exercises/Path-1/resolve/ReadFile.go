package resolve

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(path string) {
	file, err := os.Open(path)

	if err != nil {
		fmt.Printf("\nHas error when read file %s, detail: %v", path, err)
		os.Exit(1)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	if scanner.Err() != nil {
		fmt.Printf("\nHas error when print value of file detail: %v", err)
		return
	}

}
