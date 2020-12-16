package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isElementInArray(el byte, array []byte) bool {
	for _, v := range array {
		if v == el {
			return true
		}
	}
	return false
}

func genInitConst() ([]byte, []byte) {

	vowels := []byte{65, 69, 73, 79, 85, 97, 101, 105, 111, 117}
	var consonants []byte
	consonants = make([]byte, 42, 42)

	var letter byte
	letter = 65
	for i := 0; i < 42; i++ {
		if !(isElementInArray(letter, vowels)) {
			consonants[i] = letter
		} else {
			i--
		}
		letter++
		if letter == 91 {
			letter = 97
		}
	}

	return vowels, consonants

}



func main() {

	var vowels, consonants []byte

	vowelAppend := "yay"
	consonantAppend := "ay"
	reader := bufio.NewReader(os.Stdin)
	text, _:= reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")

	initWords := strings.Split(text, " ")

	vowels, consonants = genInitConst()

	var length int
	consCount := 0
	wordsQuantity := len(initWords)
	var resultWords []string
	resultWords = make([]string, wordsQuantity, wordsQuantity+2)
	for i, initWord := range initWords {
		if isElementInArray(initWord[0], vowels) {
			resultWords[i] = initWord + vowelAppend
		} else if isElementInArray(initWord[0], consonants) {
			consCount = 1

			length = len(initWord)

			for j := 1; j < length; j++ {
				if isElementInArray(initWord[j], consonants) {
					consCount++
				} else {
					break
				}
			}

			resultWords[i] = initWord[consCount:length] + initWord[0:consCount] + consonantAppend

		} else {
			resultWords[i] = initWord
		}

	}

	resultText := ""

	for i, resultWord := range resultWords {
		resultText = resultText + resultWord
		if i != wordsQuantity - 1 {
			resultText = resultText + " "
		}
	}

	fmt.Println(resultText)

}

