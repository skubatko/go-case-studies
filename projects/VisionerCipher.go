package main

import (
	"fmt"
	"strings"
)

func main() {
	cipherText := "CSOITEUIWUIZNSROCNKFD"
	//cipherText := "D"
	keyword := "GOLANG"
	//keyword := "B"

	const alphabetLen = 'Z' - 'A' + 1
	keywordLen := len(keyword)
	for i, c := range cipherText {
		result := 'A' + (alphabetLen+c-int32(keyword[i%keywordLen]))%alphabetLen
		fmt.Printf("%c", result)
	}
	fmt.Println()

	// solution
	message := ""
	keyIndex := 0
	for i := 0; i < len(cipherText); i++ {
		// A=0, B=1, ... Z=25
		c := cipherText[i] - 'A'
		k := keyword[keyIndex] - 'A'

		// зашифрованная буква - ключевая буква
		c = (c-k+26)%26 + 'A'
		message += string(c)

		// увеличение keyIndex
		keyIndex++
		keyIndex %= len(keyword)
	}
	fmt.Println(message)

	// part 2

	plainText := "your message goes here"
	plainText = strings.ReplaceAll(plainText, " ", "")
	plainText = strings.ToUpper(plainText)
	for i, c := range plainText {
		result := 'A' + (alphabetLen+c+int32(keyword[i%keywordLen]))%alphabetLen
		fmt.Printf("%c", result)
	}
	fmt.Println()

	// solution
	message = "your message goes here"
	keyword = "golang"
	keyIndex = 0
	cipherText = ""

	message = strings.ToUpper(strings.Replace(message, " ", "", -1))
	keyword = strings.ToUpper(strings.Replace(keyword, " ", "", -1))

	for i := 0; i < len(message); i++ {
		c := message[i]
		if c >= 'A' && c <= 'Z' {
			// A=0, B=1, ... Z=25
			c -= 'A'
			k := keyword[keyIndex] - 'A'

			// зашифрованная буква + ключевая буква
			c = (c+k)%26 + 'A'

			// увеличить keyIndex
			keyIndex++
			keyIndex %= len(keyword)
		}
		cipherText += string(c)
	}
	fmt.Println(cipherText)
}
