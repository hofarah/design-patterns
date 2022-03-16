package main

import "fmt"

type file struct {
	name string
}

func (f *file) search(keyword string) {
	fmt.Println("searching", keyword, "in file", f.getName())
}

func (f *file) getName() string {
	return f.name
}
