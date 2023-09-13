package crypto

import (
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
	"os"
	"reflect"
)

const (
	encode = iota
	decode
)

func DecodeConfig(value interface{}, key string) error {
	if reflect.TypeOf(value).Kind() != reflect.Pointer {
		panic("DecodeConfig need pointer")
	}
	val := reflect.ValueOf(value)
	typ := reflect.TypeOf(value)

	return decodeOrEncodeConfig(val, typ, key, decode)
}

func EncodeConfig(value interface{}, key string) error {
	if reflect.TypeOf(value).Kind() != reflect.Pointer {
		panic("DecodeConfig need pointer")
	}
	val := reflect.ValueOf(value)
	typ := reflect.TypeOf(value)

	return decodeOrEncodeConfig(val, typ, key, encode)
}

func decodeOrEncodeConfig(val reflect.Value, typ reflect.Type, key string, types int) error {
	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}
	if typ.Kind() == reflect.Pointer {
		typ = typ.Elem()
	}
	numField := val.NumField()
	for i := 0; i < numField; i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)
		tag := fieldType.Tag.Get("crypto")
		if field.Kind() == reflect.Struct {
			if err := decodeOrEncodeConfig(field, fieldType.Type, key, types); err != nil {
				return err
			}
		} else if field.Kind() == reflect.Pointer {
			field = field.Elem()
			if err := decodeOrEncodeConfig(field, fieldType.Type, key, types); err != nil {
				return err
			}
		} else if field.Kind() == reflect.String {
			if field.IsZero() {
				continue
			}
			if tag != "-" {
				continue
			}
			s := field.Interface().(string)
			var (
				str string
				err error
			)
			if types == encode {
				str, err = DesEncoding(s, key)
			} else {
				str, err = DesDecoding(s, key)
			}
			if err != nil {
				return err
			}
			if field.CanSet() {
				field.SetString(str)
			}
		} else {
			continue
		}
	}

	return nil
}

func WriteEncodeConfig(value interface{}, key, src, out string, options ...viper.DecoderConfigOption) error {
	viper.SetConfigFile(src)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	if err := viper.Unmarshal(value, options...); err != nil {
		return err
	}
	if err := EncodeConfig(value, key); err != nil {
		return err
	}
	bytes, err := yaml.Marshal(value)
	if err != nil {
		return err
	}
	file, err := os.OpenFile(out, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	defer file.Close()
	if _, err := file.Write(bytes); err != nil {
		return err
	}

	return nil
}
