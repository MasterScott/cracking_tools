package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"unicode"
)

func main() {

	input := os.Args[1]
	wordchan = readFile(input)
	masks := make(map[string]int)

	sortedMasks := rankByCount(masks)
	printTopMasks(sortedMasks, 10)

	fmt.Printf("Processed %v words\n", count)
}

func readFile(infile string) <-chan []byte {
	var count int64 = 1
	out := make(chan []byte)
	fileHandle, err := os.Open(input)
	if err != nil {
		log.Fatal(input, " file not found")
	}
	defer fileHandle.Close()
	fileReader := bufio.NewReader(fileHandle)

	for {
		bytes, err := fileReader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Printf("Error processing %v", count)
			log.Printf("Error %v", err)
		}

		out <- bytes
		count += 1
	}
}

func getMask(in <-chan []byte) <-chan string {
	out := make(chan string)
	go func() {
		for text := range in {
			var mask string = ""
			for _, char := range text {
				char := rune(char)
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
			out <- mask
		}
		close(out)
	}
	return out
}

func printTopMasks(masks PairList, num int) {
	i := 0
	for _, pair := range masks {
		if i == num {
			break
		}
		fmt.Printf("%v,%v\n", pair.Key, pair.Value)
		i++
	}
}

// https://stackoverflow.com/questions/18695346/how-to-sort-a-mapstringint-by-its-values

func rankByCount(masks map[string]int) PairList {
	pl := make(PairList, len(masks))
	i := 0
	for k, v := range masks {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
