package main

import (
	"fmt"
	"strings"
)

func main() {
	userNameList := UserNameSlice()
	userInputChar := userInput()
	resNormal := normalHomework(userInputChar, userNameList)
	fmt.Println("Результат задания уровня normal:", resNormal)
}

// UserNameSlice asks user how many names he wants to enter, and then
// asks for each name. Returns slice of entered names.
func UserNameSlice() []string {
	nameSlice := []string{}
	var point int
	fmt.Println("Сколько имен вы хотите ввести? ")
	fmt.Scanln(&point)
	for i := 0; i < point; i++ {
		var name string
		fmt.Println("Введите имя: ")
		fmt.Scanln(&name)
		nameSlice = append(nameSlice, name)
	}
	return nameSlice
}

// userInput asks user for a character, trims the input and returns it.
func userInput() string {
	fmt.Println("С какой буквы ищем имена? ")
	var input string
	fmt.Scanln(&input)
	return strings.TrimSpace(input)
}

// normalHomework filters and returns the list of names from nameSlice
// that start with the given prefix 's', ignoring case sensitivity.
func normalHomework(s string, nameSlice []string) []string {
	var res []string
	for _, name := range nameSlice {
		if strings.HasPrefix(strings.ToLower(name), strings.ToLower(s)) {
			res = append(res, name)
		}
	}
	return res
}