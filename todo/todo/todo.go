package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type todo struct {
	Text string `json:"text"`
}

func New(text string) (todo, error) {
	if text == "" {
		return todo{}, errors.New("All fields are required")
	}

	return todo{
		Text: text,
	}, nil
}

func (todo todo) Display() {
	fmt.Println(todo.Text)
}

func (todo *todo) Save() error {
	fileName := "todo"
	jsonStuct, err := json.Marshal(todo)

	if err != nil {
		return err
	}

	return os.WriteFile("./saveFiles/"+fileName+".json", jsonStuct, 0644)
}
