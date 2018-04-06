package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func wc(str []string) map[string]int {
	counts := make(map[string]int)
	for _, word := range str {
		_, ok := counts[word]
		if ok {
			counts[word]++
		} else {
			counts[word] = 1
		}
	}
	return counts
}

func main() {
	byteData, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	data := string(byteData[:])
	replacer := strings.NewReplacer(",", "", ".", "", ";", "", "-", "", "?", "")
	data = replacer.Replace(data)
	words := strings.Fields(data)
	counts := wc(words)
	for count, element := range counts {
		fmt.Println(count, "=>", element)
	}
}
