package pkg

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var missedENV []interface{}

func LoadConfig(config interface{}) error {
	missedENV = nil

	if err := Deep(config); err != nil {
		return err
	}

	if missedENV != nil {
		result := make([]string, 0, len(missedENV))
		for _, val := range missedENV {
			result = append(result, fmt.Sprintf("%v", val))
		}

		return errors.New("missed ENV configs: [" + strings.Join(result, ", ") + "]")
	}

	return nil
}

func Deep(cfg interface{}) error {
	v := reflect.ValueOf(cfg).Elem()
	t := reflect.TypeOf(cfg).Elem()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)

		switch field.Kind() {
		case reflect.Struct:
			if err := Deep(field.Addr().Interface()); err != nil {
				return err
			}
		case reflect.Ptr:
			if field.Elem().Kind() == reflect.Struct {
				if err := Deep(field.Addr().Interface()); err != nil {
					return err
				}
			}
		}

		tag := t.Field(i).Tag.Get("env")
		if tag == "" {
			continue
		}

		required := strings.ToLower(t.Field(i).Tag.Get("required")) == "true"
		envByTag := os.Getenv(tag)

		if envByTag == "" && required {
			missedENV = append(missedENV, tag)
			continue
		}

		typ := field.Type()
		switch field.Kind() {
		case reflect.String:
			field.SetString(envByTag)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			var val int64
			var err error

			if field.Kind() == reflect.Int64 && typ.PkgPath() == "time" && typ.Name() == "Duration" {
				var d time.Duration
				d, err = time.ParseDuration(envByTag)
				val = int64(d)
			} else {
				val, err = strconv.ParseInt(envByTag, 0, typ.Bits())
			}
			if err != nil {
				return err
			}
			field.SetInt(val)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			val, err := strconv.ParseUint(envByTag, 0, typ.Bits())
			if err != nil {
				return err
			}
			field.SetUint(val)
		case reflect.Bool:
			val, err := strconv.ParseBool(envByTag)
			if err != nil {
				return err
			}
			field.SetBool(val)
		case reflect.Float32, reflect.Float64:
			val, err := strconv.ParseFloat(envByTag, typ.Bits())
			if err != nil {
				return err
			}
			field.SetFloat(val)

		case reflect.Slice:
			sl := reflect.MakeSlice(typ, 0, 0)
			if typ.Elem().Kind() == reflect.Uint8 {
				sl = reflect.ValueOf([]byte(envByTag))
			} else if len(strings.TrimSpace(envByTag)) != 0 {
				vals := strings.Split(envByTag, ",")
				sl = reflect.MakeSlice(typ, len(vals), len(vals))

				for i, val := range vals {
					f := sl.Index(i)
					switch f.Type().Kind() {
					case reflect.String:
						f.SetString(val)
					}
				}
			}
			field.Set(sl)
		}
	}

	return nil
}

func setVar(v interface{}, s string) error {
	if t := reflect.TypeOf(v); t.Kind() == reflect.Ptr &&
		t.Elem().Kind() == reflect.String {
		s = strconv.Quote(s)
	}
	return json.Unmarshal([]byte(s), v)
}
