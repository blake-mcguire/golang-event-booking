package models
import (
	"time"
	"github.com/blake-mcguire/golang-event-booking/main/db"
)
type Event struct {
	ID int64
	Name string `binding:"required"`
	Description string `binding:"required"`
	Location string `binding:"required"`
	DateTime time.Time `binding:"required"`
	UserID int 
}


var events = []Event{}
func (e *Event) Save() error {
	//LATER ADD THE EVENT TO A DATBASE
	query := `INSERT INTO events(name, description, location, datetime, user_id) VALUES ("?, ?, ?, ?, ?")`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err 
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)

	if err != nil {
	 return err
	}
	id, err := result.LastInsertId()
	e.ID = id
	events = append(events, *e)
	return err 
}	


func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query) //if you have a query that just fetches data the function is query, If you have a function that changes stuff it is exec
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		rows.Scan()
	}


	return events, nil
}