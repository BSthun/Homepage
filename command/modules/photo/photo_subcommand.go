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

	fmt.Println("photo_section_id;shooter_id;root;image_path;thumbnail_path;raw_path;exif;created_at;updated_at")
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

		var exif *ExifRaw
		err = json.Unmarshal(jsonByte, &exif)

		t, err2 := time.Parse("2006:01:02 15:04:05 -0700", exif.DateTimeOriginal+" +0700")
		if err2 != nil {
			log.Fatal(err.Error())
		}

		var aperture = "1.2"
		if exif.ApertureValue != nil {
			aperture = strings.Split(exif.ApertureValue[0], "/")[0]

			// Force 1.8
			aperture = "1.8"
		}

		var fl = "35"
		var fl35 = "52"
		if exif.FocalLength != nil {
			foc := strings.Split(exif.FocalLength[0], "/")
			fl = foc[0]
			if len(foc) > 1 {
				foc0, _ := strconv.ParseInt(foc[0], 10, 32)
				foc1, _ := strconv.ParseInt(foc[1], 10, 32)
				fl = fmt.Sprintf("%d", foc0/foc1)
			}
			fl35 = strconv.Itoa(exif.FocalLengthIn35mmFilm[0])
		}

		if exif.LensModel == "----" {
			exif.LensModel = "RISESPRAY 35mm F1.2"
		}

		photos = append(photos, &PhotoItem{
			PhotoSectionId: 3,
			Root:           "https://proxy.bsthun.com/mixkoraspi/photo/csfdwalk2022",
			ImagePath:      "/image/" + file.Name(),
			ThumbnailPath:  "/thumbnail/" + strings.Split(file.Name(), ".")[0] + "_thumbnail.jpg",
			RawPath:        "/raw/" + strings.Split(file.Name(), ".")[0] + ".dng",
			ShooterId:      1,
			Exif: &Exif{
				Aperture:        aperture,
				Timestamp:       t.UTC(),
				ShutterSpeed:    exif.ExposureTime[0],
				FocalLength:     fl,
				FocalLength35mm: fl35,
				Iso:             strconv.Itoa(exif.IsoSpeedRatings[0]),
				Lens:            exif.LensModel,
				Camera:          exif.Make + exif.Model,
			},
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
		})

		byte, err := json.Marshal(photos[i].Exif)
		fmt.Printf("%d;%d;%s;%s;%s;%s;%s;%s;%s\n", photos[i].PhotoSectionId, photos[i].ShooterId, photos[i].Root, photos[i].ImagePath, photos[i].ThumbnailPath, photos[i].RawPath, string(byte), photos[i].CreatedAt.Format(time.RFC3339), photos[i].UpdatedAt.Format(time.RFC3339))
	}

	// jsbyte, _ := json.Marshal(photos)
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
	//
	// fmt.Println(string(jsbyte))
	return nil
}
