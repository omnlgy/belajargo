package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
)

func main() {
	fmt.Println("Hello world")

	title := getUserInput("Title : ")
	content := getUserInput("Content : ")

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println("Error :", err)
		return
	}

	userNote.Display()

	err = userNote.Save()

	if err != nil {
		fmt.Println("Error saving note:", err)
		return
	}

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
