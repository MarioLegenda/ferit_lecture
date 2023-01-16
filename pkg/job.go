package pkg

type Job[T any] interface {
	Validate() error
	Authenticate() error
	Authorize() error
	Handle() (T, error)
}
