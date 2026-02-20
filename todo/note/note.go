package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}

func New(title, content string) (note, error) {
	if title == "" || content == "" {
		return note{}, errors.New("All fields are required")
	}

	return note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

func (note note) Display() {
	fmt.Printf("%s\n%s\n", note.Title, note.Content)
}

func (note *note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName)
	jsonStuct, err := json.Marshal(note)

	if err != nil {
		return err
	}

	return os.WriteFile("./saveFiles/"+fileName+".json", jsonStuct, 0644)
}
