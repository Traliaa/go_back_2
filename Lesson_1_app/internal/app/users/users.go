package users

type ImplementationDatabase interface {
	CreateUser()
	DeleteUser()
	SearchUser()
}

type UserRepositories struct {
	db ImplementationDatabase
}

func NewUserRepositories(db ImplementationDatabase) *UserRepositories {
	return &UserRepositories{db: db}
}

func (u *UserRepositories) Create() {
	u.db.CreateUser()
}
func (u *UserRepositories) Search() {
	u.db.SearchUser()
}
func (u *UserRepositories) Delete() {
	u.db.DeleteUser()
}
