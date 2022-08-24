package text

import (
	"errors"
	"reflect"
	"regexp"

	"github.com/go-playground/validator/v10"
)

var Validator = validator.New()

func ValidateIso8601Date(fl validator.FieldLevel) bool {
	ISO8601DateRegexString := "^(?:[1-9]\\d{3}-(?:(?:0[1-9]|1[0-2])-(?:0[1-9]|1\\d|2[0-8])|(?:0[13-9]|1[0-2])-(?:29|30)|(?:0[13578]|1[02])-31)|(?:[1-9]\\d(?:0[48]|[2468][048]|[13579][26])|(?:[2468][048]|[13579][26])00)-02-29)T(?:[01]\\d|2[0-3]):[0-5]\\d:[0-5]\\d(?:\\.\\d{1,9})?(?:Z|[+-][01]\\d:[0-5]\\d)$"
	ISO8601DateRegex := regexp.MustCompile(ISO8601DateRegexString)
	return ISO8601DateRegex.MatchString(fl.Field().String())
}

func init() {
	_ = Validator.RegisterValidation("iso8601date", ValidateIso8601Date)
}

func Validate(schema any, tag ...string) error {
	s := reflect.ValueOf(schema)

	if s.Kind() == reflect.Ptr {
		s = s.Elem()
	}

	if s.Kind() == reflect.Struct {
		var excepts []string
		for i := 0; i < s.NumField(); i++ {
			field := s.Field(i)
			kind := field.Kind()

			if kind == reflect.Ptr {
				field = field.Elem()
				kind = field.Kind()
			}

			if kind == reflect.Struct {
				excepts = append(excepts, s.Type().Field(i).Name)
				if err := Validate(field.Interface()); err != nil {
					return err
				}
			}

			if kind == reflect.Map || kind == reflect.Array || kind == reflect.Slice {
				excepts = append(excepts, s.Type().Field(i).Name)
				if err := Validate(field.Interface(), s.Type().Field(i).Tag.Get("validate")); err != nil {
					return err
				}
			}
		}
		return Validator.StructExcept(schema, excepts...)
	}

	if s.Kind() == reflect.Map {
		kind := s.Type().Elem().Kind()
		if kind == reflect.Ptr {
			kind = s.Type().Elem().Elem().Kind()
		}

		if kind == reflect.Struct {
			for _, e := range s.MapKeys() {
				val := s.MapIndex(e)
				if err := Validate(val.Interface()); err != nil {
					return err
				}
			}
		} else {
			for _, e := range s.MapKeys() {
				val := s.MapIndex(e)
				if err := Validator.Var(val, tag[0]); err != nil {
					return err
				}
			}
		}
	}

	if s.Kind() == reflect.Array || s.Kind() == reflect.Slice {
		if len(tag) != 1 {
			return errors.New("no tag inserted for an array")
		}

		kind := s.Type().Elem().Kind()
		if kind == reflect.Ptr {
			kind = s.Type().Elem().Elem().Kind()
		}

		if kind == reflect.Struct {
			for i := 0; i < s.Len(); i++ {
				if err := Validate(s.Index(i).Interface()); err != nil {
					return err
				}
			}
		} else {
			for i := 0; i < s.Len(); i++ {
				if err := Validator.Var(s.Index(i).Interface(), tag[0]); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
