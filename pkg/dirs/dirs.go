package dirs

import (
	"io/fs"
	"io/ioutil"
	"log"
	"os"
)

func ListDirs(dir string) []fs.FileInfo {
	if dirExists(dir) {
		dirlisting, err := ioutil.ReadDir(dir)
		if err != nil {
			log.Fatalf("Error reading dir: %s", dir)
		}
		return dirlisting
	} else {
		return nil
	}

}
func dirExists(dir string) bool {
	_, err := os.Stat(dir)

	if os.IsNotExist(err) {
		return false
	}
	return true

}
