package data

import (
	"crypto/rand"
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"
)

// exported member to get access to the dabatase
var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "dbname=chitchat sslmode= disable")
	if err != nil {
		log.Fatal(err)
	}
	//TODO: why do this fuction have a return statement ?
	return
}

// Not really sure how sha works, but something to note here
// is how the param and return value names were declared
// the names belongs to the functionality not to the very
// specific usage of it, in other case the names would by something
// like password, and encryptedpass respectively or something like that
func Encrypt(plaintext string) (crptext string) {
	crptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return
}

// TODO: cannot they import any package that already have this fuction?
func createUUID() (uuid string) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatalln("Cannot generate UUID", err)
	}

	// 0x40 is reserved variant from RFC 4122
	u[8] = (u[8] | 0x40) & 0x7F
	// Set the four most significant bits (bits 12 through 15) of the
	// time_hi_and_version field to the 4-bit version number.
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
}
