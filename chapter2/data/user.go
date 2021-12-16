package data

import "time"

// user struct created to handle the information that gets queried
type User struct {
	Id        int
	Uuid      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

// session struct created to hable the session of a user that gets authenticated
type Session struct {
	Id        int
	Uuid      string
	Email     string
	UserId    int
	CreatedAt time.Time
}

// returns an user or error based on the query
// i really like the fact the values that are going to be returned
// are already declared on the statemente just after the params of the fucntion
// and the fact how Go expects to receive an error, or just populate the user pointer
// that is created just aboive the QueryRow call.
// Another curious think is how they return without the need to specify
// the values that are needed, the return statement can infer those values
// by its usage on the function
func UserByEmail(email string) (user User, err error) {
	user = User{}
	err = Db.QueryRow("SELECT id, uuid, name, email, password, created_at FROM users WHERE email = $1", email).
		Scan(&user.Id, &user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	// return user, err (is not needed at all)
	return
}

// Couple things to note here:
// user as a pointer have sense because we are just getting the attribute values to be used
// as part of the session, (we don't really need a copy user User to use the values of session to modify
// their value stored on its attributes)
func (user *User) CreateSession() (session Session, err error) {
	// TODO: on the original code they use a variable to store the string
	// that is being used on the Prepare method, don't really know if
	// they want to be clear about what was that string, or if is something
	// related on the way how the string is being passed to the Prepare method
	stmt, err := Db.Prepare("insert into sessions (uuid, email, user_id, created_at) values ($1, $2, $3, $4) returning id, uuid, email, user_id, created_at")
	if err != nil {
		return
	}
	// as any other Prepare statement this need to be closed.
	defer stmt.Close()

	// Looks like query row is actually an execution of the prepared statement so far
	// this execution gets a response that can be extracted with the scan Method
	err = stmt.QueryRow(createUUID(), user.Email, user.Id, time.Now()).
		Scan(&session.Id, &session.Email, &session.Id, &session.CreatedAt)
	return
}
