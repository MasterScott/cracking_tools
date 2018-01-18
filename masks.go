package main

import (
	"bufio"
	"fmt"
	"os"
	"log"
	"unicode"
	"sort"
)

func main() {

	var count int64 = 0
	input := os.Args[1]
	masks := make(map[string]int)
	fileHandle, err := os.Open(input)
	if err != nil {
		log.Fatal(input, " file not found")
	}
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		text := fileScanner.Text()
		mask := getMask(text)
		//fmt.Printf("%v,%v\n", text, mask)
		masks[mask] += 1
		count += 1
	}

	sortedMasks := rankByCount(masks)
	printTopMasks(sortedMasks, 10)

	fmt.Printf("Processed %v words", count)
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

func rankByCount(masks map[string]int) PairList{
	pl := make(PairList, len(masks))
	i := 0
	for k,v := range masks {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}

type Pair struct {
	Key string
	Value int
}

type PairList []Pair

func (p PairList) Len() int { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int){ p[i], p[j] = p[j], p[i] }
