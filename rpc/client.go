package whatis_rpc

type User struct {
	FirstName string
	LastName  string
	Age       int32
}

func NewUser(firstName, lastName string, age int32) *User {
	return &User{firstName, lastName, age}
}

func (u *User) GetFullName() string {
	return u.FirstName + " " + u.LastName
}
