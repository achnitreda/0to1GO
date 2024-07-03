package methods

import (
	"bufio"
	"errors"
	"log"
	"os"
)

var errTemplates = errors.New("INVALID TEMPLATE, use those:\n-->  shadow\n-->  standard\n-->  thinkertoy")

func ProcessBanner(fileName, input string) map[byte][]string {
	file, err := os.Open(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal(errTemplates)
		} else {
			log.Fatal(err)
		}
	}
	defer file.Close()

	chars := CharsNeeded(input)
	scanner := bufio.NewScanner(file)
	graphics := make(map[byte][]string)
	lines := 0
	for i := 0; i < len(chars); i++ {
		if chars[i] != 0 {
			temp := make([]string, 8)
			for scanner.Scan() {
				lines += 1
				if lines == ((int(chars[i])-32)*9 + 2) {
					for j := 0; j < 8; j++ {
						temp[j] = scanner.Text()
						scanner.Scan()
						lines += 1
					}
					graphics[chars[i]] = temp
					break
				}
			}
			if err := scanner.Err(); err != nil {
				log.Fatal(err)
			}
		}
	}
	return graphics
}

func CharsNeeded(str string) []byte {
	slice := make([]byte, 95)
	if ValidInput(str) {
		for i := 0; i < len(str); i++ {
			if slice[str[i]-32] == 0 {
				slice[str[i]-32] = str[i]
			}
		}
	}
	return slice
}
