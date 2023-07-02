package extern

import (
	"encoding/json"
	"errors"
	"time"
)

type Exif struct {
	Aperture      string    `json:"apt"`
	Timestamp     time.Time `json:"t"`
	ShutterSpeed  string    `json:"ss"`
	FocalLength   string    `json:"fl"`
	FocalLengthFF string    `json:"fl_ff"`
	Iso           string    `json:"iso"`
	Lens          string    `json:"l"`
	Camera        string    `json:"c"`
	Software      string    `json:"s"`
	Width         int       `json:"w"`
	Height        int       `json:"h"`
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
