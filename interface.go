package main

type Migration interface {
	Up() string
	Down() string
}
