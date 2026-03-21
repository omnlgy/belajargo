package models

import (
	"fmt"
	"time"

	"example.com/rest/db"
	"github.com/google/uuid"
)

type Event struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name" binding:"required,alphaspace"`
	Description string    `json:"description" binding:"required,alphaspace"`
	Location    string    `json:"location" binding:"required,alphanumspace"`
	DateTime    time.Time `json:"dateTime" binding:"required"`
	UserId      uuid.UUID `json:"userId"`
}

func (e Event) Save() error {
	query := `
	INSERT INTO events (id, name, description, location, dateTime, userId) 
	VALUES (?, ?, ?, ?, ?, ?)
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return fmt.Errorf("failed to prepare the save query %w", err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID, e.Name, e.Description, e.Location, e.DateTime, e.UserId)

	if err != nil {
		return fmt.Errorf("failed to execute the save query %w", err)
	}

	return nil
}

func GetAllEvents() ([]Event, error) {
	var events []Event

	rows, err := db.DB.Query("SELECT * FROM events")

	if err != nil {
		return events, fmt.Errorf("failed to query events: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var event Event

		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
		if err != nil {
			return events, fmt.Errorf("failed to scan event: %w", err)
		}

		events = append(events, event)
	}

	return events, nil
}

func (e Event) Update() error {
	query := `
	UPDATE events 
	SET name = ?, description = ?, location = ?, dateTime = ?, userId = ?
	WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return fmt.Errorf("failed to prepare the update query %w", err)
	}

	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId, e.ID)

	if err != nil {
		return fmt.Errorf("failed to execute the update query %w", err)
	}

	if affectedCount, _ := result.RowsAffected(); affectedCount == 0 {
		return fmt.Errorf("no rows affected")
	}

	return nil
}

func DeleteEvent(eventId uuid.UUID) error {
	query := `
	DELETE FROM events 
	WHERE id = ?
	`

	result, err := db.DB.Exec(query, eventId)

	if err != nil {
		return fmt.Errorf("failed to execute the delete query %w", err)
	}

	if affectedCount, _ := result.RowsAffected(); affectedCount == 0 {
		return fmt.Errorf("no rows affected")
	}

	return nil
}
