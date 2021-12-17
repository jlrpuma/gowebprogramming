package main

import "database/sql"

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "user=gwp dbname=gwp password=gwp sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func retrieve(id int) (post Post, err error) {
	post = Post{}
	err = Db.QueryRow("select id, content, author from posts where id = $1", id).
		Scan(&post.Id, &post.Content, &post.Author)
	return
}

// TODO: a bad habit that i get from Java is passing the Post
// that i wan't to create on the params of this method
// but in go it's diferent because you can make the Post struct
// compose the create method on it
// With this oyu have the right implementation just to be used from the
// post created struct
func (post *Post) create() (err error) {
	// I did not know about that returning word...
	statement := "insert into posts (content, author) values ($1, $2) returning id"
	// Db.Prepare, remember that, Db.Prepare
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	// QueryRow executes the statement and the params can be passed just right there
	// then we will need the id of the created post
	// Wy do we use post.Content and post.Author instead
	// &postContent and &post.Author... thats because we dont need
	// to dereferenciate the post to get the content on it
	// but in case we want to change those values we wold have to
	// use it like &post.id is doing it, to actually change the value
	// that is being stored there
	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
	if err != nil {
		return
	}
	return
}

func (post *Post) update() (err error) {
	_, err = Db.Exec("update posts set content = $2, author = $3 where id = $1", post.Id, post.Content, post.Author)
	return
}

func (post *Post) delete() (err error) {
	_, err = Db.Exec("delete from posts where id = $1", post.Id)
	return
}
