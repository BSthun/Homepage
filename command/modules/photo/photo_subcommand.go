package photo

import (
	"encoding/json"
	"flag"
	"fmt"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	_ "golang.org/x/image/tiff"

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

	fmt.Println("photo_section_id;shooter_id;image_path;thumbnail_path;raw_path;exif;created_at;updated_at")
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

		// im, _, err := image.DecodeConfig(imgFile)
		// if err != nil {
		// 	fmt.Fprintf(os.Stderr, "%s: %v\n", imgFile.Name(), err)
		// 	continue
		// }

		jsonByte, err := metaData.MarshalJSON()
		if err != nil {
			log.Fatal(err.Error())
		}

		var exifRaw *ExifRaw
		err = json.Unmarshal(jsonByte, &exifRaw)

		t, err2 := time.Parse("2006:01:02 15:04:05 -0700", exifRaw.DateTimeOriginal+" +0700")
		if err2 != nil {
			log.Fatal(err.Error())
		}

		var aperture = "1.2"
		if exifRaw.ApertureValue != nil {
			// fmt.Println(exifRaw.ApertureValue)
			// aperture = strings.Split(exifRaw.ApertureValue[0], "/")[0]
			aperture = "1.8"
		}

		fl := "35"
		fl35 := "52"
		if exifRaw.FocalLength != nil {
			foc := strings.Split(exifRaw.FocalLength[0], "/")
			fl = foc[0]
			if len(foc) > 1 {
				foc0, _ := strconv.ParseInt(foc[0], 10, 32)
				foc1, _ := strconv.ParseInt(foc[1], 10, 32)
				fl = fmt.Sprintf("%d", foc0/foc1)
			}
			fl35 = strconv.Itoa(exifRaw.FocalLengthIn35mmFilm[0])
		} else {
			exifRaw.LensModel = "RISESPRAY 35mm F1.2"
		}

		if exifRaw.LensModel == "----" {
			exifRaw.LensModel = "RISESPRAY 35mm F1.2"
		}

		photos = append(photos, &PhotoItem{
			PhotoSectionId: 5,
			ImagePath:      "/image/" + file.Name(),
			ThumbnailPath:  "/thumbnail/" + strings.Split(file.Name(), ".")[0] + "_thumbnail.jpg",
			RawPath:        "/raw/" + strings.Split(file.Name(), ".")[0] + ".arw",
			ShooterId:      1,
			Exif: &Exif{
				Aperture:        aperture,
				Timestamp:       t.UTC(),
				ShutterSpeed:    exifRaw.ExposureTime[0],
				FocalLength:     fl,
				FocalLength35mm: fl35,
				Iso:             strconv.Itoa(exifRaw.IsoSpeedRatings[0]),
				Lens:            exifRaw.LensModel,
				Camera:          exifRaw.Make + " " + exifRaw.Model,
			},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		})

		byte, err := json.Marshal(photos[i].Exif)
		fmt.Printf("%d;%d;%s;%s;%s;%s;%s;%s\n", photos[i].PhotoSectionId, photos[i].ShooterId, photos[i].ImagePath, photos[i].ThumbnailPath, photos[i].RawPath, string(byte), photos[i].CreatedAt.Format(time.RFC3339), photos[i].UpdatedAt.Format(time.RFC3339))
	}

	// jsbyte, _ := json.Marshal(photos)
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
	//
	// fmt.Println(string(jsbyte))
	return nil
}
