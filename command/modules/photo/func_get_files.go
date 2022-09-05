package photo

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func GetFiles(dir string) ([]fs.FileInfo, error) {
	// * Get absolute file path
	abs, err := filepath.Abs(dir)
	if err != nil {
		return nil, fmt.Errorf("unable to get absolute path: %s", err.Error())
	}

	// * Check that the folder exist
	if _, err := os.Stat(abs); os.IsNotExist(err) {
		return nil, fmt.Errorf("directory is not exist")
	}

	// * Read directory information
	files, err := ioutil.ReadDir(abs)
	if err != nil {
		log.Fatal(err)
	}

	// * Filter only image files
	var filtered []fs.FileInfo
	for _, file := range files {
		if strings.HasPrefix(file.Name(), "DSC") && strings.HasSuffix(file.Name(), ".jpg") {
			filtered = append(filtered, file)
		}
	}

	return filtered, nil
}
