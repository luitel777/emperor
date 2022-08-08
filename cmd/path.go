package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	vips "github.com/davidbyttow/govips/v2/vips"
)

var errorFolderPermission string = "[Error] Cannot create folder. Might be a permission issue"

func checkErrorAndDie(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func createFolderIfNotExists() {
	var arrayOfFolders []string = []string{"./data", "./data/books", "./data/cache"}
	for _, folderNames := range arrayOfFolders {
		// one can also use fs.ErrPermission
		if err := os.Mkdir(folderNames, 0750); err != nil && !os.IsExist(err) {
			fmt.Println(errorFolderPermission)
		}
	}

	listFiles()
}

func listFiles() []string {
	files, err := filepath.Glob("./data/books/*")
	checkErrorAndDie(err)
	return files
}

func listFileNames() []string {
	files, err := ioutil.ReadDir("./data/books/")
	checkErrorAndDie(err)

	var filesArray []string
	for _, f := range files {
		filesArray = append(filesArray, f.Name())
	}
	return filesArray
}

func listCacheFiles() []string {
	os.DirFS("./data/cache")
	imageFiles, err := filepath.Glob("./data/cache/*")
	checkErrorAndDie(err)

	return imageFiles
}

func createCacheImages() {
	vips.Startup(nil)
	defer vips.Shutdown()

	for i, k := range listFiles() {
		image1, err := vips.NewImageFromFile(k)
		checkErrorAndDie(err)

		ep := vips.NewDefaultJPEGExportParams()
		image1bytes, _, err := image1.Export(ep)

		err = ioutil.WriteFile("./data/cache/"+listFileNames()[i]+".jpeg", image1bytes, 0644)
		checkErrorAndDie(err)
	}

}
