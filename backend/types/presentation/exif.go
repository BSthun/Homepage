package presentation

import (
	"encoding/json"
	"errors"
	"time"
)

type Exif struct {
	Aperture        string    `json:"aperture"`
	Brightness      string    `json:"brightness"`
	Credit          string    `json:"credit"`
	Timestamp       time.Time `json:"timestamp"`
	ExposureTime    string    `json:"exposure_time"`
	FocalLength     string    `json:"focal_length"`
	FocalLength35mm string    `json:"focal_length_35mm"`
	IsoSpeed        string    `json:"iso_speed"`
	LensModel       string    `json:"lens_model"`
	CameraMaker     string    `json:"camera_maker"`
	CameraModel     string    `json:"camera_model"`
}

func (r *Exif) Scan(src any) error {
	var source []byte
	switch src.(type) {
	case string:
		source = []byte(src.(string))
	case []byte:
		source = src.([]byte)
	default:
		return errors.New("incompatible type for Exif")
	}

	if err := json.Unmarshal(source, r); err != nil {
		return err
	}
	return nil
}