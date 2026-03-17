package models

import (
	"log/slog"
	"time"

	"github.com/look4suman/events-api/db"
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"date_time" binding:"required"`
	UserID      int       `json:"user_id"`
}

func (e Event) Save() (Event, error) {
	query := `INSERT INTO events (name, description, location, date_time, user_id) VALUES (?, ?, ?, ?, ?)`
	result, err := db.DB.Exec(query, e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return e, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return e, err
	}
	e.ID = id
	return e, nil
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT id, name, description, location, date_time, user_id FROM events`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

// pointer *Event is returned instead of Event because nil cannot be returned unless its a pointer.
// another alternative was to use Event{}
func GetEventById(id int64) (*Event, error) {
	query := `SELECT id, name, description, location, date_time, user_id FROM events where id = ?`
	result := db.DB.QueryRow(query, id)

	var event Event
	err := result.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (e Event) Update() error {
	query := `
	Update events
	set name = ?, description = ?, location = ?, date_time = ?, user_id = ?
	where id = ?
	`
	slog.Info("sql",
		"query", query,
		"args", []any{e.Name, e.Description, e.Location, e.DateTime, e.UserID, e.ID},
	)

	_, err := db.DB.Exec(query, e.Name, e.Description, e.Location, e.DateTime, e.UserID, e.ID)
	if err != nil {
		return err
	}
	return nil
}
