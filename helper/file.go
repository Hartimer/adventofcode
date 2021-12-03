package helper

import (
	"bufio"
	"os"
)

func FileLineReader(inputFilePath string) chan string {
	c := make(chan string)
	go func() {
		file, err := os.Open(inputFilePath)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		defer close(c)
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			c <- scanner.Text()
		}
	}()
	return c
}
