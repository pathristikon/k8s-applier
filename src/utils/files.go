package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)


/** Check file extension */
func CheckExtension(i os.FileInfo) bool {
	extensions := []string{".yaml", ".yml"}

	for _, ext := range extensions {
		if filepath.Ext(i.Name()) == ext {
			return true
		}
	}

	return false
}


/** Loop in the files */
func ReadFiles(dirname string, configParams Config) []string {
	dir := fmt.Sprintf("%s/%s", configParams.ConfigFolder, dirname)
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	var list []string

	for _, f := range files {
		// passing over dirs
		if f.IsDir() {
			continue
		}

		// checking if extension is available
		if CheckExtension(f) {
			list = append(list, f.Name())
		}
	}

	return list
}
