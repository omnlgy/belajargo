package models

import (
	"errors"
	"fmt"
	"time"

	"example.com/rest/db"
	"github.com/google/uuid"
)

var (
	ErrNoRowsAffected = errors.New("no rows affected")
)

type Event struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name" binding:"required,alphaspace"`
	Description string    `json:"description" binding:"required,alphaspace"`
	Location    string    `json:"location" binding:"required,alphanumspace"`
	DateTime    time.Time `json:"dateTime" binding:"required"`
	UserId      uuid.UUID `json:"userId"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
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

	rows, err := db.DB.Query("SELECT id, name, description, location, dateTime, userId, createdAt, updatedAt FROM events")

	if err != nil {
		return events, fmt.Errorf("failed to query events: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var event Event

		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId, &event.CreatedAt, &event.UpdatedAt)
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
	SET name = ?, description = ?, location = ?, dateTime = ?, userId = ?, updatedAt = ?
	WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return fmt.Errorf("failed to prepare the update query %w", err)
	}

	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId, time.Now(), e.ID)

	if err != nil {
		return fmt.Errorf("failed to execute the update query %w", err)
	}

	if affectedCount, _ := result.RowsAffected(); affectedCount == 0 {
		return ErrNoRowsAffected
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
		return ErrNoRowsAffected
	}

	return nil
}

func GetEventById(eventId uuid.UUID) (Event, error) {
	var event Event

	row := db.DB.QueryRow("SELECT id, name, description, location, dateTime, userId, createdAt, updatedAt FROM events WHERE id = ?", eventId)

	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId, &event.CreatedAt, &event.UpdatedAt)

	if err != nil {
		return event, fmt.Errorf("failed to scan event: %w", err)
	}

	return event, nil
}
