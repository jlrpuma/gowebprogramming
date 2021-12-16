package data

import "time"

type Thread struct {
	Id        int
	Uuid      string
	Topic     string
	UserId    int
	CreatedAt time.Time
}

// get all threads
func Threads() (threads []Thread, err error) {
	// executing the query and store the value on rows variable
	rows, err := Db.Query("SELECT id, uuid, topic, user_id, created_at FROM threads ORDER BY created_at DESC")
	// at any error on the query execution is being handled here
	if err != nil {
		return
	}

	//TODO: this is something cool to see
	// an iterator Next() funcion on the rows
	for rows.Next() {
		// creating the empty structure to hold the values
		conv := Thread{}
		// TODO: interesting part, they handle an error in case Scan fail, or
		// populate the conv variable with the values that are being scanned from
		// the rows variable

		// TODO: wold be great if we can avoid this for and get an implementation
		// were you can get on the elements scanned/appended  in one line
		if err = rows.Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAt); err != nil {
			return
		}
		threads = append(threads, conv)
	}
	// closing rows that were queried
	rows.Close()
	return
}
