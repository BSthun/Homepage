package photo

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/rwcarlsen/goexif/exif"
)

type Subcommand struct {
	Dir string
}

func (r *Subcommand) Name() string {
	return "photo"
}

func (r *Subcommand) Parse(args []string) error {
	fs := flag.NewFlagSet(r.Name(), flag.ContinueOnError)
	fs.StringVar(&r.Dir, "dir", "", "Dir of the photo")

	return fs.Parse(args)
}

func (r *Subcommand) Run() error {
	fmt.Printf("INFO | Variable DIR: %s\n", r.Dir)

	files, err := GetFiles(r.Dir)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(files)

	imgFile, err := os.Open(filepath.Join(r.Dir, files[500].Name()))
	if err != nil {
		log.Fatal(err.Error())
	}

	metaData, err := exif.Decode(imgFile)
	if err != nil {
		log.Fatal(err.Error())
	}

	jsonByte, err := metaData.MarshalJSON()
	if err != nil {
		log.Fatal(err.Error())
	}

	jsonString := string(jsonByte)
	fmt.Println(jsonString)

	return nil
}
