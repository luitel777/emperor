package main

import (
	"fmt"
	"os"
	"path/filepath"
)

var errorFolderPermission string = "[Error] Cannot create folder. Might be a permission issue"

func createFolderIfNotExists() {
	if err := os.Mkdir("./data", 0750); err != nil && !os.IsExist(err) {
		// one can also use fs.ErrPermission
		fmt.Println(errorFolderPermission)
	}
	listFiles()
}

func listFiles() int {
	os.DirFS("./data")
	// a, _ := os.Open("./data/1658279373215053.png")
	// b, _ := io.ReadAll(a)
	a, _ := filepath.Glob("./data/*")
	return len(a)
}
