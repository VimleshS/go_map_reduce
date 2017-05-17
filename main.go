package main

import (
	"bytes"
	"fmt"
	"strings"
	"sync"
)

func main() {
	sentence := "The quick brown fox jumps over the lazy dog"
	words := _map(sentence)
	rwords := process(words)
	fmt.Println(reduce(rwords))
}

func process(words []string) []string {
	nosOfWords := len(words)
	buffChannel := make(chan string, nosOfWords)
	task := new(sync.WaitGroup)
	task.Add(nosOfWords)

	for _, word := range words {
		go func(word string) {
			defer task.Done()
			buffChannel <- reverse(word)
		}(word)
	}
	task.Wait()
	close(buffChannel)

	rwords := make([]string, 0)
	for rword := range buffChannel {
		rwords = append(rwords, rword)
	}
	return rwords
}

func _map(sentence string) []string {
	return strings.Split(sentence, " ")
}

func reduce(reverseWords []string) string {
	return strings.Join(reverseWords, " ")
}

func reverse(word string) string {
	var buff bytes.Buffer
	for index := len(word) - 1; index >= 0; index-- {
		buff.WriteString(string(word[index]))
	}
	return buff.String()
}
