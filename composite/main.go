package main

func main() {
	file1 := &file{
		name: "test.js",
	}
	file2 := &file{
		name: "index.js",
	}
	file3 := &file{
		name: "env.js",
	}
	folder1 := &folder{
		name: "modules",
	}
	folder2 := &folder{
		name: "fold",
	}
	folder1.add(file1)
	folder1.add(file2)
	folder1.add(folder2)
	folder2.add(file3)
	folder1.search("ali")
}
