package dirs

import (
    "io/fs"
	"os"
	"path/filepath"
	"testing"
)

// func getFileContents() map[string]string {
// 	files := make(map[string]string)
// 	files["hello"] = "Hello World"
// 	files["sello"] = "Sello World"
// 	return files
// }
func TestListDirs(t *testing.T) {
	t.Run("Directory exists", func(t *testing.T) {
		dirpath, err := os.MkdirTemp("", "listing")
		if err != nil {
			t.Fatalf("Error creating temp dir %v ", err)
		}
		defer os.RemoveAll(dirpath)
		for _, name := range []string{"hello" , "sello"} {
			tmpfilepath := filepath.Join(dirpath, name)
			os.WriteFile(tmpfilepath, []byte(""), 0644)

		}
		wanted := []string{
			"hello",
			"sello",
		}
		got := ListDirs(dirpath)
		for i, want := range wanted {
			if got[i].Name() != want {
				t.Errorf("wanted %v  ,got %v", got, want)
			}
		}

	})
    t.Run("Directory does not exist" , func(t *testing.T) {
        notExistentDir := "/tmp/get-out"

        var want []fs.FileInfo = nil
        got := ListDirs(notExistentDir)

        if want != nil {
            t.Errorf("wanted %v , got %v" , got , want)
        }

    })
}
