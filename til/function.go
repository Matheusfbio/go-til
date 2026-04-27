package main

import "fmt"

type User struct {
	Name string
	Age  int
}

var users []User

func main() {
	CreateUser("Matheus Fabio", 27)
	CreateUser("Matheus Fabio", 27)
	CreateUser("Matheus Fabio", 27)
	fmt.Println("--------lista de usuarios---------")
	listUser()

}

func CreateUser(name string, age int) {
	if age < 0 {
		fmt.Println("invalid age")
		return
	}
	user := User{Name: name, Age: age}
	users = append(users, user)
	fmt.Printf("User created successfully!\n")
	fmt.Printf("Name: %s, Age: %d\n", user.Name, user.Age)
}

func listUser() {
	if len(users) == 0 {
		fmt.Println("No users found.")
		return
	}


	for i, user := range users {
		fmt.Printf("User %d: Name: %s, Age: %d\n", i, user.Name, user.Age)
	}
}

// func sum(x int, y int) int{
// 	return x + y
// }