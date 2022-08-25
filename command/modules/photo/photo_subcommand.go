package photo

import (
	"encoding/json"
	"flag"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

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
	// fmt.Printf("INFO | Variable DIR: %s\n", r.Dir)

	files, err := GetFiles(r.Dir)
	if err != nil {
		log.Fatal(err)
	}

	var photos []*PhotoItem

	fmt.Println("photo_section_id;shooter_id;root;image_path;thumbnail_path;raw_path;exif;ratio;created_at;updated_at")
	for i, file := range files {
		if i == -1 {
			continue
		}

		imgFile, err := os.Open(filepath.Join(r.Dir, file.Name()))
		if err != nil {
			log.Fatal(err.Error())
		}

		metaData, err := exif.Decode(imgFile)
		if err != nil {
			log.Fatal(err.Error())
		}

		imgFile.Seek(0, 0)

		im, _, err := image.DecodeConfig(imgFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", imgFile.Name(), err)
			continue
		}

		jsonByte, err := metaData.MarshalJSON()
		if err != nil {
			log.Fatal(err.Error())
		}

		var exif *ExifRaw
		err = json.Unmarshal(jsonByte, &exif)

		t, err2 := time.Parse("2006:01:02 15:04:05", exif.DateTimeOriginal)
		if err2 != nil {
			log.Fatal(err.Error())
		}

		var aperture string = "1.2"
		if exif.ApertureValue != nil {
			aperture = strings.Split(exif.ApertureValue[0], "/")[0]
		}

		var fl string = "35"
		var fl35 string = "52"
		if exif.FocalLength != nil {
			fl = strings.Split(exif.FocalLength[0], "/")[0]
			fl35 = strconv.Itoa(exif.FocalLengthIn35mmFilm[0])
		}

		if exif.LensModel == "----" {
			exif.LensModel = "RISESPRAY 35mm F1.2"
		}

		photos = append(photos, &PhotoItem{
			PhotoSectionId: 2,
			Ratio:          float32(im.Width) / float32(im.Height),
			Root:           "https://macmini.lan.bsthun.com/static",
			ImagePath:      "/csfd2022/" + file.Name(),
			ThumbnailPath:  "/csfd2022/" + strings.Split(file.Name(), ".")[0] + "_thumbnail.jpg",
			RawPath:        "/csfd2022/" + strings.Split(file.Name(), ".")[0] + ".ARW",
			ShooterId:      1,
			Exif: &Exif{
				Aperture:        aperture,
				Brightness:      exif.BrightnessValue[0],
				Credit:          exif.Artist + ", " + exif.Copyright,
				Timestamp:       t,
				ExposureTime:    exif.ExposureTime[0],
				FocalLength:     fl,
				FocalLength35mm: fl35,
				IsoSpeed:        strconv.Itoa(exif.IsoSpeedRatings[0]),
				LensModel:       exif.LensModel,
				CameraMaker:     exif.Make,
				CameraModel:     exif.Model,
			},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		})

		byte, err := json.Marshal(photos[i].Exif)
		fmt.Printf("%d;%d;%s;%s;%s;%s;%s;%f;%s;%s\n", photos[i].PhotoSectionId, photos[i].ShooterId, photos[i].Root, photos[i].ImagePath, photos[i].ThumbnailPath, photos[i].RawPath, string(byte), photos[i].Ratio, photos[i].CreatedAt.Format(time.RFC3339), photos[i].UpdatedAt.Format(time.RFC3339))
	}

	// jsbyte, _ := json.Marshal(photos)
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
	//
	// fmt.Println(string(jsbyte))
	return nil
}
