package compare

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"command/procedures/file"
)

type Subcommand struct {
	Dir1   string
	Dir2   string
	Delete bool
}

func (r *Subcommand) Name() string {
	return "compare"
}

func (r *Subcommand) Parse(args []string) error {
	fs := flag.NewFlagSet(r.Name(), flag.ContinueOnError)
	fs.StringVar(&r.Dir1, "dir1", "", "Dir 1")
	fs.StringVar(&r.Dir2, "dir2", "", "Dir 2")
	fs.BoolVar(&r.Delete, "delete", false, "Delete files from dir2")

	return fs.Parse(args)
}

func (r *Subcommand) Run() error {
	files1, err := file.GetDscFiles(r.Dir1)
	if err != nil {
		log.Fatal(err)
	}

	files2, err := file.GetDscFiles(r.Dir2)
	if err != nil {
		log.Fatal(err)
	}

	for _, file2 := range files2 {
		found := false
		for _, file1 := range files1 {
			if file.RemoveFileExtension(file1.Name()) == file.RemoveFileExtension(file2.Name()) {
				found = true
				break
			}
		}
		if !found {
			fmt.Printf("File %s not found in %s\n", file2.Name(), r.Dir2)

			if r.Delete {
				if err := os.Remove(filepath.Join(r.Dir2, file2.Name())); err != nil {
					log.Fatal(err)
				}
			}
		}
	}

	return nil
}
