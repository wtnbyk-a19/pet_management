package User

// User TODO: 値オブジェクトの作成
type User struct {
	Id       int
	Name     string
	Email    string
	Password string
}

func NewUser(id int, name string, email string, password string) (*User, error) {
	// User TODO: ガード節の作成
	return &User{
		id, name, email, password,
	}, nil
}
