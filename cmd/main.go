package main

import "go-project/domain/user"

func main() {
	createUser(user.CreateUser{})
}

func createUser(createUser user.CreateUserContract) {
	createUser.CreateUser()
}
