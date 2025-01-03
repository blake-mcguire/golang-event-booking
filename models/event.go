package models

import (
	"time"
	"log"
	"github.com/blake-mcguire/golang-event-booking/main/db"
)

type Event struct {
	ID          int64     
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"datetime" binding:"required"`
	UserID      int       `json:"user_id"`
}

var events = []Event{}

func (e *Event) Save() error {
	//LATER ADD THE EVENT TO A DATBASE
	query := `INSERT INTO events(name, description, location, datetime, user_id) VALUES (?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		log.Printf("Error preparing query: %v", err)
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)

	if err != nil {
		log.Printf("Error executing query: %v", err)
		return err
		
	}
	id, err := result.LastInsertId()
	e.ID = id
	events = append(events, *e)
	return err
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events" //creating a query string to use on the sql lite db
	rows, err := db.DB.Query(query) //if you have a query that just fetches data the function is query, If you have a function that changes stuff it is exec
	if err != nil {
		return nil, err
	}

	defer rows.Close() // make sure to close the db connection

	var events []Event //setting a slice of events to populate with the data from the rows

	for rows.Next() { //prepares the row for scanning,
		var event Event                                                                                               //setting a single event to populate for each row and then to add that event to the events slice
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID) // scans the row and then populates the data referenced in the params ie all the props of the event struct we jsut defiend
		if err != nil {
			return nil, err

		}

		events = append(events, event)
	}

	return events, nil
}


func GetEventById(id int64) (*Event, error){
	query := "SELECT * FROM events WHERE id = ?"

	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}

	return &event, nil 

}
