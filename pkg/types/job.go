package types

type Job[T any] interface {
	Validate() error
	Authenticate() error
	Authorize() error
	Logic() (T, error)
	Handle() (T, error)
}
