package main

import (
	"fmt"
	"github.com/armando-couto/goutils"
)

func main() {
	Up()
}

func Up() {
	files := goutils.ListFolderFiles(".")
	for _, file := range files {
		fmt.Println(file)
	}
}

func Down() {

}
