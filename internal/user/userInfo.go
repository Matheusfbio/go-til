package user

type User struct {
	Name string
	Age  int
	autenticated bool
}

func (u User) IsAdult() bool {
    return u.Age >= 18
}

func NewUser(name string, age int, autenticated bool) User {
    return User{Name: name, Age: age, autenticated: autenticated}
}

