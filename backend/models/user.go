package models

// User represent the basic information for user
// through to the room
type User struct {
	Name   string `db:"name"`
	Gender string `db:"gender"`
	Age    int    `db:"age"`
}
