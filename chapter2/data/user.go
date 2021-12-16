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
