package user

// User represent the basic information for user
// throug to the room
type User struct {
	Name   string `db:"name"`
	Gender string `db:"gender"`
	Age    int    `db:"age"`
}
