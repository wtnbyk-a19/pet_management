package Repository

import "pet_management/src/Domain/User"

type UserRepository interface {
	Insert(user *User.User) (*User.User, error)
}
