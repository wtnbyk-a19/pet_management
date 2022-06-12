package Repository

import "pet_management/src/Domain/User"

type UserRepository interface {
	Insert(user *User.User) error
	Update(user *User.User) error
	FindById(id int) (*User.User, error)
}
