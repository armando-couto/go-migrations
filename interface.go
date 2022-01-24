package main

/*
	https://gobyexample.com/interfaces
*/
type Migration interface {
	Up() string
	Down() string
}
