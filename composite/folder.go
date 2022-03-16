package main

import "fmt"

type folder struct {
	name       string
	components []component
}

func (f *folder) search(keyword string) {
	fmt.Println("searching", keyword, "in folder", f.getName())
	for _, composite := range f.components {
		composite.search(keyword)
	}
}

func (f *folder) getName() string {
	return f.name
}
func (f *folder) add(c component) {
	f.components = append(f.components, c)
}
