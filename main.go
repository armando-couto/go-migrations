package main

import (
	"fmt"
	"github.com/armando-couto/goutils"
)

const (
	PATH = "./migrations"
)

func main() {
	Run()
}

func Run() {
	// Sempre checa se a Paste existe
	goutils.CreateDirectory(PATH)
	// Find de Files
	files := goutils.ListFolderFiles(PATH + "/")
	for _, file := range files {
		fmt.Println(file)
	}
}
