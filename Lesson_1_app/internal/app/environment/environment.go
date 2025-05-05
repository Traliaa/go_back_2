package environment

type ImplementationDatabase interface {
	CreateEnvironment()
	DeleteEnvironment()
	SearchEnvironment()
	SearchByUserInEnvironment()
	SearchByEnvInEnvironment()
}
type EnvRepositories struct {
	db ImplementationDatabase
}

func NewEnvironmentRepositories(db ImplementationDatabase) *EnvRepositories {
	return &EnvRepositories{db: db}
}
func (e *EnvRepositories) Create() {
	e.db.CreateEnvironment()
}
func (e *EnvRepositories) Search() {
	e.db.SearchEnvironment()
}
func (e *EnvRepositories) Delete() {
	e.db.DeleteEnvironment()
}
func (e *EnvRepositories) SearchByUser() {
	e.db.SearchByUserInEnvironment()
}
func (e *EnvRepositories) SearchByEnv() {
	e.db.SearchByEnvInEnvironment()
}
