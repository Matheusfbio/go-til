package user

import "fmt"

func AddUser(users *[]User, u User) {
    *users = append(*users, u)
}

func ListUsers(users []User) {
    for _, u := range users {
        adult := "menor"
        if u.IsAdult() {
            adult = "adulto"
        }
        fmt.Printf("Nome: %s | Idade: %d | %s\n", u.Name, u.Age, adult)
    }
}
