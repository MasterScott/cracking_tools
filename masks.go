package main

import (
	"bufio"
	"fmt"
	"os"
	"log"
	"unicode"
)

func main() {

	input := "rockyou.txt"
	masks := make(map[string]int)
	fileHandle, err := os.Open(input)
	if err != nil {
		log.Fatal(input, " file not found")
	}
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)

	for fileScanner.Scan() {
		text := fileScanner.Text()
		mask := getMask(text)
		//fmt.Printf("%v,%v\n", text, mask)
		masks[mask] += 1
	}

	printTopMasks(masks, 10)
}

func getMask(text string) (mask string) {
	for _, char := range text {
		if unicode.IsDigit(char) {
			mask += "d"
		} else if unicode.IsUpper(char) {
			mask += "u"
		} else if unicode.IsLower(char) {
			mask += "l"
		} else {
			mask += "s"
		}
	}

	return mask
}

func printTopMasks(masks map[string]int, num int) {
	i := 0
	for mask, count := range masks {
		if i == num {
			break
		}
		fmt.Printf("%v,%v\n", mask, count)
		i++
	}
}
