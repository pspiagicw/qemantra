package dir

import (
	"github.com/pspiagicw/goreland"
	"io/fs"
	"os"
)

func ListDir(dir string) []fs.DirEntry {
	if dirExists(dir) {
		dirlisting, err := os.ReadDir(dir)
		if err != nil {
			goreland.LogFatal("Error reading dir: %s", dir)
		}
		return dirlisting

	}
	return []fs.DirEntry{}
}
func dirExists(dir string) bool {
	_, err := os.Stat(dir)

	if os.IsNotExist(err) {
		return false
	}
	return true

}
