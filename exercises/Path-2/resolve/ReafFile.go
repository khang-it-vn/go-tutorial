package resolve

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(path string) {

	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("\nOpen file from path %s has error: %s ", path, err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	if scanner.Err() != nil {
		fmt.Printf("\n Read file has error: %s", scanner.Err())
		os.Exit(1)
	}
}
