package pkg

import (
	"bufio"
	"log"
	"os"
)

func ReadInput() []string {
	scanner := bufio.NewScanner(os.Stdin)
	var urls []string

	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		urls = append(urls, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln("Error: during scanning")
	}

	return urls
}