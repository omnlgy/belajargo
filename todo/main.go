package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface {
	Save() error
}

// type diplayer interface {
// 	Display()
// }

type outputter interface {
	saver
	// diplayer
	Display()
}

func main() {
	fmt.Println("Hello world")

	title := getUserInput("Title : ")
	content := getUserInput("Content : ")
	todoText := getUserInput("Todo : ")

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println("Error :", err)
		return
	}

	userTodo, err := todo.New(todoText)
	if err != nil {
		fmt.Println("Error :", err)
		return
	}

	err = outputData(&userNote)

	if err != nil {
		fmt.Println("Error saving note:", err)
		return
	}

	err = outputData(&userTodo)

	if err != nil {
		fmt.Println("Error saving todo:", err)
		return
	}
	printSometing(3.1)
	maxValue := max(1, 2)
	fmt.Println("Max value:", maxValue)
	maxValueFloat := max(2.1, 3.2)
	fmt.Println("Max value:", maxValueFloat)

	fmt.Println("Note saved successfully!")
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	text, err := bufio.NewReader(os.Stdin).ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input:", err)
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}

func saveData(data saver) error {
	return data.Save()
}

func outputData(data outputter) error {
	data.Display()
	// return data.Save()
	return saveData(data)
}

// func printSometing(val interface{}) {
func printSometing(val any) {
	switch val.(type) {
	case int:
		fmt.Println("Integer", val)
	case float64:
		fmt.Println("Float", val)
	case string:
		fmt.Println("String", val)
	default:
		fmt.Println("Unknown type")
	}

	typedVal, ok := val.(int)

	if ok {
		fmt.Println("Integer", typedVal)
	}

}

// generic
func max[T int | float32 | float64](a, b T) T {
	if a > b {
		return a
	}
	return b
}
