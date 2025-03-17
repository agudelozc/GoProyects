package main

import (
	"fmt"
)

// UserManager se encarga solo de la gestión de usuarios
type UserManager struct {
	users map[int]string
}

func NewUserManager() *UserManager {
	return &UserManager{users: make(map[int]string)}
}

func (um *UserManager) AddUser(id int, name string) {
	um.users[id] = name
}

func (um *UserManager) GetUser(id int) string {
	return um.users[id]
}

func main() {
	um := NewUserManager()
	um.AddUser(1, "Alice")
	fmt.Println(um.GetUser(1))
}
