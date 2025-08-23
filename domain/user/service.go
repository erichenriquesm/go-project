package user

import "fmt"

type CreateUser struct{}

func (c CreateUser) CreateUser() int {
	fmt.Println("User created!")
	return 1
}
