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

func UserNameSlice() []string {
	nameSlice := []string{}
	var point int
	fmt.Println("Сколько имен вы хотите ввести? ")
	fmt.Scanln(&point)
	for i := 0; i < point; i++ {
		var name string
		fmt.Scanln(&name)
		nameSlice = append(nameSlice, name)
	}
	return nameSlice
}

func userInput() string {
	fmt.Println("С какой буквы ищем имена? ")
	var input string
	fmt.Scanln(&input)
	return strings.TrimSpace(input)
}

func normalHomework(s string, nameSlice []string) []string {
	var res []string
	for _, name := range nameSlice {
		if strings.HasPrefix(strings.ToLower(name), strings.ToLower(s)) {
			res = append(res, name)
		}
	}
	return res
}
