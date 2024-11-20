package main

import (
	"bufio"
	"fmt"
	"os"
)

type Writer interface {
	Write([]byte) (int, error)
}

type File struct {
	Name string
	Writer
}

func (f File) Write(data []byte) (int, error) {
	file, err := os.Create(f.Name)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	fmt.Println("Файл создан")
	return file.Write(data)
}

func UserInput() string {
	fmt.Println("Введите текст: ")
	text, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return ""
	}
	return text
}

type Console struct {}

type ConsolePrinter interface {
	Print(string)
}

func (c Console) Print(text string) {
	fmt.Println(text)
}

func main() {
	f := File{Name: "file.txt"}
	res, _ := f.Write([]byte(UserInput()))
	c := Console{}
	c.Print(string(res))
}
