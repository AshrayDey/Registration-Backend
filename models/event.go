package models

import (
	"registrationApp/db"
	"time"
)

type Event struct {
	ID          int64
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Location    string `binding:"required"`
	DateTime    time.Time
	UserID      int64
}

func (e *Event) Save() error {
	query := `
	INSERT INTO events(name, description, location, datetime, user_id)
	VALUES (?,?,?,?,?)
	`
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()
	result, err := statement.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	e.ID = id
	// events = append(events, e)
	return err
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * from events"
	rows, err := db.DB.Query(query)
	var events []Event
	var event Event

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id)
	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}

	return &event, err
}
func (e Event) Update() error {
	query := `
  	UPDATE events
   	SET name = ?, description = ?, location = ?, datetime = ?
    WHERE id = ?
	`
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(e.Name, e.Description, e.Location, e.DateTime, e.ID)
	return err
}

func (event Event) Delete() error {
	query := `
	DELETE FROM events
	WHERE id = ?
	`
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(event.ID)
	return err
}
