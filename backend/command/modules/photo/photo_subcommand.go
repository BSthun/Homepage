package photo

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/buckket/go-blurhash"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"path/filepath"
	"share/types/extern"
	"share/types/model"
	"share/utils/value"
	"strconv"
	"strings"
	"time"

	_ "golang.org/x/image/tiff"

	"command/procedures/file"
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

	files, err := file.GetDscFiles(r.Dir)
	if err != nil {
		log.Fatal(err)
	}

	var photos []*model.PhotoItem

	//fmt.Println("photo_section_id;shooter_id;image_path;thumbnail_path;raw_path;blurhash;exif;created_at;updated_at")
	for i, fileInfo := range files {
		if i == -1 {
			continue
		}

		imgFile, err := os.Open(filepath.Join(r.Dir, fileInfo.Name()))
		if err != nil {
			log.Fatal(err.Error())
		}

		metaData, err := exif.Decode(imgFile)
		if err != nil {
			log.Fatal(err.Error())
		}

		_, err = imgFile.Seek(0, 0)
		if err != nil {
			return err
		}

		img, _, err := image.Decode(imgFile)
		if err != nil {
			log.Fatal(err.Error())
		}

		hash, err := blurhash.Encode(5, 5, img)
		if err != nil {
			log.Fatal(err.Error())
		}

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

		currentTime := time.Now()
		camera := exifRaw.Make + " " + exifRaw.Model
		if strings.HasPrefix(exifRaw.Model, "NIKON") {
			camera = exifRaw.Model
		}

		photos = append(photos, &model.PhotoItem{
			Id:             nil,
			PhotoSection:   nil,
			PhotoSectionId: value.Ptr[uint64](8),
			Shooter:        nil,
			ShooterId:      value.Ptr[uint64](1),
			ImagePath:      value.Ptr("img/" + fileInfo.Name()),
			ThumbnailPath:  value.Ptr("tmb/" + strings.Split(fileInfo.Name(), ".")[0] + ".jpg"),
			RawPath:        value.Ptr("raw/" + strings.Split(fileInfo.Name(), ".")[0] + ".arw"),
			Blurhash:       &hash,
			Exif: &extern.Exif{
				Aperture:      aperture,
				Timestamp:     t.UTC(),
				ShutterSpeed:  exifRaw.ExposureTime[0],
				FocalLength:   fl,
				FocalLengthFF: fl35,
				Iso:           strconv.Itoa(exifRaw.IsoSpeedRatings[0]),
				Lens:          exifRaw.LensModel,
				Camera:        camera,
			},
			CreatedAt: &currentTime,
			UpdatedAt: &currentTime,
		})

		//bytes, err := json.Marshal(photos[i].Exif)
		//fmt.Printf("%d;%d;%s;%s;%s;%s;%s;%s;%s\n", *photos[i].PhotoSectionId, *photos[i].ShooterId, *photos[i].ImagePath, *photos[i].ThumbnailPath, *photos[i].Blurhash, *photos[i].RawPath, string(bytes), photos[i].CreatedAt.Format(time.RFC3339), photos[i].UpdatedAt.Format(time.RFC3339))
		fmt.Printf("UPDATE photo_items SET blurhash = '%s' WHERE image_path = '%s'\n", *photos[i].Blurhash, *photos[i].ImagePath)
	}

	// jsbyte, _ := json.Marshal(photos)
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
	//
	// fmt.Println(string(jsbyte))
	return nil
}
