package main

import (
	"fmt"
	"go-til/internal/user"
)

func main() {
		user := user.NewUser("Matheus fabio", 27)
		fmt.Printf("Seu nome é: %s e sua idade é: %d\n", user.Name, user.Age)
}

