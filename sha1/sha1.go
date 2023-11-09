package main

import "os"

func sha1(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {

	}

	defer file.Close()
	return "", nil
}

func main() {
}
