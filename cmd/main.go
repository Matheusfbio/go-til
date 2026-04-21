package main

import (
	"go-til/internal/user"
)

func main()  {
	var users []user.User

	user.AddUser(&users, user.NewUser("Matheus", 27, true))

	user.ListUsers(users)
}


// "go-til/internal/user"

// type User struct {
// 	name string
// 	age int
// }

// func (u User) IsAdult() bool {
// 	return u.age >= 18
// }


// 	userMap := map[string]int{
//     "Matheus": 25,
// }

// 	fmt.Println(userMap)
	// nums := []int{1, 2, 3}
	// nums = append(nums, 4)

	// fmt.Print(nums)

	// user := User{
	// 	name: "Matheus",
	// 	age: 27,
	// }

	// fmt.Println(user.IsAdult())
	// user := User{name: "Matheus Fabio", age: 27}
	// fmt.Println("user:", user.name, user.age)


	// for i := 1; i < 5; i++ {
	// 	fmt.Println(i)
	// }

	// age := 12

	// if age >= 18 {
	// 	fmt.Println("Ja pode ser preso kkkk")
	// } else {
	// 	fmt.Println("Jaja chega a sua hora")
	// }
// }



// func main() {
// 		user := user.NewUser("Matheus fabio", 27, "Authenticated")
// 		fmt.Println("user:", user.Name, user.Age, user.Autenticated)
// }



