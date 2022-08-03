package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	vips "github.com/davidbyttow/govips/v2/vips"
)

var errorFolderPermission string = "[Error] Cannot create folder. Might be a permission issue"

func createFolderIfNotExists() {
	if err := os.Mkdir("./data", 0750); err != nil && !os.IsExist(err) {
		// one can also use fs.ErrPermission
		fmt.Println(errorFolderPermission)
	}

	if err := os.Mkdir("./data/books", 0750); err != nil && !os.IsExist(err) {
		// one can also use fs.ErrPermission
		fmt.Println(errorFolderPermission)
	}

	if err := os.Mkdir("./data/cache", 0750); err != nil && !os.IsExist(err) {
		// one can also use fs.ErrPermission
		fmt.Println(errorFolderPermission)
	}
	listFiles()
}

func listFiles() []string {
	os.DirFS("./data/books")
	// a, _ := os.Open("./data/1658279373215053.png")
	// b, _ := io.ReadAll(a)
	a, _ := filepath.Glob("./data/books/*")
	return a
}

func listFileNames() []string {
	a, _ := ioutil.ReadDir("./data/books/")
	var b []string
	for _, f := range a {
		b = append(b, f.Name())
	}
	return b
}

func listCacheFiles() []string {
	os.DirFS("./data/cache")
	// a, _ := os.Open("./data/1658279373215053.png")
	// b, _ := io.ReadAll(a)
	a, _ := filepath.Glob("./data/cache/*")
	return a

}

func createCacheImages() {
	// vips starts here
	vips.Startup(nil)
	defer vips.Shutdown()

	for i, k := range listFiles() {
		image1, err := vips.NewImageFromFile(k)
		checkError(err)

		ep := vips.NewDefaultJPEGExportParams()
		image1bytes, _, err := image1.Export(ep)

		err = ioutil.WriteFile("./data/cache/"+listFileNames()[i]+".jpeg", image1bytes, 0644)
		checkError(err)
	}

	// vips ends here

}
