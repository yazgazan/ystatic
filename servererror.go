package main

type ServerError struct {
	What string
	Code int
}

func (e ServerError) Error() string {
	return e.What
}
